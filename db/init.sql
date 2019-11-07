CREATE DATABASE closer_event;
CREATE DATABASE grafana;

-- Create events table to the closer_event db
\c closer_event;
CREATE TABLE IF NOT EXISTS events (
    id UUID NOT NULL PRIMARY KEY,
    event VARCHAR(255) NOT NULL,
    source_id VARCHAR(255) NOT NULL,
    source_type VARCHAR(255) NOT NULL,
    data JSONB DEFAULT '{}'::jsonb,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL
);
CREATE INDEX events_event_source_id ON events (event, source_id);
CREATE INDEX events_timestamp ON events (timestamp);

-- Create user role
CREATE ROLE readonly_role;

-- Grant access to existing tables
GRANT USAGE ON SCHEMA public TO readonly_role;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly_role;

-- Grant access to future tables
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO readonly_role;

-- Create a grafana user
CREATE USER grafana_user; --WITH PASSWORD 6r4f4n4_p455w0rd;
GRANT readonly_role TO grafana_user;

-- Grant all access on the grafana database
GRANT ALL PRIVILEGES ON DATABASE grafana TO grafana_user;

CREATE OR REPLACE FUNCTION TIMESTAMPTZ_TRUNC(_unit TEXT, _date TIMESTAMP WITH TIME ZONE, _timezone TEXT) RETURNS TIMESTAMP WITH TIME ZONE
    LANGUAGE plpgsql
AS
$$
BEGIN RETURN DATE_TRUNC(_unit, _date AT TIME ZONE _timezone)::TIMESTAMP WITHOUT TIME ZONE AT TIME ZONE _timezone; END
$$;
