CREATE TABLE brands
(
    id BIGSERIAL UNIQUE,
    name VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)