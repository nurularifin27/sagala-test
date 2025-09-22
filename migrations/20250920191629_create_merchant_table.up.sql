CREATE TABLE merchants (
    id SERIAL PRIMARY KEY,
    branch_id INTEGER NOT NULL REFERENCES branches(id),
    brand_id INTEGER NOT NULL REFERENCES brands(id),
    channel_id INTEGER NOT NULL REFERENCES channels(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (branch_id, brand_id, channel_id)
);