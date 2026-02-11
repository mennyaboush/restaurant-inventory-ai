package repository

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
	"github.com/mennyaboush/restaurant-inventory-ai/internal/models"
)

// PostgresStore implements Repository using PostgreSQL
type PostgresStore struct {
	db *sql.DB
}

// NewPostgresStore wraps an existing *sql.DB into PostgresStore
func NewPostgresStore(db *sql.DB) *PostgresStore {
	return &PostgresStore{db: db}
}

// genID creates a fallback ID from product fields
func genID(p *models.Product) string {
	return fmt.Sprintf("%s-%d-%s", strings.ToUpper(strings.ReplaceAll(p.Brand, " ", "")), p.Size, strings.ToUpper(p.ContainerType))
}

// AddProduct creates a new product and ensures a stock row exists
func (s *PostgresStore) AddProduct(p *models.Product) (string, error) {
	if err := p.Validate(); err != nil {
		return "", err
	}
	id := p.ID
	if id == "" {
		id = genID(p)
	}

	tx, err := s.db.Begin()
	if err != nil {
		return "", err
	}
	defer func() { _ = tx.Rollback() }()

	_, err = tx.Exec(`INSERT INTO products (id, name, brand, size, container_type, box_size, price, category, is_active) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) ON CONFLICT (brand,size,container_type) DO NOTHING`, id, p.Name, p.Brand, p.Size, p.ContainerType, p.BoxSize, p.Price, p.Category, p.IsActive)
	if err != nil {
		return "", err
	}

	_, err = tx.Exec(`INSERT INTO stocks (product_id, quantity_boxes, quantity_units, min_stock, last_updated) VALUES ($1,0,0,0,CURRENT_TIMESTAMP) ON CONFLICT (product_id) DO NOTHING`, id)
	if err != nil {
		return "", err
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}
	return id, nil
}

// GetProduct retrieves a product by ID
func (s *PostgresStore) GetProduct(id string) (*models.Product, error) {
	row := s.db.QueryRow(`SELECT id, name, brand, size, container_type, box_size, price, category, is_active FROM products WHERE id=$1`, id)
	var p models.Product
	if err := row.Scan(&p.ID, &p.Name, &p.Brand, &p.Size, &p.ContainerType, &p.BoxSize, &p.Price, &p.Category, &p.IsActive); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: %s", ErrProductNotFound, id)
		}
		return nil, err
	}
	return &p, nil
}

// ListProducts returns all active products
func (s *PostgresStore) ListProducts() []*models.Product {
	rows, err := s.db.Query(`SELECT id, name, brand, size, container_type, box_size, price, category, is_active FROM products WHERE is_active = true ORDER BY name`)
	if err != nil {
		return []*models.Product{}
	}
	defer rows.Close()
	var res []*models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.Size, &p.ContainerType, &p.BoxSize, &p.Price, &p.Category, &p.IsActive); err != nil {
			continue
		}
		res = append(res, &p)
	}
	return res
}

// SearchProducts by name or brand
func (s *PostgresStore) SearchProducts(query string) []*models.Product {
	q := "%" + query + "%"
	rows, err := s.db.Query(`SELECT id, name, brand, size, container_type, box_size, price, category, is_active FROM products WHERE is_active = true AND (name ILIKE $1 OR brand ILIKE $1) ORDER BY name`, q)
	if err != nil {
		return []*models.Product{}
	}
	defer rows.Close()
	var res []*models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.Size, &p.ContainerType, &p.BoxSize, &p.Price, &p.Category, &p.IsActive); err != nil {
			continue
		}
		res = append(res, &p)
	}
	return res
}

// UpdateProduct updates an existing product
func (s *PostgresStore) UpdateProduct(p *models.Product) error {
	if err := p.Validate(); err != nil {
		return err
	}
	res, err := s.db.Exec(`UPDATE products SET name=$2, brand=$3, size=$4, container_type=$5, box_size=$6, price=$7, category=$8, is_active=$9, updated_at=CURRENT_TIMESTAMP WHERE id=$1`, p.ID, p.Name, p.Brand, p.Size, p.ContainerType, p.BoxSize, p.Price, p.Category, p.IsActive)
	if err != nil {
		return err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return fmt.Errorf("%w: %s", ErrProductNotFound, p.ID)
	}
	return nil
}

// DeleteProduct soft-deletes a product
func (s *PostgresStore) DeleteProduct(id string) error {
	res, err := s.db.Exec(`UPDATE products SET is_active = false, updated_at=CURRENT_TIMESTAMP WHERE id=$1`, id)
	if err != nil {
		return err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if cnt == 0 {
		return fmt.Errorf("%w: %s", ErrProductNotFound, id)
	}
	return nil
}

// GetStock retrieves stock for a product
func (s *PostgresStore) GetStock(productID string) (*models.Stock, error) {
	row := s.db.QueryRow(`SELECT product_id, quantity_boxes, quantity_units, min_stock, last_updated FROM stocks WHERE product_id=$1`, productID)
	var st models.Stock
	if err := row.Scan(&st.ProductID, &st.QuantityBoxes, &st.QuantityUnits, &st.MinStock, &st.LastUpdated); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("%w: %s", ErrStockNotFound, productID)
		}
		return nil, err
	}
	return &st, nil
}

// UpdateStock adjusts stock (boxes and units can be negative)
func (s *PostgresStore) UpdateStock(productID string, boxes, units int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	_, err = tx.Exec(`UPDATE stocks SET quantity_boxes = quantity_boxes + $1, quantity_units = quantity_units + $2, last_updated = CURRENT_TIMESTAMP WHERE product_id = $3`, boxes, units, productID)
	if err != nil {
		return err
	}

	var qb, qu int
	if err := tx.QueryRow(`SELECT quantity_boxes, quantity_units FROM stocks WHERE product_id=$1`, productID).Scan(&qb, &qu); err != nil {
		return err
	}
	if qb < 0 || qu < 0 {
		return fmt.Errorf("%w: would result in %d boxes, %d units", ErrInsufficientStock, qb, qu)
	}

	return tx.Commit()
}

// SetMinStock sets minimum stock threshold
func (s *PostgresStore) SetMinStock(productID string, minStock int) error {
	_, err := s.db.Exec(`UPDATE stocks SET min_stock=$1 WHERE product_id=$2`, minStock, productID)
	return err
}

// GetLowStockProducts returns active products below their min stock
func (s *PostgresStore) GetLowStockProducts() []*models.Product {
	rows, err := s.db.Query(`SELECT p.id, p.name, p.brand, p.size, p.container_type, p.box_size, p.price, p.category, p.is_active FROM products p JOIN stocks s ON p.id = s.product_id WHERE (s.quantity_boxes * COALESCE(p.box_size,0) + s.quantity_units) < s.min_stock AND p.is_active = true`)
	if err != nil {
		return []*models.Product{}
	}
	defer rows.Close()
	var res []*models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Brand, &p.Size, &p.ContainerType, &p.BoxSize, &p.Price, &p.Category, &p.IsActive); err != nil {
			continue
		}
		res = append(res, &p)
	}
	return res
}

// Ensure PostgresStore implements Repository
var _ Repository = (*PostgresStore)(nil)
