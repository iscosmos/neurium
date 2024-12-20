package model

import (
	"context"
	"gorm.io/gorm"
)

type ProductSchemaVersion struct {
	gorm.Model
	ProductId uint
	Version   int64
	Status    int64
	CreateBy  string
	UpdateBy  string
}

func (p *ProductSchemaVersion) TableName() string {
	return "product_schema_version"
}

type ProductSchemaVersionModel struct {
	db *gorm.DB
}

func NewProductSchemaVersionModel(db *gorm.DB) *ProductSchemaVersionModel {
	return &ProductSchemaVersionModel{db}
}

func (m *ProductSchemaVersionModel) Create(ctx context.Context, psv *ProductSchemaVersion) (int64, error) {
	if err := m.db.WithContext(ctx).Create(psv).Error; err != nil {
		return 0, err
	}
	return int64(psv.ID), nil
}

func (m *ProductSchemaVersionModel) Update(ctx context.Context, psv *ProductSchemaVersion) error {
	return m.db.WithContext(ctx).Save(psv).Error
}

func (m *ProductSchemaVersionModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Unscoped().Delete(&ProductSchemaVersion{}, id).Error
}

func (m *ProductSchemaVersionModel) Get(ctx context.Context, id int64) (*ProductSchemaVersion, error) {
	var psv ProductSchemaVersion
	if err := m.db.WithContext(ctx).First(&psv, id).Error; err != nil {
		return nil, err
	}
	return &psv, nil
}

func (m *ProductSchemaVersionModel) List(ctx context.Context, productId, status int64) ([]*ProductEndpoint, error) {
	var pes []*ProductEndpoint
	if err := m.db.WithContext(ctx).Where("product_id = ? and status = ?", productId, status).Find(&pes).Error; err != nil {
		return nil, err
	}
	return pes, nil
}

func (m *ProductSchemaVersionModel) GetWith(ctx context.Context, productId, status int64) (*ProductSchemaVersion, error) {
	var psv ProductSchemaVersion
	if err := m.db.WithContext(ctx).Where("product_id = ? and status = ?", productId, status).First(&psv).Error; err != nil {
		return nil, err
	}
	return &psv, nil
}
