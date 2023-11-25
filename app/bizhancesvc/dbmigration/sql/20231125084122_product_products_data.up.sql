BEGIN;

CREATE TABLE product_products(
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    name VARCHAR(100) NOT NULL,
    sku VARCHAR(100) NOT NULL,
    cost_of_goods_sold INT8 NOT NULL,
    procurement_price_currency CURRENCY_ENUM NOT NULL DEFAULT 'IDR',
    expected_lifespan_months INT8,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE product_products ADD CONSTRAINT unique_product_products_user_code unique (user_id, sku);

CREATE TABLE product_product_materials(
    id uuid PRIMARY KEY,
    product_id uuid NOT NULL REFERENCES product_products(id) ON DELETE CASCADE ON UPDATE CASCADE,
    material_item_id NOT NULL REFERENCES product_material_items(id) ON DELETE CASCADE ON UPDATE CASCADE,
    quantity_per_goods_production INT8 NOT NULL,
    cost_per_quantity INT8 NOT NULL,
    cost_per_goods_production INT8 NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE product_product_materials ADD CONSTRAINT unique_product_product_materials unique (product_id, material_item_id);



COMMIT;