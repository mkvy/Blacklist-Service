package auth

import (
	"context"
	"encoding/json"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/utils"
	"github.com/shaj13/go-guardian/v2/auth"
	"github.com/shaj13/go-guardian/v2/auth/strategies/basic"
	"github.com/shaj13/go-guardian/v2/auth/strategies/jwt"
	"github.com/shaj13/go-guardian/v2/auth/strategies/union"
	"github.com/shaj13/libcache"
	_ "github.com/shaj13/libcache/fifo"
	"log"
	"net/http"
	"time"
)

var strategy union.Union
var keeper jwt.SecretsKeeper

func SetupGoGuardian() {
	keeper = jwt.StaticSecret{
		ID:        "secret-id",
		Secret:    []byte("hardcode"),
		Algorithm: jwt.HS256,
	}
	cache := libcache.FIFO.New(0)
	cache.SetTTL(time.Minute * 5)
	cache.RegisterOnExpired(func(key, _ interface{}) {
		cache.Peek(key)
	})
	basicStrategy := basic.NewCached(validateUser, cache)
	jwtStrategy := jwt.New(cache, keeper)
	strategy = union.New(jwtStrategy, basicStrategy)
}

func CreateToken(w http.ResponseWriter, r *http.Request) {
	u := auth.User(r)
	token, _ := jwt.IssueAccessToken(u, keeper)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(map[string]string{"token": token})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func validateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {

	if userName == "hardcode" && password == "hardcode" {
		return auth.NewDefaultUser(userName, "1", nil, nil), nil
	}

	return nil, utils.ErrWrongCredentials
}

func Middleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing Auth Middleware")
		_, user, err := strategy.AuthenticateRequest(r)
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		log.Printf("User %s Authenticated\n", user.GetUserName())
		r = auth.RequestWithUser(user, r)
		next.ServeHTTP(w, r)
	})
}
