package factory

import (
	"log"
	"sync"

	"github.com/IQ-tech/go-mapper"
	"github.com/diegoclair/bank-transfer/domain/service"
	"github.com/diegoclair/bank-transfer/infra/data"
	"github.com/diegoclair/bank-transfer/util/config"
)

type Services struct {
	Cfg         *config.EnvironmentVariables
	Mapper      mapper.Mapper
	AuthService service.AuthService
}

var (
	instance *Services
	once     sync.Once
)

//GetDomainServices to get instace of all services
func GetDomainServices() *Services {

	once.Do(func() {

		data, err := data.Connect()
		if err != nil {
			log.Fatalf("Error to connect data repositories: %v", err)
		}

		cfg := config.GetConfigEnvironment()
		svc := service.New(data, cfg)
		svm := service.NewServiceManager()
		mapper := mapper.New()
		instance = &Services{}
		instance.Cfg = cfg
		instance.Mapper = mapper
		instance.AuthService = svm.AuthService(svc)
	})

	return instance
}
