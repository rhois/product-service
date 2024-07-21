-- migrate:up
CREATE TABLE products (
            id BIGSERIAL NOT NULL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            description TEXT,
            price BIGINT,
            supplier_id BIGINT NOT NULL,
            deleted_at TIMESTAMP WITH TIME ZONE,
            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('AC','description AC',3500000,1);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('Refrigerator','description Refrigerator',2500000,2);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('HP','description HP',3000000,3);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('Laptop','description Laptop',4500000,4);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('TV','descrption TV',5000000,5);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('Macbook','descrption Macbook',24000000,3);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('Keyboard','descrption Keyboard',1000000,5);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('Ipad','descrption Ipad',6000000,3);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('Dispenser','descrption Dispenser',1500000,5);
INSERT INTO public.products ("name",description,price,supplier_id)
VALUES ('Monitor','descrption Monitor',1500000,5);

-- migrate:down
DROP TRIGGER set_timestamp on products;
DROP TABLE products;