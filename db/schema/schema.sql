CREATE TYPE gender AS ENUM ('M', 'F');

CREATE TYPE status AS ENUM ('ACTIVE', 'INACTIVE');

CREATE TYPE roles AS ENUM ('ADMIN', 'GENERAL', 'ASISTANT');

CREATE TABLE
    users (
        id bigserial PRIMARY KEY,
        full_name varchar NOT NULL,
        email varchar UNIQUE NOT NULL,
        gender gender NOT NULL,
        status status NOT NULL,
        id_department bigint,
        id_town bigint,
        direction varchar(200),
        create_at timestamptz NOT NULL DEFAULT (now ()),
        id_clinic bigserial NOT NULL,
        rol roles NOT NULL
    );

CREATE TABLE 
    clinic (
        id bigserial PRIMARY KEY,
        name varchar NOT NULL,
        phone varchar(50),
        id_department bigint,
        id_town bigint,
        direction varchar(200),
        status status NOT NULL,
        create_at timestamptz NOT NULL DEFAULT (now ())
    );

CREATE TABLE
    department (
        id bigserial PRIMARY KEY,
        name varchar(250) NOT NULL
    );

CREATE TABLE
    town (
        id bigserial PRIMARY KEY,
        name varchar(250) NOT NULL,
        id_department bigserial NOT NULL
    );

ALTER TABLE users ADD FOREIGN KEY (id_department) REFERENCES department (id);

ALTER TABLE users ADD FOREIGN KEY (id_town) REFERENCES town (id);

ALTER TABLE users ADD FOREIGN KEY (id_clinic) REFERENCES clinic (id);

ALTER TABLE clinic ADD FOREIGN KEY (id_department) REFERENCES department (id);

ALTER TABLE clinic ADD FOREIGN KEY (id_town) REFERENCES town (id);

ALTER TABLE town ADD FOREIGN KEY (id_department) REFERENCES department (id);