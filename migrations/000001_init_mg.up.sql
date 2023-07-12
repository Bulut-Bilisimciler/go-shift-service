CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    employee_id BIGINT UNIQUE,
    name TEXT DEFAULT NULL,
    surname TEXT DEFAULT NULL,
    email TEXT DEFAULT NULL,
    company TEXT DEFAULT NULL,
    department TEXT DEFAULT NULL,
    directorship TEXT DEFAULT NULL,
    group_name TEXT DEFAULT NULL,
    username TEXT DEFAULT NULL,
    type TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE shifts (
    id SERIAL PRIMARY KEY,
    shift_id TEXT NOT NULL,
    employee_id BIGINT UNIQUE,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    group_name TEXT DEFAULT NULL
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
    employee_id BIGINT UNIQUE,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    deleted_at TIMESTAMPTZ DEFAULT NULL
);



INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (1, 'John', 'Doe', 'john@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'johndoe', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (2, 'Jane', 'Smith', 'jane@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'janesmith', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (3, 'Jack', 'Jones', 'jake@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'jackjones', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (4, 'Jill', 'Brown', 'jill@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'jillbrown', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (5, 'James', 'Williams', 'james@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'jameswilliams', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (6, 'Jennifer', 'Davis', 'jennifer@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'jenniferdavis', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (7, 'Jeff', 'Miller', 'jeff@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'jeffmiller', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (8, 'Jessica', 'Wilson', 'jessica@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'jessicawilson', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (9, 'John', 'Moore', 'john@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'johnmoore', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (10, 'Jennifer', 'Taylor', 'jennifer@example.com', 'ACME', 'HR', 'Human Resources', 'HR', 'jennifertaylor', 'PARTNER');

INSERT INTO users (employee_id, name, surname, email, company, department, directorship, group_name, username, type)
VALUES (11, 'James', 'Anderson', 'james@example.com', 'ACME', 'IT', 'Data Center', 'IT', 'jamesanderson', 'PARTNER');

INSERT INTO shifts (shift_id, employee_id, start_time, end_time, group_name)
VALUES ('S001', 1, '2023-06-12 09:00:00', '2023-06-12 17:00:00', 'IT');

INSERT INTO shifts (shift_id, employee_id, start_time, end_time, group_name)
VALUES ('S002', 2, '2023-06-13 13:00:00', '2023-06-13 21:00:00', 'HR');


INSERT INTO shift_periods (shift_period_id, start_time, end_time)
VALUES (1, '2023-06-12 09:00:00', '2023-06-12 12:00:00');

INSERT INTO shift_periods (shift_period_id, start_time, end_time)
VALUES (2, '2023-06-12 12:00:00', '2023-06-12 15:00:00');


INSERT INTO demands (demand_id, shift_id, employee_id, start_time, end_time)
VALUES (1, 'S001', 1, '2023-06-12 09:00:00', '2023-06-12 17:00:00');

INSERT INTO demands (demand_id, shift_id, employee_id, start_time, end_time)
VALUES (2, 'S002', 2, '2023-06-13 13:00:00', '2023-06-13 21:00:00');