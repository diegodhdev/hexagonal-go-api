CREATE TABLE courses
(
    id       VARCHAR(255) NOT NULL,
    name     VARCHAR(255) NOT NULL,
    duration VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;

CREATE TABLE requests
(
    id            VARCHAR(255) NOT NULL,
    api           varchar(60)  NOT NULL,
    mode          varchar(32)  NOT NULL,
    response_type varchar(32)  NOT NULL,
    body          JSON DEFAULT NULL,

    PRIMARY KEY (id)

) CHARACTER SET utf8mb4
  COLLATE utf8mb4_bin;
