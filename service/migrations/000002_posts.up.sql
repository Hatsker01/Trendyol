CREATE TABLE posts
(
    id UUID PRIMARY KEY,
    title VARCHAR(150),
    description TEXT,
    body TEXT,
    author_id uuid,
    stars INT,
    rating INT,
    created_at TIMESTAMP,
    upated_at TIMESTAMP,
    deleted_at TIMESTAMP
)