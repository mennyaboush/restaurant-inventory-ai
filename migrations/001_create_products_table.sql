-- +migrate Up
CREATE TABLE products (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    brand VARCHAR(100) NOT NULL,
    size INTEGER NOT NULL CHECK (size > 0),
    container_type VARCHAR(50) NOT NULL,
    box_size INTEGER DEFAULT 0 CHECK (box_size >= 0),
    price DECIMAL(10,2) NOT NULL CHECK (price >= 0),
    category VARCHAR(50) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create index for unique constraint: Brand + Size + ContainerType
CREATE UNIQUE INDEX idx_products_unique ON products (brand, size, container_type);

-- Create index for category queries
CREATE INDEX idx_products_category ON products (category);

-- Create index for active products
CREATE INDEX idx_products_active ON products (is_active);

-- +migrate Down
DROP TABLE IF EXISTS products;