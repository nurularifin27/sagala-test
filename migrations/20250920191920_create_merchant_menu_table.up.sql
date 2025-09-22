CREATE TABLE merchant_menus (
    id SERIAL PRIMARY KEY,
    merchant_id INTEGER NOT NULL REFERENCES merchants(id),
    menu_id INTEGER NOT NULL REFERENCES menus(id),
    category_id INTEGER NOT NULL REFERENCES categories(id),
    sort_order INTEGER,
    price NUMERIC(10, 2) NOT NULL DEFAULT 0,
    discount NUMERIC(10, 2) DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (merchant_id, menu_id)
);