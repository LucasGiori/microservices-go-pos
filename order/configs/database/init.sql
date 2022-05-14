CREATE TABLE "order"
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    customer_id UUID NOT NULL,
    customer_name VARCHAR(100) NOT NULL,
    date_time timestamp NOT NULL,
    status VARCHAR(100) NOT NULL,
    total numeric(19, 2) NOT NULL
);

CREATE TABLE order_item
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    product_name VARCHAR(100) NOT NULL,
    quantity int NOT NULL,
    unit_value numeric(19, 2) NOT NULL,
    total numeric(19, 2) NOT NULL,
    CONSTRAINT order_item_order_id_fk FOREIGN KEY (order_id) REFERENCES "order"(id)
 );