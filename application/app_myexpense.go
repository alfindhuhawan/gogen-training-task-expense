package application

import (
	"your/path/project/domain_myexpense/controller/restapi"
	"your/path/project/domain_myexpense/gateway/sqlitelocal"
	"your/path/project/domain_myexpense/usecase/getallexpense"
	"your/path/project/domain_myexpense/usecase/runexpensecreate"
	"your/path/project/shared/driver"
	"your/path/project/shared/infrastructure/config"
	"your/path/project/shared/infrastructure/logger"
	"your/path/project/shared/infrastructure/server"
	"your/path/project/shared/infrastructure/util"
)

type myexpense struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c myexpense) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewMyexpense() func() driver.RegistryContract {
	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData("myexpense", appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandlerDefault(log, appData, cfg)

		datasource := sqlitelocal.NewGateway(log, appData, cfg)

		return &myexpense{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:                    log,
				Config:                 cfg,
				Router:                 httpHandler.Router,
				RunExpenseCreateInport: runexpensecreate.NewUsecase(datasource),
				GetAllExpenseInport:    getallexpense.NewUsecase(datasource),
			},
		}

	}
}
