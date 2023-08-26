CREATE TABLE inputs
(
    id          SERIAL PRIMARY KEY,
    name        varchar,
    type_inputs BIGINT
);

CREATE TABLE services
(
    id   SERIAL PRIMARY KEY,
    name varchar
);


CREATE TABLE service_inputs
(
    id         SERIAL PRIMARY KEY,
    input_id   INT REFERENCES inputs (id),
    service_id INT REFERENCES services (id),
    required   boolean
);

CREATE TABLE users
(
    id       SERIAL PRIMARY KEY,
    fullName varchar
);

CREATE TABLE user_service_inputs
(
    id               SERIAL PRIMARY KEY,
    service_input_id INT REFERENCES service_inputs (id),
    user_id          int REFERENCES users (id),
    val              varchar,
    required         boolean
);