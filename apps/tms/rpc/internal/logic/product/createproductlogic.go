package productlogic

import (
	"context"
	"github.com/iscosmos/neurium/apps/tms/model"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/code"
	"github.com/iscosmos/neurium/apps/tms/rpc/internal/config"
	"github.com/jinzhu/copier"
	"github.com/lithammer/shortuuid/v4"
	"gorm.io/gorm"
	"strings"
	"time"

	"github.com/iscosmos/neurium/apps/tms/rpc/internal/svc"
	"github.com/iscosmos/neurium/apps/tms/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateProduct 创建产品
func (l *CreateProductLogic) CreateProduct(in *pb.CreateProductReq) (*pb.CreateProductResp, error) {
	var p model.Product
	_ = copier.Copy(&p, in)

	p.Status = 2
	p.Key = shortuuid.New()

	if err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		// 创建产品
		id, err := model.NewProductModel(tx).Create(l.ctx, &p)
		if err != nil {
			return err
		}

		// 初始化通信地址
		productEndpointModel := model.NewProductEndpointModel(tx)
		if p.DataProtocol == 1 {
			if p.DataFormat == 1 {
				if err := productEndpointModel.CreateMany(l.ctx, initMqttModelEndpoint(id, p.Key, l.svcCtx.Config)); err != nil {
					return err
				}
			} else if p.DataFormat == 2 {
				if err := productEndpointModel.CreateMany(l.ctx, initMqttRawEndpoint(id, p.Key, l.svcCtx.Config)); err != nil {
					return err
				}
			}
		} else if p.DataProtocol == 2 {
			if err := productEndpointModel.CreateMany(l.ctx, initHttpEndpoint(id, p.Key, l.svcCtx.Config)); err != nil {
				return err
			}
		}

		// 初始化物模型版本(已发布)
		productSchemaVersionId, err := model.NewProductSchemaVersionModel(tx).Create(l.ctx, &model.ProductSchemaVersion{
			ProductId: uint(id),
			Version:   time.Now().UnixMilli(),
			Status:    2,
		})
		if err != nil {
			return err
		}

		// 初始化物模型默认模块
		if _, err := model.NewProductSchemaModuleModel(tx).Create(l.ctx, &model.ProductSchemaModule{
			VersionId:   uint(productSchemaVersionId),
			Name:        "默认模块",
			Identifier:  "default",
			Description: "默认模块",
			CreateBy:    "system",
		}); err != nil {
			return err
		}

		// 初始化物模型版本(草稿版本)
		productSchemaDraftVersionId, err := model.NewProductSchemaVersionModel(tx).Create(l.ctx, &model.ProductSchemaVersion{
			ProductId: uint(id),
			Version:   time.Now().UnixMilli(),
			Status:    1,
		})
		if err != nil {
			return err
		}

		// 初始化物模型草稿版本模块
		if _, err := model.NewProductSchemaModuleModel(tx).Create(l.ctx, &model.ProductSchemaModule{
			VersionId:   uint(productSchemaDraftVersionId),
			Name:        "默认模块",
			Identifier:  "default",
			Description: "默认模块",
			CreateBy:    "system",
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, code.UnknownErr
	}

	return &pb.CreateProductResp{
		Id: int64(p.ID),
	}, nil
}

func initMqttModelEndpoint(productId int64, productKey string, c config.Config) []*model.ProductEndpoint {
	mqttPropertyPost := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "属性上报",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.PropertyPost, "${productKey}", productKey, 1),
		Access:      10,
		Description: "设备属性上报",
		CreateBy:    "system",
	}
	mqttPropertyPostReply := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "属性上报",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.PropertyPostReply, "${productKey}", productKey, 1),
		Access:      11,
		Description: "云端响应属性上报",
		CreateBy:    "system",
	}
	mqttPropertySet := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "属性设置",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.PropertySet, "${productKey}", productKey, 1),
		Access:      11,
		Description: "设备属性设置",
		CreateBy:    "system",
	}
	mqttEventPost := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "事件上报",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.EventPost, "${productKey}", productKey, 1),
		Access:      10,
		Description: "设备事件上报",
		CreateBy:    "system",
	}
	mqttEventPostReply := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "事件上报",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.EventPostReply, "${productKey}", productKey, 1),
		Access:      11,
		Description: "云端响应事件上报",
		CreateBy:    "system",
	}
	mqttServiceCall := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "服务调用",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.ServiceCall, "${productKey}", productKey, 1),
		Access:      11,
		Description: "设备服务调用",
		CreateBy:    "system",
	}
	mqttServiceCallReply := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "服务调用",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.ServiceCallReply, "${productKey}", productKey, 1),
		Access:      10,
		Description: "设备端响应服务调用",
		CreateBy:    "system",
	}

	var list []*model.ProductEndpoint
	list = append(list, mqttPropertyPost)
	list = append(list, mqttPropertyPostReply)
	list = append(list, mqttPropertySet)
	list = append(list, mqttEventPost)
	list = append(list, mqttEventPostReply)
	list = append(list, mqttServiceCall)
	list = append(list, mqttServiceCallReply)

	return list
}

func initMqttRawEndpoint(productId int64, productKey string, c config.Config) []*model.ProductEndpoint {
	mqttRawUp := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "透传上行",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.RawUp, "${productKey}", productKey, 1),
		Access:      10,
		Description: "用于设备发布属性、事件和扩展信息的数据",
		CreateBy:    "system",
	}
	mqttRawUpReply := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "透传上行",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.RawUpReply, "${productKey}", productKey, 1),
		Access:      11,
		Description: "云端响应",
		CreateBy:    "system",
	}
	mqttRawDown := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "透传下行",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.RawDown, "${productKey}", productKey, 1),
		Access:      11,
		Description: "用于设备订阅获取云端设置的属性和云端调用服务的指令",
		CreateBy:    "system",
	}
	mqttRawDownReply := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "透传下行",
		Protocol:    1,
		Path:        strings.Replace(c.MqttEndpoint.RawDownReply, "${productKey}", productKey, 1),
		Access:      10,
		Description: "设备端响应",
		CreateBy:    "system",
	}

	var list []*model.ProductEndpoint
	list = append(list, mqttRawUp)
	list = append(list, mqttRawUpReply)
	list = append(list, mqttRawDown)
	list = append(list, mqttRawDownReply)

	return list
}

func initHttpEndpoint(productId int64, productKey string, c config.Config) []*model.ProductEndpoint {
	httpPropertyPost := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "属性上报",
		Protocol:    2,
		Path:        strings.Replace(c.HttpEndpoint.PropertyPost, "${productKey}", productKey, 1),
		Access:      20,
		Description: "设备属性上报",
		CreateBy:    "system",
	}
	httpEventPost := &model.ProductEndpoint{
		ProductId:   uint(productId),
		Name:        "事件上报",
		Protocol:    2,
		Path:        strings.Replace(c.HttpEndpoint.EventPost, "${productKey}", productKey, 1),
		Access:      20,
		Description: "设备事件上报",
		CreateBy:    "system",
	}

	var list []*model.ProductEndpoint
	list = append(list, httpPropertyPost)
	list = append(list, httpEventPost)

	return list
}
