package entities

import "time"

type Category struct {
	ID        uint64     `gorm:"primaryKey" json:"id"`
	Name      string     `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at_at" json:"deleted_at"`
	Products  []Product  `gorm:"foreignKey:category_id" json:"products,omitempty"`
}
