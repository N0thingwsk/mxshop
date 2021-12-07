package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/satori/go.uuid"
	"mxshop/usersrv/config"
)

func ConsulRegister(port int, addr string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		return err
	}
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", addr, port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "5s",
	}
	reg := new(api.AgentServiceRegistration)
	reg.Name = config.UserConfig.UserServiceName
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	reg.ID = serviceId
	reg.Port = port
	reg.Tags = []string{"微服务", "user", "srv"}
	reg.Address = addr
	reg.Check = check
	err = client.Agent().ServiceRegister(reg)
	return nil
}
