-- migrate:up
CREATE TABLE suppliers (
            id BIGSERIAL NOT NULL PRIMARY KEY,
            name VARCHAR(100) NOT NULL,
            contact_info TEXT,
            deleted_at TIMESTAMP WITH TIME ZONE,
            created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
            updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
CREATE TRIGGER set_timestamp
    BEFORE UPDATE ON suppliers
    FOR EACH ROW
    EXECUTE PROCEDURE trigger_set_timestamp();

INSERT INTO public.suppliers ("name",contact_info)
VALUES ('Samsung','Jakarta Selatan');
INSERT INTO public.suppliers ("name",contact_info)
VALUES ('Toshiba','Jakarta Selatan');
INSERT INTO public.suppliers ("name",contact_info)
VALUES ('Apple','Jakarta Selatan');
INSERT INTO public.suppliers ("name",contact_info)
VALUES ('Microsoft','Jakarta Selatan');
INSERT INTO public.suppliers ("name",contact_info)
VALUES ('LG','Jakarta Selatan');

-- migrate:down
DROP TRIGGER set_timestamp on suppliers;
DROP TABLE suppliers;