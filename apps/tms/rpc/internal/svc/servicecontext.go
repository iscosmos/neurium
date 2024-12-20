package svc

import (
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/config"
	"github.com/iscosmos/neurium/pkg/orm"
	"github.com/jinzhu/copier"
)

type ServiceContext struct {
	Config                    config.Config
	DB                        *orm.DB
	ProductModel              *model.ProductModel
	ProductCategoryModel      *model.ProductCategoryModel
	ProductEndpointModel      *model.ProductEndpointModel
	ProductSchemaVersionModel *model.ProductSchemaVersionModel
	ProductSchemaModuleModel  *model.ProductSchemaModuleModel
	ProductSchemaModel        *model.ProductSchemaModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化数据库连接
	var ormConfig orm.Config
	_ = copier.Copy(&ormConfig, c.DB)
	db := orm.MustNewMySQL(&ormConfig)

	return &ServiceContext{
		Config:                    c,
		DB:                        db,
		ProductModel:              model.NewProductModel(db.DB),
		ProductCategoryModel:      model.NewProductCategoryModel(db.DB),
		ProductEndpointModel:      model.NewProductEndpointModel(db.DB),
		ProductSchemaVersionModel: model.NewProductSchemaVersionModel(db.DB),
		ProductSchemaModuleModel:  model.NewProductSchemaModuleModel(db.DB),
		ProductSchemaModel:        model.NewProductSchemaModel(db.DB),
	}
}
