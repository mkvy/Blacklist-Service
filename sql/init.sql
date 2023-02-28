CREATE TABLE IF NOT EXISTS blacklist(
    id varchar(50) PRIMARY KEY,
    phone_number varchar(20) NOT NULL,
    user_name varchar(150) NOT NULL,
    ban_reason varchar(500) NOT NULL,
    date_banned timestamp NOT NULL,
    username_who_banned varchar(150) NOT NULL,
    UNIQUE (phone_number, user_name, username_who_banned)
    );