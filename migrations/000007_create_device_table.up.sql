CREATE TABLE IF NOT EXISTS device_type (
    id VARCHAR(255) NOT NULL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS devices (
    id VARCHAR(255) NOT NULL,
    type VARCHAR(255) REFERENCES device_type(id),
    address VARCHAR(255),
    imei VARCHAR(255),
    driver VARCHAR (255) REFERENCES drivers(id)
)