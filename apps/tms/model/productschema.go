package model

import (
	"context"
	"github.com/iscosmos/neurium/pkg/orm"
	"gorm.io/gorm"
)

type ProductSchema struct {
	gorm.Model
	ModuleId        uint
	ParentId        uint
	Type            int64
	Name            string
	Identifier      string
	Specification   string
	DataType        int64
	AccessType      int64
	EventType       int64
	ServiceCallType int64
	ParameterType   int64
	Description     string
	CreateBy        string
	UpdateBy        string
}

func (p *ProductSchema) TableName() string {
	return "product_schema"
}

type ProductSchemaModel struct {
	db *gorm.DB
}

func NewProductSchemaModel(db *gorm.DB) *ProductSchemaModel {
	return &ProductSchemaModel{db: db}
}

func (m *ProductSchemaModel) Create(ctx context.Context, p *ProductSchema) (int64, error) {
	if err := m.db.WithContext(ctx).Create(p).Error; err != nil {
		return 0, err
	}
	return int64(p.ID), nil
}

func (m *ProductSchemaModel) Update(ctx context.Context, p *ProductSchema) error {
	return m.db.WithContext(ctx).Save(p).Error
}

func (m *ProductSchemaModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Unscoped().Delete(&ProductSchema{}, id).Error
}

func (m *ProductSchemaModel) Get(ctx context.Context, id int64) (*ProductSchema, error) {
	var p ProductSchema
	if err := m.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

func (m *ProductSchemaModel) List(ctx context.Context, pageNum, pageSize, moduleId int64, keyword string) (int64, []*ProductSchema, error) {
	var total int64
	var result []*ProductSchema

	db := m.db.WithContext(ctx).Order("type").Where("module_id = ?", moduleId)

	if len(keyword) > 0 {
		db = db.Where("name like ? or identifier like ?", "%"+keyword+"%", "%"+keyword+"%")
	}

	if err := db.Scopes(orm.Paginate(pageNum, pageSize)).Find(&result).Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		return 0, nil, err
	}

	return total, result, nil
}
