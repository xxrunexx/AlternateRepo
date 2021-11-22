package account

import (
	"movie-api/business/account"
	"movie-api/presentation/account/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	accountBusiness account.Business
}

func NewHandlerAccount(accountBusiness account.Business) *AccountHandler {
	return &AccountHandler{accountBusiness}
}

func (accHandler *AccountHandler) CreateAccount(c echo.Context) error {
	newAccount := request.ReqAccount{}

	if err := c.Bind(&newAccount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	if err, _ := accHandler.accountBusiness.CreateAccount(request.ToAccountCore(newAccount)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusAccepted, map[string]interface{}{
		"message": "Success",
		"data":    newAccount,
	})
}
