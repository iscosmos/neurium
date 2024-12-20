package model

import (
	"context"
	"github.com/iscosmos/neurium/pkg/orm"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryId         uint
	Name               string
	Key                string
	NodeType           int64
	NetworkMode        int64
	IngressGatewayMode int64
	DataProtocol       int64
	DataFormat         int64
	HeartBeat          int64
	Status             int64
	Vendor             string
	BrandModel         string
	Description        string
	CreateBy           string
	UpdateBy           string
}

func (p *Product) TableName() string {
	return "product"
}

type ProductModel struct {
	db *gorm.DB
}

func NewProductModel(db *gorm.DB) *ProductModel {
	return &ProductModel{db}
}

func (m *ProductModel) Create(ctx context.Context, p *Product) (int64, error) {
	if err := m.db.WithContext(ctx).Create(p).Error; err != nil {
		return 0, err
	}
	return int64(p.ID), nil
}

func (m *ProductModel) Update(ctx context.Context, p *Product) error {
	return m.db.WithContext(ctx).Save(p).Error
}

func (m *ProductModel) Delete(ctx context.Context, id int64) error {
	return m.db.WithContext(ctx).Unscoped().Delete(&Product{}, id).Error
}

func (m *ProductModel) Get(ctx context.Context, id int64) (*Product, error) {
	var p Product
	if err := m.db.WithContext(ctx).First(&p, id).Error; err != nil {
		return nil, err
	}
	return &p, nil
}

type ProductListFilter struct {
	PageNum      int64
	PageSize     int64
	Keyword      string
	NodeType     int64
	DataProtocol int64
	Status       int64
	Categories   []int64
}

func (m *ProductModel) List(ctx context.Context, filter *ProductListFilter) (int64, []*Product, error) {
	var total int64
	var result []*Product

	db := m.db.WithContext(ctx).Where("category_id in ?", filter.Categories)

	if filter.NodeType != 0 {
		db = db.Where("node_type = ?", filter.NodeType)
	}

	if filter.DataProtocol != 0 {
		db = db.Where("data_protocol = ?", filter.DataProtocol)
	}

	if filter.Status != 0 {
		db = db.Where("status = ?", filter.Status)
	}

	if len(filter.Keyword) > 0 {
		db = db.Where("name like ? or vendor like ?", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")
	}

	if err := db.Scopes(orm.Paginate(filter.PageNum, filter.PageSize)).Find(&result).Offset(-1).Limit(-1).Count(&total).Error; err != nil {
		return 0, nil, err
	}

	return total, result, nil
}
