package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Base model with UUID
type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// User represents a user in the system
type User struct {
	BaseModel
	Email         string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash  string    `gorm:"not null" json:"-"`
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Phone         string    `json:"phone"`
	Role          string    `gorm:"default:'customer'" json:"role"` // customer, admin
	EmailVerified bool      `gorm:"default:false" json:"email_verified"`
	Addresses     []Address `gorm:"foreignKey:UserID" json:"addresses,omitempty"`
	Orders        []Order   `gorm:"foreignKey:UserID" json:"orders,omitempty"`
}

// Address represents a user's address
type Address struct {
	BaseModel
	UserID        uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
	Type          string    `json:"type"` // shipping, billing
	StreetAddress string    `json:"street_address"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	PostalCode    string    `json:"postal_code"`
	Country       string    `json:"country"`
	IsDefault     bool      `gorm:"default:false" json:"is_default"`
}

// Category represents a product category
type Category struct {
	BaseModel
	Name        string     `gorm:"not null" json:"name"`
	Slug        string     `gorm:"uniqueIndex;not null" json:"slug"`
	Description string     `json:"description"`
	ParentID    *uuid.UUID `gorm:"type:uuid" json:"parent_id,omitempty"`
	Parent      *Category  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
}

// Product represents a product
type Product struct {
	BaseModel
	Name           string          `gorm:"not null" json:"name"`
	Slug           string          `gorm:"uniqueIndex;not null" json:"slug"`
	Description    string          `json:"description"`
	Price          float64         `gorm:"not null" json:"price"`
	CompareAtPrice float64         `json:"compare_at_price"`
	Cost           float64         `json:"cost"`
	SKU            string          `gorm:"uniqueIndex" json:"sku"`
	StockQuantity  int             `gorm:"default:0" json:"stock_quantity"`
	IsActive       bool            `gorm:"default:true" json:"is_active"`
	Images         []ProductImage  `gorm:"foreignKey:ProductID" json:"images,omitempty"`
	Categories     []Category      `gorm:"many2many:product_categories;" json:"categories,omitempty"`
}

// ProductImage represents a product image
type ProductImage struct {
	BaseModel
	ProductID uuid.UUID `gorm:"type:uuid;not null;index" json:"product_id"`
	URL       string    `gorm:"not null" json:"url"`
	AltText   string    `json:"alt_text"`
	Position  int       `gorm:"default:0" json:"position"`
	IsPrimary bool      `gorm:"default:false" json:"is_primary"`
}

// ProductCategory is the join table for products and categories
type ProductCategory struct {
	ProductID  uuid.UUID `gorm:"type:uuid;primaryKey"`
	CategoryID uuid.UUID `gorm:"type:uuid;primaryKey"`
}

// Cart represents a shopping cart
type Cart struct {
	BaseModel
	UserID    *uuid.UUID `gorm:"type:uuid;index" json:"user_id,omitempty"`
	SessionID string     `gorm:"index" json:"session_id"` // For guest users
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items,omitempty"`
}

// CartItem represents an item in a cart
type CartItem struct {
	BaseModel
	CartID      uuid.UUID `gorm:"type:uuid;not null;index" json:"cart_id"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null;index" json:"product_id"`
	Product     Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	PriceAtAdd  float64   `json:"price_at_add"`
}

// Order represents an order
type Order struct {
	BaseModel
	UserID              uuid.UUID   `gorm:"type:uuid;not null;index" json:"user_id"`
	OrderNumber         string      `gorm:"uniqueIndex;not null" json:"order_number"`
	Status              string      `gorm:"default:'pending'" json:"status"` // pending, processing, shipped, delivered, cancelled
	Subtotal            float64     `json:"subtotal"`
	Tax                 float64     `json:"tax"`
	Shipping            float64     `json:"shipping"`
	Total               float64     `json:"total"`
	PaymentStatus       string      `gorm:"default:'pending'" json:"payment_status"` // pending, paid, failed, refunded
	ShippingAddressID   uuid.UUID   `gorm:"type:uuid" json:"shipping_address_id"`
	BillingAddressID    uuid.UUID   `gorm:"type:uuid" json:"billing_address_id"`
	ShippingAddress     Address     `gorm:"foreignKey:ShippingAddressID" json:"shipping_address"`
	BillingAddress      Address     `gorm:"foreignKey:BillingAddressID" json:"billing_address"`
	Items               []OrderItem `gorm:"foreignKey:OrderID" json:"items,omitempty"`
	Payment             *Payment    `gorm:"foreignKey:OrderID" json:"payment,omitempty"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	BaseModel
	OrderID   uuid.UUID `gorm:"type:uuid;not null;index" json:"order_id"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index" json:"product_id"`
	Product   Product   `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `json:"price"`
	Total     float64   `json:"total"`
}

// Payment represents a payment
type Payment struct {
	BaseModel
	OrderID          uuid.UUID `gorm:"type:uuid;not null;index" json:"order_id"`
	PaymentMethod    string    `json:"payment_method"` // card, paypal, etc.
	TransactionID    string    `gorm:"uniqueIndex" json:"transaction_id"`
	Amount           float64   `json:"amount"`
	Status           string    `gorm:"default:'pending'" json:"status"` // pending, completed, failed, refunded
	ProviderResponse string    `gorm:"type:jsonb" json:"provider_response,omitempty"`
}

// BeforeCreate hook to generate UUID
func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return nil
}
