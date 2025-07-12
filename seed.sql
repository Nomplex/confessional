DROP TABLE IF EXISTS confessions;


CREATE TABLE confessions (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	title VARCHAR(30) NOT NULL,
	content VARCHAR(280) NOT NULL,
	created DATETIME NOT NULL
);

CREATE INDEX idx_confessions_created ON confessions(created);

INSERT INTO confessions (title, content, created) 
VALUES ("I once killed 80 people 367890", "Embarassing Confession", UTC_TIMESTAMP());

INSERT INTO confessions (title, content, created) 
VALUES ("Silly Boy Hours", "Silly Confession", UTC_TIMESTAMP());

INSERT INTO confessions (title, content, created) 
VALUES ("Sad Boy Hours", "Sad Confession", UTC_TIMESTAMP());
