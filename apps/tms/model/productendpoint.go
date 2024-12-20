package model

import (
	"context"
	"gorm.io/gorm"
)

type ProductEndpoint struct {
	gorm.Model
	ProductId   uint
	Name        string
	Protocol    int64
	Path        string
	Access      int64
	Description string
	CreateBy    string
	UpdateBy    string
}

func (p *ProductEndpoint) TableName() string {
	return "product_endpoint"
}

type ProductEndpointModel struct {
	db *gorm.DB
}

func NewProductEndpointModel(db *gorm.DB) *ProductEndpointModel {
	return &ProductEndpointModel{db}
}

func (m *ProductEndpointModel) Create(ctx context.Context, pe *ProductEndpoint) (int64, error) {
	if err := m.db.WithContext(ctx).Create(pe).Error; err != nil {
		return 0, err
	}
	return int64(pe.ID), nil
}

func (m *ProductEndpointModel) CreateMany(ctx context.Context, pes []*ProductEndpoint) error {
	return m.db.WithContext(ctx).Create(&pes).Error
}

func (m *ProductEndpointModel) Update(ctx context.Context, pe *ProductEndpoint) error {
	return m.db.WithContext(ctx).Save(pe).Error
}

func (m *ProductEndpointModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Unscoped().Delete(&ProductEndpoint{}, id).Error
}

func (m *ProductEndpointModel) Get(ctx context.Context, id int64) (*ProductEndpoint, error) {
	var pe ProductEndpoint
	if err := m.db.WithContext(ctx).First(&pe, id).Error; err != nil {
		return nil, err
	}
	return &pe, nil
}

func (m *ProductEndpointModel) List(ctx context.Context, productId int64) ([]*ProductEndpoint, error) {
	var pes []*ProductEndpoint
	if err := m.db.WithContext(ctx).Where("product_id = ?", productId).Find(&pes).Error; err != nil {
		return nil, err
	}
	return pes, nil
}
