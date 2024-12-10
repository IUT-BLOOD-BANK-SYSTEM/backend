package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create blood_donation godoc
// @ID create_blood_donation
// @Router /blood_donation [POST]
// @Summary Create BloodDonation
// @Description Create BloodDonation
// @Tags BloodDonation
// @Accept json
// @Procedure json
// @Param BloodDonation body models.CreateBloodDonation true "CreateBloodDonationRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateBloodDonation(c *gin.Context) {

	var createBloodDonation models.CreateBloodDonation
	err := c.ShouldBindJSON(&createBloodDonation)
	if err != nil {
		h.handlerResponse(c, "error BloodDonation should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.BloodDonation().Create(c.Request.Context(), &createBloodDonation)
	if err != nil {
		h.handlerResponse(c, "storage.BloodDonation.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.BloodDonation().GetByID(c.Request.Context(), &models.BloodDonationPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodDonation.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodDonation resposne", http.StatusCreated, resp)
}

// GetByID blood_donation godoc
// @ID get_by_id_blood_donation
// @Router /blood_donation/{id} [GET]
// @Summary Get By ID BloodDonation
// @Description Get By ID BloodDonation
// @Tags BloodDonation
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdBloodDonation(c *gin.Context) {
	var id string = c.Param("id")

	// Here We Check id from Token
	// val, exist := c.Get("Auth")
	// if !exist {
	// 	h.handlerResponse(c, "Here", http.StatusInternalServerError, nil)
	// 	return
	// }

	// ToolData := val.(helper.TokenInfo)
	// if len(ToolData) > 0 {
	// 	id = ToolData.ToolID
	// } else {
	// 	id = c.Param("id")
	// }

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	resp, err := h.strg.BloodDonation().GetByID(c.Request.Context(), &models.BloodDonationPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodDonation.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id BloodDonation resposne", http.StatusOK, resp)
}

// GetList blood_donation godoc
// @ID get_list_blood_donation
// @Router /blood_donation [GET]
// @Summary Get List BloodDonation
// @Description Get List BloodDonation
// @Tags BloodDonation
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListBloodDonation(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list BloodDonation offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list BloodDonation limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.BloodDonation().GetList(c.Request.Context(), &models.BloodDonationGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.BloodDonation.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list BloodDonation resposne", http.StatusOK, resp)
}

// Update blood_donation godoc
// @ID update_blood_donation
// @Router /blood_donation/{id} [PUT]
// @Summary Update BloodDonation
// @Description Update BloodDonation
// @Tags BloodDonation
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param BloodDonation body models.UpdateBloodDonation true "UpdateBloodDonationRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateBloodDonation(c *gin.Context) {

	var (
		id                  string = c.Param("id")
		updateBloodDonation models.UpdateBloodDonation
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateBloodDonation)
	if err != nil {
		h.handlerResponse(c, "error BloodDonation should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateBloodDonation.Id = id
	rowsAffected, err := h.strg.BloodDonation().Update(c.Request.Context(), &updateBloodDonation)
	if err != nil {
		h.handlerResponse(c, "storage.BloodDonation.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.BloodDonation.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.BloodDonation().GetByID(c.Request.Context(), &models.BloodDonationPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodDonation.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodDonation resposne", http.StatusAccepted, resp)
}

// Delete blood_donation godoc
// @ID delete_blood_donation
// @Router /blood_donation/{id} [DELETE]
// @Summary Delete BloodDonation
// @Description Delete BloodDonation
// @Tags BloodDonation
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteBloodDonation(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.BloodDonation().Delete(c.Request.Context(), &models.BloodDonationPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodDonation.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodDonation resposne", http.StatusNoContent, nil)
}
