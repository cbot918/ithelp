-- Database CRUD

-- Create
INSERT INTO users(name, email, password, discount) VALUES ('yale','yale@gmail.com', '12345', 1);
INSERT INTO users(name, email, password, discount) VALUES ('node','node@gmail.com', '12345', 0.9);
INSERT INTO users(name, email, password, discount) VALUES ('hi','hi@gmail.com', '12345', 0.85);
INSERT INTO users(name, email, password, discount) VALUES ('delete','delete@gmail.com', '12345', 0.1);

-- Read
SELECT * FROM users;

-- Update
UPDATE users SET discount = 0.85 WHERE name='hi';

-- Delete
DELETE FROM users WHERE name = 'delete';

-- Read
SELECT * FROM users;