CREATE OR REPLACE FUNCTION store_deleted_time_of_driver()
RETURNS TRIGGER AS 
$$
BEGIN 
    UPDATE city SET deleted_at = NOW() WHERE id=OLD.id;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER store_deleted_time
  AFTER UPDATE
  ON "city"
  FOR EACH ROW
  EXECUTE PROCEDURE store_deleted_time_of_driver();