package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"your/path/project/domain_myexpense/model/entity"
	"your/path/project/domain_myexpense/usecase/getallexpense"
	"your/path/project/shared/infrastructure/logger"
	"your/path/project/shared/infrastructure/util"
	"your/path/project/shared/model/payload"
)

// getAllExpenseHandler ...
func (r *Controller) getAllExpenseHandler(inputPort getallexpense.Inport) gin.HandlerFunc {

	type request struct {
		Page int `form:"page,omitempty,default=0"`
		Size int `form:"size,omitempty,default=0"`
	}

	type response struct {
		Count int               `json:"count"`
		Items []*entity.Expense `json:"items"`
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.Bind(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req getallexpense.InportRequest
		req.Page = jsonReq.Page
		req.Size = jsonReq.Size

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		jsonRes.Count = res.Count
		jsonRes.Items = res.Items

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
