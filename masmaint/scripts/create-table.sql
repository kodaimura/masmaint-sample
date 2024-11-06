CREATE TABLE IF NOT EXISTS department (
	code TEXT,
	name TEXT NOT NULL,
	description TEXT,
	manager_id INTEGER,
	location TEXT,
	budget NUMERIC NOT NULL,
	created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
	updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
	PRIMARY KEY(code)
);

CREATE TABLE IF NOT EXISTS employee (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	first_name TEXT NOT NULL,
	last_name TEXT,
	email TEXT,
	phone_number TEXT,
	address TEXT,
	hire_date TEXT,
	job_title TEXT,
	department_code TEXT,
	salary NUMERIC,
	created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
	updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);


CREATE TRIGGER IF NOT EXISTS trg_department_upd AFTER UPDATE ON department
BEGIN
	UPDATE department
	SET updated_at = DATETIME('now', 'localtime')
	WHERE rowid == NEW.rowid;
END;

CREATE TRIGGER IF NOT EXISTS trg_employee_upd AFTER UPDATE ON employee
BEGIN
	UPDATE employee
	SET updated_at = DATETIME('now', 'localtime')
	WHERE rowid == NEW.rowid;
END;