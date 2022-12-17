CREATE TABLE product_sale 
(
    id UUID PRIMARY KEY,
    user_id UUID,
    post_id UUID,
    count INT,
    price INT,
    saled_at TIMESTAMP,
    created_at TIMESTAMP,
    deleted_at TIMESTAMP    
)