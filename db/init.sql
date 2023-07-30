CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE CHECK (username <> ''),
  email VARCHAR(100) NOT NULL UNIQUE CHECK (email <> ''),
  password_hash CHAR(60) NOT NULL CHECK (password_hash <> '')
);

CREATE TABLE tweet (
  tweet_id SERIAL PRIMARY KEY,
  content VARCHAR(280) NOT NULL,
  created_at TIMESTAMP NOT NULL,
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
VALUES ('testuser', 'testuser@example.com', '$2a$10$0dRFAzNNqBzcoB4.xxy2ieNLRBflO6AeVsl0TXmCSbhd/XHt4WvIS');

INSERT INTO tweet (content, created_at, user_id)
VALUES ('This is a sample tweet!', '2023-07-23 12:34:56', 1);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Hello Twitter!', '2023-07-23 13:45:00', 2);

INSERT INTO tweet (content, created_at, user_id)
VALUES ('Tweeting from a different user account.', '2023-07-23 15:00:00', 3);
