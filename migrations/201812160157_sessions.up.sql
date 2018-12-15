CREATE TABLE IF NOT EXISTS sessions (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'internal session id',
  account_id INT UNSIGNED COMMENT 'user account id - linked to accounts table',
  ipv4 VARCHAR(16) COMMENT 'initial ip address (v4) of the session',
  ipv6 VARCHAR(40) COMMENT 'initial ip address (v6) of the session',
  source VARCHAR(16) COMMENT 'arbitrary consumer-defined code (eg. app, website)',
  device TEXT COMMENT 'device string (eg. user-agent in web/iPhone/Android/etc)',
  date_created TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT 'session created timestamp',
  last_modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'last active timestamp',
  PRIMARY KEY (id)
) ENGINE=InnoDB;
