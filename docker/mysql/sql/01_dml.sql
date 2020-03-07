USE `library`;

INSERT INTO `author` (`id`, `name`)
VALUES
	(1,'William Shakespeare'),
	(2,'J. K. Rowling');


INSERT INTO `book` (`id`, `title`, `author_id`)
VALUES
	(1,'Hamlet',1),
	(2,'Romeo and Juliet',1),
	(3,'Harry Potter and the Philosopher\'s Stone',2),
	(4,'HarryPotter and the Chamber of Secrets',2);
