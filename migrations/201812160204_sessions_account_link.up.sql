ALTER TABLE sessions
  ADD CONSTRAINT fk_sessions_accounts_account_id
  FOREIGN KEY (account_id) REFERENCES accounts(id)
  ON DELETE CASCADE
  ON UPDATE CASCADE;
