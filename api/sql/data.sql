INSERT INTO users(name, nick, email, password)
values
("user1", "user1", "user1@mail", "$2a$10$pQtfO6S01K9mlv/bIYeD/.8weVsBajl.5dzf4XoTWtt0P6OBKu20e"),
("user2", "user2", "user2@mail", "$2a$10$pQtfO6S01K9mlv/bIYeD/.8weVsBajl.5dzf4XoTWtt0P6OBKu20e"),
("user3", "user3", "user3@mail", "$2a$10$pQtfO6S01K9mlv/bIYeD/.8weVsBajl.5dzf4XoTWtt0P6OBKu20e"),
("user4", "user4", "user4@mail", "$2a$10$pQtfO6S01K9mlv/bIYeD/.8weVsBajl.5dzf4XoTWtt0P6OBKu20e"),
("user5", "user5", "user5@mail", "$2a$10$pQtfO6S01K9mlv/bIYeD/.8weVsBajl.5dzf4XoTWtt0P6OBKu20e");

INSERT INTO followers(user_id, follower_id)
values
(1,2),
(2,3),
(3,4),
(2,5),
(4,6);

INSERT INTO publications (title, content, author_id)
values
("publication user 1", "Content publication user 1", 1),
("publication user 2", "Content publication user 2", 2),
("publication user 3", "Content publication user 3", 3),
("publication user 4", "Content publication user 4", 4),
("publication user 5", "Content publication user 5", 5);