CREATE TABLE users (
  id INT IDENTITY(1,1) PRIMARY KEY,
  name VARCHAR(64) NOT NULL,
  surname VARCHAR(64) NOT NULL,
  email VARCHAR(255) DEFAULT NULL,
  made_field VARCHAR(255) DEFAULT NULL,
  nickname VARCHAR(64) NOT NULL,
  user_id INT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE shifts (
  id INT IDENTITY(1,1)  PRIMARY KEY, 
  shift_id VARCHAR(16) NOT NULL,
  user_id INT DEFAULT NULL,
  start_time TIMESTAMPTZ NOT NULL,
  end_time TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  made_field VARCHAR(255) DEFAULT NULL
);
CREATE TABLE shift_periods (
  id INT IDENTITY(1,1)  PRIMARY KEY,
  shift_periods_id VARCHAR(16) NOT NULL,
  start_time TIMESTAMPTZ NOT NULL,
  end_time TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE TABLE demands (
  id INT IDENTITY(1,1)  PRIMARY KEY,
  demand_id VARCHAR(64) NOT NULL,
  user_id VARCHAR(64) NOT NULL,
  old_shift_id VARCHAR(64) NOT NULL,
  new_shift_id  VARCHAR(255) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW(),
  deleted_at TIMESTAMPTZ DEFAULT NULL
);



/* INSERT INTO public.users
(id, "name", surname, email, made_field, nickname, user_id, created_at, updated_at)
VALUES(4, 'asdasdasd', 'asdasdad', 'email', 'made_field', 'asdasd', 0, now(), now());

INSERT INTO public.shifts
(id, shift_id, user_id, start_time, end_time, created_at, updated_at, made_field)
VALUES(1, 'qwerqwe', 1,  now(),  now(), now(), now(),  'qweqwe'); 
*/
