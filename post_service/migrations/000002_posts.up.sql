CREATE TABLE posts
(
    id UUID PRIMARY KEY,
    title VARCHAR(150),
    description TEXT,
    body TEXT,
    author_id uuid,
    stars INT,
    rating INT,
    price VARCHAR(100),
    product_type VARCHAR(50),
    size VARCHAR(50)[],
    color  VARCHAR,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)