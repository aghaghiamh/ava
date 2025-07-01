# User Table
For saving time, instead of using the migrator, I just applied the Create Query directly on the Mysql.
    CREATE TABLE users ( id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(256) NOT NULL, phone_number VARCHAR(50), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP );