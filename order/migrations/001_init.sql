-- +goose Up
CREATE TABLE orders
(
    order_uuid text,
    user_uuid text,
    part_uuids text[],
    total_price float,
    transaction_uuid text,
    payment_method text,
    status text
);

-- +goose Down
drop table if exists orders;