package handler

import (
	"app/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		value := c.GetHeader("Authorization")

		info_other, err := helper.ParseClaimsForOther(value, h.cfg.SecretKey)

		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}

		info_tadqiqotchi, err := helper.ParseClaimsForTadqiqotchi(value, h.cfg.SecretKey)

		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}
		info_oqituvchi, err := helper.ParseClaimsForOqituvchi(value, h.cfg.SecretKey)

		if err != nil {
			c.AbortWithError(http.StatusForbidden, err)
			return
		}
		c.Set("Auth", info_other)
		c.Set("Auth", info_tadqiqotchi)
		c.Set("Auth", info_oqituvchi)
		c.Next()
	}
}
