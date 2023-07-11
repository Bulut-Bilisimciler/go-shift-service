CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    user_id BIGINT DEFAULT NULL,
    name TEXT DEFAULT NULL,
    surname TEXT DEFAULT NULL,
    email TEXT DEFAULT NULL,
    company TEXT DEFAULT NULL,
    department TEXT DEFAULT NULL,
    directorship TEXT DEFAULT NULL,
    made_field TEXT DEFAULT NULL,
    username TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE shifts (
    id SERIAL PRIMARY KEY,
    shift_id TEXT NOT NULL,
    user_id BIGINT DEFAULT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    made_field TEXT DEFAULT NULL
);

CREATE TABLE shift_periods (
    id SERIAL PRIMARY KEY,
    shift_period_id BIGINT NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE demands (
    id SERIAL PRIMARY KEY,
    demand_id BIGINT NOT NULL,
    shift_id TEXT NOT NULL,
    user_id BIGINT DEFAULT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);



INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (1, 'John', 'Doe', 'john@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'johndoe');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (2, 'Jane', 'Smith', 'jane@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'janesmith');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (3, 'Jack', 'Jones', 'jake@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'jackjones');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (4, 'Jill', 'Brown', 'jill@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'jillbrown');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (5, 'James', 'Williams', 'james@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'jameswilliams');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (6, 'Jennifer', 'Davis', 'jennifer@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'jenniferdavis');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (7, 'Jeff', 'Miller', 'jeff@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'jeffmiller');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (8, 'Jessica', 'Wilson', 'jessica@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'jessicawilson');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (9, 'John', 'Moore', 'john@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'johnmoore');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (10, 'Jennifer', 'Taylor', 'jennifer@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'jennifertaylor');

INSERT INTO users (user_id, name, surname, email, company, department, directorship, made_field, username)
VALUES (11, 'James', 'Anderson', 'james@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'jamesanderson');

INSERT INTO shifts (shift_id, user_id, start_time, end_time, made_field)
VALUES ('S001', 1, '2023-06-12 09:00:00', '2023-06-12 17:00:00', 'IT');

INSERT INTO shifts (shift_id, user_id, start_time, end_time, made_field)
VALUES ('S002', 2, '2023-06-13 13:00:00', '2023-06-13 21:00:00', 'HR');


INSERT INTO shift_periods (shift_period_id, start_time, end_time)
VALUES (1, '2023-06-12 09:00:00', '2023-06-12 12:00:00');

INSERT INTO shift_periods (shift_period_id, start_time, end_time)
VALUES (2, '2023-06-12 12:00:00', '2023-06-12 15:00:00');


INSERT INTO demands (demand_id, shift_id, user_id, start_time, end_time)
VALUES (1, 'S001', 1, '2023-06-12 09:00:00', '2023-06-12 17:00:00');

INSERT INTO demands (demand_id, shift_id, user_id, start_time, end_time)
VALUES (2, 'S002', 2, '2023-06-13 13:00:00', '2023-06-13 21:00:00');