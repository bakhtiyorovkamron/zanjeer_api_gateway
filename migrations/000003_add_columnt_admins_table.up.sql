CREATE TYPE user_type AS ENUM ('superadmin', 'admin');
ALTER TABLE admins ADD COLUMN type user_type NOT NULL;
