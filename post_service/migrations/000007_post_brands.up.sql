CREATE TABLE post_brand 
(
    id BIGSERIAL PRIMARY KEY UNIQUE,
    brand_id UUID,
    post_id UUID,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP,
)