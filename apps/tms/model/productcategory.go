package model

import (
	"context"
	"gorm.io/gorm"
)

type ProductCategory struct {
	gorm.Model
	ParentId   int64
	Name       string
	ParentName string
	CreateBy   string
	UpdateBy   string
}

func (p *ProductCategory) TableName() string {
	return "product_category"
}

type ProductCategoryModel struct {
	db *gorm.DB
}

func NewProductCategoryModel(db *gorm.DB) *ProductCategoryModel {
	return &ProductCategoryModel{db: db}
}

func (m *ProductCategoryModel) Create(ctx context.Context, pc *ProductCategory) (int64, error) {
	if err := m.db.WithContext(ctx).Create(pc).Error; err != nil {
		return 0, err
	}
	return int64(pc.ID), nil
}

func (m *ProductCategoryModel) Update(ctx context.Context, pc *ProductCategory) error {
	return m.db.WithContext(ctx).Save(pc).Error
}

func (m *ProductCategoryModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Unscoped().Delete(&ProductCategory{}, id).Error
}

func (m *ProductCategoryModel) Get(ctx context.Context, id int64) (*ProductCategory, error) {
	var pc ProductCategory
	if err := m.db.WithContext(ctx).First(&pc, id).Error; err != nil {
		return nil, err
	}
	return &pc, nil
}

func (m *ProductCategoryModel) List(ctx context.Context) ([]*ProductCategory, error) {
	var pcs []*ProductCategory
	if err := m.db.WithContext(ctx).Find(&pcs).Error; err != nil {
		return nil, err
	}
	return pcs, nil
}

func (m *ProductCategoryModel) GetChildren(ctx context.Context, id int64) ([]*ProductCategory, error) {
	var pcs []*ProductCategory
	if err := m.db.WithContext(ctx).Where("parent_id = ?", id).Find(&pcs).Error; err != nil {
		return nil, err
	}
	return pcs, nil
}

func (m *ProductCategoryModel) UpdateParentName(ctx context.Context, ids []int64, parentName string) error {
	return m.db.WithContext(ctx).Model(&ProductCategory{}).Where("id in ?", ids).Update("parent_name", parentName).Error
}
