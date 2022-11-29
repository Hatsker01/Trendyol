CREATE TABLE posts
(
    id UUID PRIMARY KEY,
    title VARCHAR(150),
    description TEXT,
    body TEXT,
    author_id uuid,
    stars INT,
    rating INT,
    price INT,
    product_type VARCHAR,
    size INT[],
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)