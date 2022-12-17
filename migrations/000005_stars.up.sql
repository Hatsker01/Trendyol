CREATE TABLE stars
(
    id UUID PRIMARY KEY,
    post_id UUID,
    user_id UUID,
    star INT,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP
)