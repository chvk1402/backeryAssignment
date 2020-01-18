package handlers

import (
	"fmt"
	"github.com/backery/order"
	"github.com/backery/structs"
	"github.com/backery/utils/responses"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

func Calculate(c *gin.Context) {
	req := &structs.OrderReq{}

	err := c.MustBindWith(req, binding.JSON)
	if err != nil {
		responses.ResponseWithError(c, http.StatusBadRequest, fmt.Errorf("incorrect request format"))
		return
	}

	resp, err := order.ProcessOrder(req.Code, req.Quantity)
	if err != nil {
		responses.ResponseWithError(c, http.StatusInternalServerError, fmt.Errorf("could not calcualte the order packs"))
		return
	}

	responses.ResponseWithData(c, http.StatusOK, resp)
}
