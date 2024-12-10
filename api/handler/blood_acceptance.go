package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create blood_acceptance godoc
// @ID create_blood_acceptance
// @Router /blood_acceptance [POST]
// @Summary Create BloodAcceptance
// @Description Create BloodAcceptance
// @Tags BloodAcceptance
// @Accept json
// @Procedure json
// @Param BloodAcceptance body models.CreateBloodAcceptance true "CreateBloodAcceptanceRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateBloodAcceptance(c *gin.Context) {

	var createBloodAcceptance models.CreateBloodAcceptance
	err := c.ShouldBindJSON(&createBloodAcceptance)
	if err != nil {
		h.handlerResponse(c, "error BloodAcceptance should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.BloodAcceptance().Create(c.Request.Context(), &createBloodAcceptance)
	if err != nil {
		h.handlerResponse(c, "storage.BloodAcceptance.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.BloodAcceptance().GetByID(c.Request.Context(), &models.BloodAcceptancePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodAcceptance.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodAcceptance resposne", http.StatusCreated, resp)
}

// GetByID blood_acceptance godoc
// @ID get_by_id_blood_acceptance
// @Router /blood_acceptance/{id} [GET]
// @Summary Get By ID BloodAcceptance
// @Description Get By ID BloodAcceptance
// @Tags BloodAcceptance
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdBloodAcceptance(c *gin.Context) {
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

	resp, err := h.strg.BloodAcceptance().GetByID(c.Request.Context(), &models.BloodAcceptancePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodAcceptance.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id BloodAcceptance resposne", http.StatusOK, resp)
}

// GetList blood_acceptance godoc
// @ID get_list_blood_acceptance
// @Router /blood_acceptance [GET]
// @Summary Get List BloodAcceptance
// @Description Get List BloodAcceptance
// @Tags BloodAcceptance
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListBloodAcceptance(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list BloodAcceptance offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list BloodAcceptance limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.BloodAcceptance().GetList(c.Request.Context(), &models.BloodAcceptanceGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.BloodAcceptance.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list BloodAcceptance resposne", http.StatusOK, resp)
}

// Update blood_acceptance godoc
// @ID update_blood_acceptance
// @Router /blood_acceptance/{id} [PUT]
// @Summary Update BloodAcceptance
// @Description Update BloodAcceptance
// @Tags BloodAcceptance
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param BloodAcceptance body models.UpdateBloodAcceptance true "UpdateBloodAcceptanceRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateBloodAcceptance(c *gin.Context) {

	var (
		id                    string = c.Param("id")
		updateBloodAcceptance models.UpdateBloodAcceptance
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateBloodAcceptance)
	if err != nil {
		h.handlerResponse(c, "error BloodAcceptance should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateBloodAcceptance.Id = id
	rowsAffected, err := h.strg.BloodAcceptance().Update(c.Request.Context(), &updateBloodAcceptance)
	if err != nil {
		h.handlerResponse(c, "storage.BloodAcceptance.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.BloodAcceptance.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.BloodAcceptance().GetByID(c.Request.Context(), &models.BloodAcceptancePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodAcceptance.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodAcceptance resposne", http.StatusAccepted, resp)
}

// Delete blood_acceptance godoc
// @ID delete_blood_acceptance
// @Router /blood_acceptance/{id} [DELETE]
// @Summary Delete BloodAcceptance
// @Description Delete BloodAcceptance
// @Tags BloodAcceptance
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteBloodAcceptance(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.BloodAcceptance().Delete(c.Request.Context(), &models.BloodAcceptancePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodAcceptance.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodAcceptance resposne", http.StatusNoContent, nil)
}
