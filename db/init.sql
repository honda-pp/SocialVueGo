CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE CHECK (username <> ''),
  email VARCHAR(100) NOT NULL UNIQUE CHECK (email <> ''),
  password_hash CHAR(60) NOT NULL CHECK (password_hash <> ''),
  self_introduction TEXT,
  date_of_birth DATE,
  icon_url VARCHAR(255),
  location VARCHAR(100)
);

CREATE TABLE tweet (
  tweet_id SERIAL PRIMARY KEY,
  content VARCHAR(280) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE follow (
  follower_id INTEGER NOT NULL,
  following_id INTEGER NOT NULL,
  PRIMARY KEY (follower_id, following_id),
  FOREIGN KEY (follower_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE,
  FOREIGN KEY (following_id) REFERENCES users (id) ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE favorite (
  favorite_id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  tweet_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (tweet_id) REFERENCES tweet (tweet_id)
);

CREATE TABLE retweet (
  retweet_id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  original_tweet_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id),
  FOREIGN KEY (original_tweet_id) REFERENCES tweet (tweet_id)
);

INSERT INTO users (username, email, password_hash, self_introduction, date_of_birth, icon_url, location)
VALUES
  ('user1', 'user1@example.com', '$2a$10$1IbewfZwgFNMJiU6D4KXgeL72uMctFaDRkouNhxHdFJddJ2Fx6Dd2', 'Hello, I''m user1!', '1990-01-15', 'http://localhost:8080/api/userIcon/user1_icon.png', 'Tokyo'),
  ('user2', 'user2@example.com', '$2a$10$FzWZxcRDBfWJZCfBN9BQ6.Yvjl98B1mdtv7M7CbrW9kQZVpfrK.E2', 'Nice to meet you! I''m user2.', '1985-09-22', 'http://localhost:8080/api/userIcon/user2_icon.png', 'New York'),
  ('user3', 'user3@example.com', '$2a$10$BO95bE0G17LNsDeOKgOxEeBujNJG7StOKS0dOROvagbXS70rKRdZC', 'Hello there! I''m user3.', '1992-03-10', 'http://localhost:8080/api/userIcon/user3_icon.png', 'London'),
  ('user4', 'user4@example.com', '$2a$10$M5hzbYL8h3X7qw0AbKS2P.uPH9U6qp56OW5LkcnYLKqLJ2UNFXRim', 'Greetings from user4!', '1998-06-05', 'http://localhost:8080/api/userIcon/user4_icon.png', 'Paris'),
  ('user5', 'user5@example.com', '$2a$10$B5L8fuIW7jhPn6YFZB6MsuguavFVYBOBaG6q/bUtjr4.ifu8E2wza', 'Hello everyone! I''m user5.', '1995-11-30', 'http://localhost:8080/api/userIcon/user5_icon.png', 'Sydney'),
  ('user6', 'user6@example.com', '$2a$10$CzX7d65EOq53ZMPtqsz7xO6tk6r1exstMe0tBpZoVbSGmkhRrOWpa', 'Hi, I''m user6. Nice to meet you!', '1993-07-20', 'http://localhost:8080/api/userIcon/user6_icon.png', 'Berlin'),
  ('user7', 'user7@example.com', '$2a$10$YL9CmzgqjU17pp4tZUGiEuwclFzKSyGYeAzrP2GwIWxSNIdiLlCJ2', 'Greetings, I''m user7!', '1997-04-18', 'http://localhost:8080/api/userIcon/user7_icon.png', 'Los Angeles');

INSERT INTO tweet (content, created_at, user_id)
VALUES
  ('Hello, Twitterverse!', '2023-07-24 09:15:30', 1),
  ('Having a great day!', '2023-07-24 10:30:15', 2),
  ('Just tweeted something cool.', '2023-07-24 11:45:45', 3),
  ('Testing tweet functionality.', '2023-07-24 13:00:00', 4),
  ('Final tweet of the day!', '2023-07-24 14:30:00', 5),
  ('Another tweet from user 1.', '2023-07-25 09:00:00', 1),
  ('Enjoying the weekend!', '2023-07-25 10:30:00', 2),
  ('Trying out new features.', '2023-07-25 11:45:00', 3),
  ('Feeling productive today.', '2023-07-25 13:15:00', 4),
  ('End of the day tweet.', '2023-07-25 14:45:00', 5),
  ('Hello from user 1!', '2023-07-26 09:30:00', 1),
  ('Another sunny day.', '2023-07-26 11:00:00', 2),
  ('Sharing some thoughts.', '2023-07-26 12:30:00', 3),
  ('Exploring new horizons.', '2023-07-26 14:00:00', 4),
  ('Last tweet of the week!', '2023-07-26 15:30:00', 5),
  ('More tweets to share!', '2023-07-27 09:30:00', 6),
  ('Exciting times ahead.', '2023-07-27 11:00:00', 7),
  ('Exploring the city!', '2023-07-28 10:00:00', 3),
  ('New adventures await.', '2023-07-28 12:30:00', 4),
  ('Chasing dreams!', '2023-07-28 15:00:00', 5),
  ('Another day, another tweet.', '2023-07-29 09:30:00', 6),
  ('More tweets to share!', '2023-07-27 09:30:00', 6),
  ('Exciting times ahead.', '2023-07-27 11:00:00', 7),
  ('Exploring the city!', '2023-07-28 10:00:00', 3),
  ('New adventures await.', '2023-07-28 12:30:00', 4),
  ('Chasing dreams!', '2023-07-28 15:00:00', 5),
  ('Another day, another tweet.', '2023-07-29 09:30:00', 6),
  ('Having a wonderful time!', '2023-07-29 11:00:00', 7),
  ('Reflecting on the week.', '2023-07-30 10:00:00', 3),
  ('Looking forward to the weekend.', '2023-07-30 12:30:00', 4),
  ('Enjoying the sunshine!', '2023-07-30 15:00:00', 5),
  ('Ready for new challenges.', '2023-07-31 09:30:00', 6),
  ('Feeling inspired!', '2023-07-31 11:00:00', 7),
  ('Making memories.', '2023-08-01 10:00:00', 3),
  ('Exploring nature.', '2023-08-01 12:30:00', 4),
  ('Taking a break.', '2023-08-01 15:00:00', 5);

INSERT INTO follow (follower_id, following_id)
VALUES
  (1, 2),
  (1, 3),
  (1, 4),
  (2, 3),
  (2, 4),
  (2, 5),
  (3, 4),
  (3, 5),
  (3, 1),
  (4, 5),
  (4, 1),
  (4, 2),
  (5, 1),
  (5, 2),
  (5, 3),
  (6, 7),
  (7, 1),
  (7, 2),
  (7, 3),
  (5, 6),
  (6, 1),
  (6, 2),
  (6, 3),
  (6, 4),
  (6, 5),
  (7, 4),
  (4, 7);