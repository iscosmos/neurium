CREATE DATABASE tms;
USE tms;

# 产品表
CREATE TABLE `product`
(
    `id`                   BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `category_id`          BIGINT UNSIGNED NOT NULL COMMENT '产品分类ID',
    `name`                 VARCHAR(32)     NOT NULL COMMENT '产品名称',
    `key`                  VARCHAR(32)     NOT NULL COMMENT '产品密钥',
    `node_type`            TINYINT         NOT NULL DEFAULT 0 COMMENT '节点类型 1:直连设备 2:网关设备 3:网关子设备',
    `network_mode`         TINYINT         NOT NULL DEFAULT 0 COMMENT '联网方式 1:以太网 2:蜂窝(2G/3G/4G/5G) 3:WiFi 4:其它 | 仅有直连设备和网关设备存在',
    `ingress_gateway_mode` TINYINT         NOT NULL DEFAULT 0 COMMENT '接入网关协议 1:自定义 2:Modbus 3:OPCUA 4: Zigbee 5:BLE | 仅有子设备存在',
    `data_protocol`        TINYINT         NOT NULL DEFAULT 0 COMMENT '数据协议 1:MQTT 2:HTTP | 仅有直连设备和网关设备存在',
    `data_format`          TINYINT         NOT NULL DEFAULT 0 COMMENT '数据格式 1:物模型 2:自定义/透传',
    `heart_beat`           INT             NOT NULL DEFAULT 0 COMMENT '最大心跳间隔 | 判断设备数据是否异常',
    `status`               TINYINT         NOT NULL DEFAULT 0 COMMENT '状态 1:已发布 2:未发布',
    `vendor`               VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '产品厂商',
    `brand_model`          VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '品牌型号',
    `description`          VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '产品描述',
    `create_by`            VARCHAR(32)     NOT NULL COMMENT '创建用户',
    `update_by`            VARCHAR(32)     NOT NULL DEFAULT '' COMMENT '更新用户',
    `created_at`           DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`           DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`           DATETIME(3)     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_product_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='产品表';

# 产品分类表
CREATE TABLE `product_category`
(
    `id`          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `parent_id`   BIGINT UNSIGNED NOT NULL COMMENT '父级ID',
    `name`        VARCHAR(32)     NOT NULL COMMENT '分类名称',
    `parent_name` VARCHAR(255)    NOT NULL COMMENT '上级名称',
    `create_by`   VARCHAR(32)     NOT NULL COMMENT '创建用户',
    `update_by`   VARCHAR(32)     NOT NULL DEFAULT '' COMMENT '更新用户',
    `created_at`  DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`  DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`  DATETIME(3)     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_product_category_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='产品分类表';

# 产品通信地址
CREATE TABLE `product_endpoint`
(
    `id`                   BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `product_id`           BIGINT UNSIGNED NOT NULL COMMENT '产品ID',
    `name`                 VARCHAR(32)     NOT NULL COMMENT '名称',
    `protocol`             TINYINT         NOT NULL DEFAULT 0 COMMENT '协议 1:MQTT 2:HTTP',
    `path`                 VARCHAR(255)    NOT NULL COMMENT '通信地址',
    `access`               TINYINT         NOT NULL DEFAULT 0 COMMENT '权限 10:发布 11:订阅 12:发布和订阅 20:POST',
    `description`          VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '描述',
    `create_by`            VARCHAR(32)     NOT NULL COMMENT '创建用户',
    `update_by`            VARCHAR(32)     NOT NULL DEFAULT '' COMMENT '更新用户',
    `created_at`           DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`           DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`           DATETIME(3)     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_product_endpoint_deleted_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT ='产品通信地址';

# 产品物模型模块版本
CREATE TABLE `product_schema_version`
(
    `id`                    BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `product_id`            BIGINT UNSIGNED NOT NULL COMMENT '产品ID',
    `version`               BIGINT UNSIGNED NOT NULL COMMENT '版本号 发布的时间戳',
    `status`                TINYINT         NOT NULL DEFAULT 0 COMMENT '状态 1:草稿状态 2:发布状态 3:已归档',
    `create_by`             VARCHAR(32)     NOT NULL COMMENT '创建用户',
    `update_by`             VARCHAR(32)     NOT NULL DEFAULT '' COMMENT '更新用户',
    `created_at`            DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`            DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`            DATETIME(3)     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_product_schema_version_deleted_at` (`deleted_at`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4
 COLLATE = utf8mb4_unicode_ci COMMENT ='产品物模型版本表';

# 产品物模型模块
CREATE TABLE `product_schema_module`
(
    `id`                    BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `version_id`            BIGINT UNSIGNED NOT NULL COMMENT '版本ID',
    `name`                  VARCHAR(32)     NOT NULL COMMENT '名称',
    `identifier`            VARCHAR(32)     NOT NULL COMMENT '标识符',
    `description`           VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '描述',
    `create_by`             VARCHAR(32)     NOT NULL COMMENT '创建用户',
    `update_by`             VARCHAR(32)     NOT NULL DEFAULT '' COMMENT '更新用户',
    `created_at`            DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`            DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`            DATETIME(3)     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_product_schema_module_deleted_at` (`deleted_at`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4
 COLLATE = utf8mb4_unicode_ci COMMENT ='产品物模型模块表';

# 产品物模型
CREATE TABLE `product_schema`
(
    `id`                    BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `module_id`             BIGINT UNSIGNED NOT NULL COMMENT '模块ID',
    `parent_id`             BIGINT UNSIGNED NOT NULL COMMENT '父ID',
    `type`                  TINYINT         NOT NULL DEFAULT 0 COMMENT '功能类型 1:属性 2:事件 3:服务 4:参数',
    `name`                  VARCHAR(32)     NOT NULL COMMENT '名称 示例:环境温度',
    `identifier`            VARCHAR(32)     NOT NULL COMMENT '标识符 示例:temperature',
    `specification`         TEXT            NOT NULL COMMENT '数据定义 示例:{"type":"float", "min": 0.0, "max": 99.0}',
    `data_type`             TINYINT         NOT NULL DEFAULT 0 COMMENT '数据类型 1:整数 2:浮点数 3:布尔 4:字符串 5:枚举 6:结构体 7:数组',
    `access_type`           TINYINT         NOT NULL DEFAULT 0 COMMENT '读写类型 1:读写 2:只读',
    `event_type`            TINYINT         NOT NULL DEFAULT 0 COMMENT '事件类型 1:信息 2:告警 3:故障',
    `service_call_type`     TINYINT         NOT NULL DEFAULT 0 COMMENT '服务调用方式 1:异步 2:同步',
    `parameter_type`        TINYINT         NOT NULL DEFAULT 0 COMMENT '参数类型 1:输入参数 2:输出参数',
    `description`           VARCHAR(255)    NOT NULL DEFAULT '' COMMENT '描述',
    `create_by`             VARCHAR(32)     NOT NULL COMMENT '创建用户',
    `update_by`             VARCHAR(32)     NOT NULL DEFAULT '' COMMENT '更新用户',
    `created_at`            DATETIME(3)     NULL COMMENT '创建时间',
    `updated_at`            DATETIME(3)     NULL COMMENT '更新时间',
    `deleted_at`            DATETIME(3)     NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_product_schema_deleted_at` (`deleted_at`)
)ENGINE = InnoDB
 DEFAULT CHARSET = utf8mb4
 COLLATE = utf8mb4_unicode_ci COMMENT ='产品物模型表';