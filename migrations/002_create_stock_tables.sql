-- +migrate Up
CREATE TABLE stocks (
    product_id VARCHAR(50) PRIMARY KEY,
    quantity_boxes INTEGER DEFAULT 0 CHECK (quantity_boxes >= 0),
    quantity_units INTEGER DEFAULT 0 CHECK (quantity_units >= 0),
    min_stock INTEGER DEFAULT 0 CHECK (min_stock >= 0),
    last_updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE stock_movements (
    id VARCHAR(50) PRIMARY KEY,
    product_id VARCHAR(50) NOT NULL REFERENCES products(id),
    type VARCHAR(20) NOT NULL,
    boxes INTEGER DEFAULT 0,
    units INTEGER DEFAULT 0,
    performed_by VARCHAR(100) NOT NULL,
    reported_by VARCHAR(100) NOT NULL,
    reason TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_stocks_min ON stocks (min_stock);
CREATE INDEX idx_movements_product ON stock_movements (product_id);

-- +migrate Down
DROP TABLE IF EXISTS stock_movements;
DROP TABLE IF EXISTS stocks;
