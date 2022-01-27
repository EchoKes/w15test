CREATE USER 'user'@'%' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'user'@'%';

-- Sample Database Setup
CREATE DATABASE sample_db;

USE sample_db;

-- CREATE TABLE Sample
-- ( 
--     id INT AUTO_INCREMENT PRIMARY KEY,
--     sample VARCHAR(100),
--     datetime DATETIME,
--     boolValue TINYINT(0) DEFAULT 0
-- );

-- INSERT INTO Sample() 
-- VALUES
