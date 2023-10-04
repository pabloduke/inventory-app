CREATE TABLE item (
	category VARCHAR(50),
	name VARCHAR(50) PRIMARY KEY NOT NULL,
	desc VARCHAR(100)  NULL,
	price DECIMAL  NULL
);

INSERT INTO item
VALUES ('Food', 'Apple', 'Kinda healthy, better options available', 1.99);

INSERT INTO item
VALUES ('Food', 'Orange', 'Kinda healthy, better options available', 1.99);

INSERT INTO item
VALUES ('Drink', 'Coffee', 'Makes you smarter', 1.99);