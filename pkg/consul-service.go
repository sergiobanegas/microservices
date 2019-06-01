package pkg

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"math/rand"
	"strconv"
)

type ConsulConfig struct {
	ServiceName string
	Address     string
	GRPCPort    int
	HTTPPort    int
}

func RegisterServiceWithConsul(consulConfig *ConsulConfig) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}

	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = consulConfig.ServiceName
	registration.Name = consulConfig.ServiceName
	registration.Address = consulConfig.Address
	registration.Port = consulConfig.GRPCPort
	registration.Check = new(consulapi.AgentServiceCheck)
	tags := make([]string, 1)
	tags[0] = "HTTPPort=" + strconv.Itoa(consulConfig.HTTPPort)
	registration.Tags = tags
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v/health-check", consulConfig.Address, consulConfig.HTTPPort)
	registration.Check.Interval = "5s"
	registration.Check.Timeout = "3s"
	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		log.Printf("Service %s was unabled to be registered in Consul", consulConfig.ServiceName)
	} else {
		log.Printf("Service %s registered in Consul", consulConfig.ServiceName)
	}
}

func DeregisterService(serviceId string) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatalln(err)
	}
	err = consul.Agent().ServiceDeregister(serviceId)
	if err != nil {
		log.Printf("Service %s was unabled to deregister", serviceId)

	} else {
		log.Printf("Service %s deregistered in Consul", serviceId)
	}
}

func LookupServiceWithConsul(serviceName string) (string, error) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	addrs, _, err := consul.Health().Service(serviceName, "", true, nil)
	if len(addrs) == 0 && err == nil {
		return "", errors.New(fmt.Sprintf("service ( %s ) was not found", serviceName))
	}
	if err != nil {
		return "", err
	}
	service := addrs[rand.Intn(len(addrs))].Service

	return fmt.Sprintf("%s:%v", service.Address, service.Port), nil
}
