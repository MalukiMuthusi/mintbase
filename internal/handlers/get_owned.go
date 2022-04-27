package handlers

import (
	"net/http"

	"github.com/MalukiMuthusi/mintbase/internal/adapters"
	"github.com/MalukiMuthusi/mintbase/internal/models"
	"github.com/gin-gonic/gin"
)

type OwnedHandler struct{}

func (h *OwnedHandler) Handle(c *gin.Context) {
	var userID models.UserIDParameter

	if err := c.ShouldBindUri(&userID); err != nil {
		e := models.BasicError{
			Code:    models.InvalidUserIdParam.String(),
			Message: "provide a valid user parameter",
		}

		c.JSON(http.StatusUnprocessableEntity, e)
		return
	}

	b, err := adapters.GetOwnedByUser(userID)
	if err != nil {
		e := models.BasicError{
			Code:    models.InvalidUserIdParam.String(),
			Message: "failed to fetch data",
		}

		c.JSON(http.StatusInternalServerError, e)
		return
	}

	// var res map[string]interface{}

	c.JSON(http.StatusOK, b)
}
