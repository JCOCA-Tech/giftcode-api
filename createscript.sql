CREATE DATABASE giftcode_api_db();

USE giftcode_api_db;

#DONT FORGET TO CREATE AN USER FIRST AND GRANT PRIVILIGES ON THE RIGHT TABLE!

CREAT TABLE gitcode(
    id INT NOT NULL PRIMARY KEY auto_increment,
    gifcode VARCHAR(255) NOT NULL,
    origin_url VARCHAR(255),
    reg_date TIMESTAMP NOT NULL current_timestamp() ON UPDATE current_timestamp()
);