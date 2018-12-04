ALTER TABLE security
  ADD CONSTRAINT fk_security_accounts_account_id
  FOREIGN KEY (account_id) REFERENCES accounts(id)
  ON DELETE CASCADE
  ON UPDATE CASCADE;
