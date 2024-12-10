package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
)

// GetList blood_request_history godoc
// @ID get_list_blood_request_history
// @Router /blood_request_history [GET]
// @Summary Get List BloodRequestHistory
// @Description Get List BloodRequestHistory
// @Tags BloodRequestHistory
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) BloodRequestHistory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list BloodRequest offset HandlerHistory", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list BloodRequest limit HandlerHistory", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.BloodRequest().GetList(c.Request.Context(), &models.BloodRequestGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.BloodRequest.get_list HandlerHistory", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list BloodRequest resposne", http.StatusOK, resp)
}
