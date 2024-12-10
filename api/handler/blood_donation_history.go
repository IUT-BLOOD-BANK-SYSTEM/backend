package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
)

// GetList blood_donation_history godoc
// @ID get_list_blood_donation_history
// @Router /blood_donation_history [GET]
// @Summary Get List BloodDonationHistory
// @Description Get List BloodDonationHistory
// @Tags BloodDonationHistory
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListBloodDonationHistory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list BloodDonation offset in History(Handler)", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list BloodDonation limit in History(Handler)", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.BloodDonation().GetList(c.Request.Context(), &models.BloodDonationGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.BloodDonation.get_list in History(Handler)", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list BloodDonation resposne in History(Handler)", http.StatusOK, resp)
}
