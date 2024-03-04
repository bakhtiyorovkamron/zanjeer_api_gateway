CREATE TABLE IF NOT EXISTS devices_location (
    id VARCHAR(255) NOT NULL,
    imei VARCHAR(255),
    longitude TEXT[],
    latitiude TEXT[],
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)