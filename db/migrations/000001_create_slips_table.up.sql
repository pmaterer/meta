CREATE TABLE IF NOT EXISTS slips (
    id SERIAL NOT NULL PRIMARY KEY,
    body TEXT,
    tags TEXT [],
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp
BEFORE UPDATE on slips
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();