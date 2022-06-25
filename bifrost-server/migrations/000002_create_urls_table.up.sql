CREATE TABLE IF NOT EXISTS urls (
    id UUID PRIMARY KEY,
    url VARCHAR NOT NULL,
    slug VARCHAR NOT NULL,
    user_id UUID REFERENCES users(id),
    enabled BOOLEAN DEFAULT true,
    clicks BIGINT DEFAULT 0,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);