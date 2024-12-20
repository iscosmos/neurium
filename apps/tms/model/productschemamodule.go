package model

import (
	"context"
	"gorm.io/gorm"
)

type ProductSchemaModule struct {
	gorm.Model
	VersionId   uint
	Name        string
	Identifier  string
	Description string
	CreateBy    string
	UpdateBy    string
}

func (p *ProductSchemaModule) TableName() string {
	return "product_schema_module"
}

type ProductSchemaModuleModel struct {
	db *gorm.DB
}

func NewProductSchemaModuleModel(db *gorm.DB) *ProductSchemaModuleModel {
	return &ProductSchemaModuleModel{db: db}
}

func (m *ProductSchemaModuleModel) Create(ctx context.Context, psm *ProductSchemaModule) (int64, error) {
	if err := m.db.WithContext(ctx).Create(psm).Error; err != nil {
		return 0, err
	}
	return int64(psm.ID), nil
}

func (m *ProductSchemaModuleModel) Update(ctx context.Context, psm *ProductSchemaModule) error {
	return m.db.WithContext(ctx).Save(psm).Error
}

func (m *ProductSchemaModuleModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Unscoped().Delete(&ProductSchemaModule{}, id).Error
}

func (m *ProductSchemaModuleModel) Get(ctx context.Context, id int64) (*ProductSchemaModule, error) {
	var p ProductSchemaModule
	if err := m.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (m *ProductSchemaModuleModel) List(ctx context.Context, productId int64) ([]*ProductSchemaModule, error) {
	var psms []*ProductSchemaModule

	if err := m.db.WithContext(ctx).Where("version_id = ?", productId).Find(&psms).Error; err != nil {
		return nil, err
	}

	return psms, nil
}
