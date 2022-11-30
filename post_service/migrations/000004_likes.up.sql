CREATE TABLE likes 
(
    id UUID PRIMARY KEY,
    user_id UUID,
    post_id UUID,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP
)