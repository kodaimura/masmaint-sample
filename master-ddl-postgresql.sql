create function set_update_time() returns trigger AS '
    BEGIN
      new.updated_at := NOW();
      return new;
    END;
' language 'plpgsql';

CREATE TABLE employee (
  id SERIAL PRIMARY KEY,
  first_name TEXT NOT NULL,
  last_name TEXT,
  email TEXT,
  phone_number TEXT,
  address TEXT,
  hire_date DATE,
  job_title TEXT,
  department_id INTEGER,
  salary INTEGER
);

CREATE TABLE department (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL UNIQUE,
  description TEXT,
  manager_id INTEGER,
  location TEXT,
  budget NUMERIC(10, 2),
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

create trigger trg_department_upd BEFORE UPDATE ON department FOR EACH ROW
  execute procedure set_update_time();