package rest

import (
	"net/http"
	"prisca-orchestrator/domain"
	"prisca-orchestrator/internal/utils"

	"github.com/labstack/echo"
)

type PurchaseRequestHandler struct {
	purchaseRequestUseCase domain.PurchaseRequestUseCase
}

func NewPurchaseRequestHandler(e *echo.Echo, pu domain.PurchaseRequestUseCase) {
	handler := &PurchaseRequestHandler{
		purchaseRequestUseCase: pu,
	}

	e.GET("/purchase-request", handler.PurchaseRequestGet)
	e.POST("/purchase-request", handler.PurchaseRequestPost)
}

func (ph *PurchaseRequestHandler) PurchaseRequestGet(c echo.Context) error {
	result, err := ph.purchaseRequestUseCase.PurchaseRequestGet()

	if err != nil {
		return utils.ResponseHandler(c, nil, utils.InstanceMeta(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error(), nil))
	}

	return utils.ResponseHandler(c, result, utils.InstanceMeta(http.StatusOK, http.StatusText(http.StatusOK), "Success", nil))
}

func (ph *PurchaseRequestHandler) PurchaseRequestPost(c echo.Context) error {
	req := domain.PurchaseRequestPostRequest{}

	err := c.Bind(&req)

	if err != nil {
		return utils.ResponseHandler(c, nil, utils.InstanceMeta(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error(), nil))
	}

	result, err := ph.purchaseRequestUseCase.PurchaseRequestPost(req)

	if err != nil {
		return utils.ResponseHandler(c, nil, utils.InstanceMeta(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error(), nil))
	}

	return utils.ResponseHandler(c, result, utils.InstanceMeta(http.StatusOK, http.StatusText(http.StatusOK), "Success", nil))
}
