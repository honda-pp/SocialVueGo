CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE CHECK (username <> ''),
  email VARCHAR(100) NOT NULL UNIQUE CHECK (email <> ''),
  password_hash CHAR(60) NOT NULL CHECK (password_hash <> '')
);

CREATE TABLE tweet (
  tweet_id SERIAL PRIMARY KEY,
  content VARCHAR(280) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  user_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES "users" (id)
);

CREATE TABLE follow (
  follower_id INTEGER NOT NULL,
  following_id INTEGER NOT NULL,
  PRIMARY KEY (follower_id, following_id),
  FOREIGN KEY (follower_id) REFERENCES users (id),
  FOREIGN KEY (following_id) REFERENCES users (id)
);


CREATE TABLE favorite (
  favorite_id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  tweet_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES "users" (id),
  FOREIGN KEY (tweet_id) REFERENCES tweet (tweet_id)
);

CREATE TABLE retweet (
  retweet_id SERIAL PRIMARY KEY,
  user_id INTEGER NOT NULL,
  original_tweet_id INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES "users" (id),
  FOREIGN KEY (original_tweet_id) REFERENCES tweet (tweet_id)
);

INSERT INTO users (username, email, password_hash)
VALUES ('user1', 'user1@example.com', '$2a$10$1IbewfZwgFNMJiU6D4KXgeL72uMctFaDRkouNhxHdFJddJ2Fx6Dd2');

INSERT INTO users (username, email, password_hash)
VALUES ('user2', 'user2@example.com', '$2a$10$FzWZxcRDBfWJZCfBN9BQ6.Yvjl98B1mdtv7M7CbrW9kQZVpfrK.E2');

INSERT INTO users (username, email, password_hash)
VALUES ('user3', 'user3@example.com', '$2a$10$BO95bE0G17LNsDeOKgOxEeBujNJG7StOKS0dOROvagbXS70rKRdZC');

INSERT INTO users (username, email, password_hash)
VALUES ('user4', 'user4@example.com', '$2a$10$M5hzbYL8h3X7qw0AbKS2P.uPH9U6qp56OW5LkcnYLKqLJ2UNFXRim');

INSERT INTO users (username, email, password_hash)
VALUES ('user5', 'user5@example.com', '$2a$10$B5L8fuIW7jhPn6YFZB6MsuguavFVYBOBaG6q/bUtjr4.ifu8E2wza');

INSERT INTO users (username, email, password_hash)
VALUES ('user6', 'user6@example.com', '$2a$10$CzX7d65EOq53ZMPtqsz7xO6tk6r1exstMe0tBpZoVbSGmkhRrOWpa');

INSERT INTO users (username, email, password_hash)
VALUES ('user7', 'user7@example.com', '$2a$10$YL9CmzgqjU17pp4tZUGiEuwclFzKSyGYeAzrP2GwIWxSNIdiLlCJ2');

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Hello, Twitterverse!', '2023-07-24 09:15:30', 1);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Having a great day!', '2023-07-24 10:30:15', 2);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Just tweeted something cool.', '2023-07-24 11:45:45', 3);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Testing tweet functionality.', '2023-07-24 13:00:00', 4);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Final tweet of the day!', '2023-07-24 14:30:00', 5);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Another tweet from user 1.', '2023-07-25 09:00:00', 1);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Enjoying the weekend!', '2023-07-25 10:30:00', 2);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Trying out new features.', '2023-07-25 11:45:00', 3);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Feeling productive today.', '2023-07-25 13:15:00', 4);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('End of the day tweet.', '2023-07-25 14:45:00', 5);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Hello from user 1!', '2023-07-26 09:30:00', 1);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Another sunny day.', '2023-07-26 11:00:00', 2);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Sharing some thoughts.', '2023-07-26 12:30:00', 3);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Exploring new horizons.', '2023-07-26 14:00:00', 4);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Last tweet of the week!', '2023-07-26 15:30:00', 5);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('More tweets to share!', '2023-07-27 09:30:00', 6);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Exciting times ahead.', '2023-07-27 11:00:00', 7);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Exploring the city!', '2023-07-28 10:00:00', 3);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('New adventures await.', '2023-07-28 12:30:00', 4);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Chasing dreams!', '2023-07-28 15:00:00', 5);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Another day, another tweet.', '2023-07-29 09:30:00', 6);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('More tweets to share!', '2023-07-27 09:30:00', 6);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Exciting times ahead.', '2023-07-27 11:00:00', 7);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Exploring the city!', '2023-07-28 10:00:00', 3);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('New adventures await.', '2023-07-28 12:30:00', 4);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Chasing dreams!', '2023-07-28 15:00:00', 5);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Another day, another tweet.', '2023-07-29 09:30:00', 6);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Having a wonderful time!', '2023-07-29 11:00:00', 7);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Reflecting on the week.', '2023-07-30 10:00:00', 3);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Looking forward to the weekend.', '2023-07-30 12:30:00', 4);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Enjoying the sunshine!', '2023-07-30 15:00:00', 5);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Ready for new challenges.', '2023-07-31 09:30:00', 6);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Feeling inspired!', '2023-07-31 11:00:00', 7);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Making memories.', '2023-08-01 10:00:00', 3);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Exploring nature.', '2023-08-01 12:30:00', 4);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Taking a break.', '2023-08-01 15:00:00', 5);

INSERT INTO follow (follower_id, following_id)
VALUES (1, 2);
INSERT INTO follow (follower_id, following_id)
VALUES (1, 3);
INSERT INTO follow (follower_id, following_id)
VALUES (1, 4);
INSERT INTO follow (follower_id, following_id)
VALUES (2, 3);
INSERT INTO follow (follower_id, following_id)
VALUES (2, 4);
INSERT INTO follow (follower_id, following_id)
VALUES (2, 5);
INSERT INTO follow (follower_id, following_id)
VALUES (3, 4);
INSERT INTO follow (follower_id, following_id)
VALUES (3, 5);
INSERT INTO follow (follower_id, following_id)
VALUES (3, 1);
INSERT INTO follow (follower_id, following_id)
VALUES (4, 5);
INSERT INTO follow (follower_id, following_id)
VALUES (4, 1);
INSERT INTO follow (follower_id, following_id)
VALUES (4, 2);
INSERT INTO follow (follower_id, following_id)
VALUES (5, 1);
INSERT INTO follow (follower_id, following_id)
VALUES (5, 2);
INSERT INTO follow (follower_id, following_id)
VALUES (5, 3);
INSERT INTO follow (follower_id, following_id)
VALUES (6, 7);
INSERT INTO follow (follower_id, following_id)
VALUES (7, 1);
INSERT INTO follow (follower_id, following_id)
VALUES (7, 2);
INSERT INTO follow (follower_id, following_id)
VALUES (7, 3);
INSERT INTO follow (follower_id, following_id)
VALUES (5, 6);
INSERT INTO follow (follower_id, following_id)
VALUES (6, 1);
INSERT INTO follow (follower_id, following_id)
VALUES (6, 2);
INSERT INTO follow (follower_id, following_id)
VALUES (6, 3);
INSERT INTO follow (follower_id, following_id)
VALUES (6, 4);
INSERT INTO follow (follower_id, following_id)
VALUES (6, 5);
INSERT INTO follow (follower_id, following_id)
VALUES (7, 4);
INSERT INTO follow (follower_id, following_id)
VALUES (4, 7);
