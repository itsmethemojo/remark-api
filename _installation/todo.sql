DROP TABLE todo CASCADE;
CREATE TABLE todo
(
   id       INT                                         NOT NULL,
   user_id  INT                                         NOT NULL,
   data     TEXT CHARSET utf8 COLLATE utf8_general_ci   NOT NULL
)
ENGINE=InnoDB
COLLATE=utf8_general_ci;

ALTER TABLE todo
   ADD CONSTRAINT pk_todo
   PRIMARY KEY (id);