CREATE TABLE couriers (
    id serial NOT NULL,
    type text NOT NULL,
    regions int[],
    working_hours jsonb,
    PRIMARY KEY (id)
);

CREATE TABLE orders (
    id serial NOT NULL,
    weight int NOT NULL,
    region int NOT NULL,
    delivery_time jsonb NOT NULL,
    price int NOT NULL,
    complete_time timestamp,
    courier_id int,
    PRIMARY KEY (id),
    CONSTRAINT fk_courier FOREIGN KEY (courier_id) REFERENCES couriers(id) ON UPDATE NO ACTION ON DELETE NO ACTION
)
