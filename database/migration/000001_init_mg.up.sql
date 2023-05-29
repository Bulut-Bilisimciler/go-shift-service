CREATE TABLE users (
                     id uuid PRIMARY KEY,
                     name VARCHAR(64) NOT NULL,
                     surname VARCHAR(64) NOT NULL,
                     email VARCHAR(255) DEFAULT NULL,
                     made_field VARCHAR(255) DEFAULT NULL,
                     nickname VARCHAR(64) NOT NULL,
                     user_id INT NOT NULL,
                     created_at TIMESTAMPTZ DEFAULT NOW(),
                     updated_at TIMESTAMPTZ DEFAULT NOW(),
                     is_active BOOLEAN DEFAULT true
);
CREATE TABLE shifts (
                      id uuid PRIMARY KEY,
                      shift_id VARCHAR(16) NOT NULL,
                      user_id uuid DEFAULT NULL,
                      start_time TIMESTAMPTZ NOT NULL,
                      end_time TIMESTAMPTZ NOT NULL,
                      created_at TIMESTAMPTZ DEFAULT NOW(),
                      updated_at TIMESTAMPTZ DEFAULT NOW(),
                      made_field VARCHAR(255) DEFAULT NULL,
                      is_active BOOLEAN DEFAULT true
);
CREATE TABLE shift_periods (
                             id uuid PRIMARY KEY,
                             shift_periods_id VARCHAR(16) NOT NULL,
                             start_time TIMESTAMPTZ NOT NULL,
                             end_time TIMESTAMPTZ NOT NULL,
                             created_at TIMESTAMPTZ DEFAULT NOW(),
                             updated_at TIMESTAMPTZ DEFAULT NOW(),
                             is_active BOOLEAN DEFAULT true
);
CREATE TABLE demands (
                       id uuid PRIMARY KEY,
                       demand_id VARCHAR(64) NOT NULL,
                       user_id VARCHAR(64) NOT NULL,
                       old_shift_id VARCHAR(64) NOT NULL,
                       new_shift_id  VARCHAR(255) NOT NULL,
                       created_at TIMESTAMPTZ DEFAULT NOW(),
                       updated_at TIMESTAMPTZ DEFAULT NOW(),
                       is_active BOOLEAN DEFAULT true
);

select * from users


