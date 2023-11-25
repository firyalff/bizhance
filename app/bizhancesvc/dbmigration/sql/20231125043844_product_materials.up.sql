BEGIN;

CREATE TABLE product_material_brands(
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE product_material_brands ADD CONSTRAINT unique_product_material_brands_user_name unique (user_id, name);

CREATE TABLE product_material_categories(
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE product_material_categories ADD CONSTRAINT unique_product_material_categories_user_name unique (user_id, name);

CREATE TYPE CURRENCY_ENUM AS ENUM ('IDR');

CREATE TABLE product_material_items(
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    brand_id uuid NOT NULL REFERENCES product_material_brands(id) ON DELETE CASCADE ON UPDATE CASCADE,
    category_id uuid NOT NULL REFERENCES product_material_categories(id) ON DELETE CASCADE ON UPDATE CASCADE,
    name VARCHAR(150) NOT NULL,
    code VARCHAR(100) NOT NULL UNIQUE,
    procurement_price INT8 NOT NULL,
    procurement_price_currency CURRENCY_ENUM NOT NULL DEFAULT 'IDR',
    material_procurement_uom VARCHAR(100),
    goods_production_uom VARCHAR(100),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE product_material_items ADD CONSTRAINT unique_product_material_items_user_code unique (user_id, code);

COMMIT;