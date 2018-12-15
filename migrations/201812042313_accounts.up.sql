CREATE TABLE IF NOT EXISTS accounts (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'internal account id',
  uuid VARCHAR(37) NOT NULL UNIQUE COMMENT 'public facing id of this user',
  email VARCHAR(255) UNIQUE COMMENT 'email address of the user',
  username VARCHAR(64) UNIQUE COMMENT 'username of the user if applicable',
  date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'timestamp user was created',
  last_modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'timestamp this row was last updated',
  PRIMARY KEY (id)
) ENGINE=InnoDB;
