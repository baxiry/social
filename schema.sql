
CREATE DATABASE IF NOT EXISTS social;


CREATE TABLE IF NOT EXISTS social.users (
    userid int unsigned NOT NULL AUTO_INCREMENT,
    username varchar(255) NOT NULL DEFAULT "",
    password varchar(255) NOT NULL ,
    email varchar(255) UNIQUE NOT NULL,
    gender varchar(1) NOT NULL,
    photos text NOT NULL DEFAULT "",
    number_photos int NOT NULL DEFAULT 0,
    PRIMARY KEY (userid)
);


-- UPDATE EXAMPLE
-- UPDATE  users set gender = 'f' where userid = 1 ;
-- UPDATE table_a  SET  col3 = 'some_value', col4 = 'some_other_value WHERE  id = 3 ;
