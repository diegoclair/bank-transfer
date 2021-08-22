package factory

import (
	"log"
	"sync"

	"github.com/IQ-tech/go-crypto-layer/datacrypto"
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

		instance = &Services{}
		cfg := config.GetConfigEnvironment()
		cipher := datacrypto.NewAESECB(datacrypto.AES256, cfg.MySQL.CryptoKey)
		svc := service.New(data, cfg, cipher)
		svm := service.NewServiceManager()

		instance.Cfg = cfg
		instance.Mapper = mapper.New()
		instance.AuthService = svm.AuthService(svc)
	})

	return instance
}
