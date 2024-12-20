package code

import "github.com/iscosmos/neurium/pkg/code"

var (
	UnknownErr = code.NewCode(20010, "系统未知错误")
)

var (
	ProductCategoryRootErr             = code.NewCode(20000, "根目录不允许删除")
	ProductCategoryNotFoundErr         = code.NewCode(20000, "产品分类不存在")
	ProductCategoryParentIdNotFoundErr = code.NewCode(20000, "产品分类父级ID不存在")
	ProductCategoryNameExistsErr       = code.NewCode(20010, "产品分类名称已存在")
	ProductCategoryHasChildrenErr      = code.NewCode(200020, "产品分类下存在子分类")
	ProductCategoryHasProductErr       = code.NewCode(200020, "产品分类下存在产品")
)

var (
	ProductNotFoundErr                = code.NewCode(20000, "产品不存在")
	ProductNameExistsErr              = code.NewCode(20010, "产品名称已存在")
	ProductNodeTypeValidErr           = code.NewCode(20020, "产品节点类型不合法")
	ProductNetworkModeValidErr        = code.NewCode(20020, "产品联网方式不合法")
	ProductIngressGatewayModeValidErr = code.NewCode(20020, "产品接入网关协议不合法")
	ProductDataProtocolValidErr       = code.NewCode(20020, "产品数据协议不合法")
	ProductDataFormatValidErr         = code.NewCode(20020, "产品数据格式不合法")
)

var (
	ProductMqttTopicNotFoundErr = code.NewCode(20000, "产品MQTT主题不存在")
)
