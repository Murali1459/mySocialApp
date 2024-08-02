CREATE USER 'admin'@'172.18.0.1' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'admin'@'172.18.0.1';
FLUSH PRIVILEGES;

-- Create database and use it
CREATE DATABASE social_media;
USE social_media;

-- Create users table
CREATE TABLE Users (
  id INT AUTO_INCREMENT,
  username VARCHAR(255) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,
  profile_picture_url VARCHAR(255),
  bio TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);

-- Create index on username and email
CREATE INDEX idx_username ON Users (username);
CREATE INDEX idx_email ON Users (email);

-- Create posts table
CREATE TABLE Posts (
  id INT AUTO_INCREMENT,
  user_id INT NOT NULL,
  content TEXT NOT NULL,
  image_url VARCHAR(255),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) REFERENCES Users(id)
);

-- Create index on user_id
CREATE INDEX idx_user_id ON Posts (user_id);

-- Create comments table
CREATE TABLE Comments (
  id INT AUTO_INCREMENT,
  post_id INT NOT NULL,
  user_id INT NOT NULL,
  content TEXT NOT NULL,

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (post_id) REFERENCES Posts(id),
  FOREIGN KEY (user_id) REFERENCES Users(id)
);

-- Create index on post_id and user_id
CREATE INDEX idx_post_id ON Comments (post_id);

CREATE INDEX idx_user_id ON Comments (user_id);

-- Create follows table
CREATE TABLE Follows (
  id INT AUTO_INCREMENT,
  follower_id INT NOT NULL,
  followee_id INT NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_on TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  PRIMARY KEY (id),
  FOREIGN KEY (follower_id) REFERENCES Users(id),

  FOREIGN KEY (followee_id) REFERENCES Users(id)
);

-- Create index on follower_id and followee_id
CREATE INDEX idx_follower_id ON Follows (follower_id);
CREATE INDEX idx_followee_id ON Follows (followee_id);
