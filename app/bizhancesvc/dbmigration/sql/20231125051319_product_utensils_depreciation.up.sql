BEGIN;

CREATE TABLE product_utensil_depreciations(
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE,
    name VARCHAR(100) NOT NULL,
    code VARCHAR(100) NOT NULL,
    procurement_price INT8 NOT NULL,
    procurement_price_currency CURRENCY_ENUM NOT NULL DEFAULT 'IDR',
    expected_lifespan_months INT8,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
);

ALTER TABLE product_utensil_depreciations ADD CONSTRAINT unique_product_utensil_depreciations_user_code unique (user_id, code);

COMMIT;