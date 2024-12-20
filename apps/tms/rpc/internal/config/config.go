package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	DB struct {
		DSN          string
		MaxIdleConns int64
		MaxOpenConns int64
		MaxLifetime  int64
	}
	MqttEndpoint struct {
		PropertyPost      string
		PropertyPostReply string
		PropertySet       string
		EventPost         string
		EventPostReply    string
		ServiceCall       string
		ServiceCallReply  string
		RawUp             string
		RawUpReply        string
		RawDown           string
		RawDownReply      string
		CustomTopicPrefix string
	}
	HttpEndpoint struct {
		PropertyPost string
		EventPost    string
	}
}
