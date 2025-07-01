A MySQL driver is required to enable Go's standard database/sql package to communicate with a MySQL database.</br></br>

### Database Schema
</br>
This project uses a users table to store user information. To set up the database for the first time, execute the following SQL statement:</br>

<p>SQL</br>
    CREATE TABLE users (</br>
        id INT AUTO_INCREMENT PRIMARY KEY,</br>
        name VARCHAR(256) NOT NULL,</br>
        phone_number VARCHAR(50) NULL,</br>
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP</br>
    );</br>
</p>


<!-- 
TODO: 
    - docker-compose environment variables of db (critical security concerns!)

-->