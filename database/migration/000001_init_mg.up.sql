CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       user_id VARCHAR(50) NOT NULL UNIQUE,
                       name VARCHAR(50) NOT NULL,
                       surname VARCHAR(50) NOT NULL,
                       email VARCHAR(100) NOT NULL UNIQUE,
                       nickname VARCHAR(50) NOT NULL UNIQUE,
                       creation_date TIMESTAMP DEFAULT NOW(),
                       update_date TIMESTAMP DEFAULT NOW(),
                       made_fiel VARCHAR(50) NOT NULL
);

CREATE TABLE shifts (
                        id SERIAL PRIMARY KEY,
                        shift_id VARCHAR(50) NOT NULL UNIQUE,
                        user_id VARCHAR(50) NOT NULL REFERENCES users(user_id),
                        start_time TIMESTAMP NOT NULL,
                        end_time TIMESTAMP NOT NULL,
                        status VARCHAR(50) NOT NULL,
                        creation_date TIMESTAMP DEFAULT NOW(),
                        update_date TIMESTAMP DEFAULT NOW(),
                        user_id_to_keep_shift VARCHAR(50) NOT NULL,
                        made_fiel VARCHAR(50) NOT NULL,
                        is_active BOOLEAN DEFAULT true
);

CREATE TABLE shift_periods (
                               id SERIAL PRIMARY KEY,
                               shift_periods_id VARCHAR(50) NOT NULL UNIQUE,
                               start_date TIMESTAMP NOT NULL,
                               end_date TIMESTAMP NOT NULL,
                               creation_date TIMESTAMP DEFAULT NOW(),
                               update_date TIMESTAMP DEFAULT NOW(),
                               is_active BOOLEAN DEFAULT true
);

CREATE TABLE demands (
                         id SERIAL PRIMARY KEY,
                         demand_id VARCHAR(50) NOT NULL UNIQUE,
                         user_id VARCHAR(50) NOT NULL REFERENCES users(user_id),
                         old_shift_id VARCHAR(50) NOT NULL REFERENCES shifts(shift_id),
                         new_shift_id VARCHAR(50) NOT NULL REFERENCES shifts(shift_id),
                         status VARCHAR(50) NOT NULL,
                         creation_date TIMESTAMP DEFAULT NOW(),
                         update_date TIMESTAMP DEFAULT NOW()
);
