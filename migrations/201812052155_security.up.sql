CREATE TABLE IF NOT EXISTS security (
  account_id INT UNSIGNED COMMENT 'user account id - linked to accounts table',
  password VARCHAR(255) COMMENT 'hashed password of user'
) ENGINE=InnoDB;
