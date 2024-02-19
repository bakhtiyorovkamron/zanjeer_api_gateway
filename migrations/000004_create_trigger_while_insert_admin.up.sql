CREATE OR REPLACE FUNCTION check_admin_existence()
RETURNS TRIGGER AS 
$$
BEGIN
    IF EXISTS (SELECT * FROM "admins" WHERE login = NEW.login ) THEN
        RAISE EXCEPTION 'Admin with login % does already exists', NEW.login;
    END IF;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER admin_insert
BEFORE INSERT ON "admins"
FOR EACH ROW
EXECUTE PROCEDURE check_admin_existence();

