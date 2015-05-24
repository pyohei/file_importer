/* Create table */
CREATE database sleep_cycle;

CREATE TABLE sleep (
    no int(10) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    sleep_from datetime DEFAULT null,
    sleep_to datetime DEFAULT null,
    confort_rate tinyint(3) DEFAULT null,
    sleep_minute int(6) DEFAULT null,
    sleep_feeling tinyint(3) DEFAULT null,
    pulsation tinyint(3) DEFAULT null,
    memo varchar(255) DEFAULT null,
    walk_count int(10) DEFAULT null,
    regist_time datetime DEFAULT NULL,
    update_time timestamp NOT NULL
);
    

