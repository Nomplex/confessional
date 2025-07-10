DROP TABLE IF EXISTS confessions;


CREATE TABLE confessions (
	id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	content VARCHAR(280) NOT NULL,
	created DATETIME NOT NULL
);

CREATE INDEX idx_confessions_created ON confessions(created);

INSERT INTO confessions (content, created) 
VALUES ("Embarassing Confession", UTC_TIMESTAMP());

INSERT INTO confessions (content, created)
VALUES ("Silly Confession", UTC_TIMESTAMP());

INSERT INTO confessions (content, created)
VALUES ("Sad Confession", UTC_TIMESTAMP());
