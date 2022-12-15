CREATE TABLE category 
(
    category_id BIGSERIAL,
    post_id UUID,
    name VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)