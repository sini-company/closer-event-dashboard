CREATE DATABASE closer_event;
CREATE DATABASE grafana;

-- Create a group
CREATE ROLE readonly_role;

-- Grant access to existing tables
GRANT USAGE ON SCHEMA public TO readonly_role;
GRANT SELECT ON ALL TABLES IN SCHEMA public TO readonly_role;

-- Grant access to future tables
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO readonly_role;

-- Create a grafana user
CREATE USER grafana_user;
GRANT readonly_role TO grafana_user;

-- Grant all access on the grafana database
GRANT ALL PRIVILEGES ON DATABASE grafana TO grafana_user;

-- Create events table to the closer_event db
\c closer_event;
CREATE TABLE IF NOT EXISTS events (
    id UUID NOT NULL PRIMARY KEY,
    event VARCHAR(255) NOT NULL,
    sourceId VARCHAR(255) NOT NULL,
    sourceType VARCHAR(255) NOT NULL,
    data JSONB DEFAULT '{}'::jsonb,
    timestamp TIMESTAMP WITH TIME ZONE NOT NULL
);
