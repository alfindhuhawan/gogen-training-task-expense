package restapi

import (
	"github.com/gin-gonic/gin"

	"your/path/project/domain_myexpense/usecase/getallexpense"
	"your/path/project/domain_myexpense/usecase/runexpensecreate"
	"your/path/project/shared/infrastructure/config"
	"your/path/project/shared/infrastructure/logger"
)

type Controller struct {
	Router                 gin.IRouter
	Config                 *config.Config
	Log                    logger.Logger
	RunExpenseCreateInport runexpensecreate.Inport
	GetAllExpenseInport    getallexpense.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.POST("/runexpensecreate", r.authorized(), r.runExpenseCreateHandler(r.RunExpenseCreateInport))
	r.Router.GET("/getallexpense", r.authorized(), r.getAllExpenseHandler(r.GetAllExpenseInport))
}
