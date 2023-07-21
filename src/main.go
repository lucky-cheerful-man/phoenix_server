package main

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	goMirco "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	pb "github.com/lucky-cheerful-man/phoenix_apis/protobuf3.pb/user_info_manage"
	"github.com/lucky-cheerful-man/phoenix_server/src/config"
	"github.com/lucky-cheerful-man/phoenix_server/src/gmysql"
	"github.com/lucky-cheerful-man/phoenix_server/src/gredis"
	"github.com/lucky-cheerful-man/phoenix_server/src/log"
	"github.com/lucky-cheerful-man/phoenix_server/src/service"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs(config.ReferGlobalConfig().ServerSetting.RegisterAddress),
	)
	// 初始化服务
	srv := goMirco.NewService(
		goMirco.Name(config.ReferGlobalConfig().ServerSetting.RegisterServerName),
		goMirco.Version(config.ReferGlobalConfig().ServerSetting.RegisterServerVersion),
		goMirco.Registry(etcdReg),
	)

	err := pb.RegisterUserServiceHandler(srv.Server(), &service.UserService{DB: gmysql.Setup(),
		Cache: gredis.Setup()})
	if err != nil {
		log.Error("RegisterUserServiceHandler failed, err:%s", err)
		return
	}

	err = srv.Run()
	if err != nil {
		log.Error("run failed, err:%s", err)
	}
}