CREATE TABLE IF NOT EXISTS login(
	id SERIAL PRIMARY KEY,
	email VARCHAR(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	age INT NOT NULL
)

SELECT age FROM login;

SELECT * FROM login;

DELETE FROM login;