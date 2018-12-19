CREATE TRIGGER accounts_before_insert
  BEFORE INSERT ON accounts
  FOR EACH ROW
  SET new.uuid = UUID();
