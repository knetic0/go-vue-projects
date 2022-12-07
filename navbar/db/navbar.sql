CREATE TABLE IF NOT EXISTS navbar(
	id SERIAL PRIMARY KEY,
	title VARCHAR(20)
);

INSERT INTO navbar(title) VALUES('Homepage');
INSERT INTO navbar(title) VALUES('About Us');
INSERT INTO navbar(title) VALUES('Downloads');
INSERT INTO navbar(title) VALUES('Products');
INSERT INTO navbar(title) VALUES('Contact');
INSERT INTO navbar(title) VALUES('Sign in');
INSERT INTO navbar(title) VALUES('Register');



SELECT * FROM navbar;
