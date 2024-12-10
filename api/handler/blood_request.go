package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create blood_request godoc
// @ID create_blood_request
// @Router /blood_request [POST]
// @Summary Create BloodRequest
// @Description Create BloodRequest
// @Tags BloodRequest
// @Accept json
// @Procedure json
// @Param BloodRequest body models.CreateBloodRequest true "CreateBloodRequestRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateBloodRequest(c *gin.Context) {

	var createBloodRequest models.CreateBloodRequest
	err := c.ShouldBindJSON(&createBloodRequest)
	if err != nil {
		h.handlerResponse(c, "error BloodRequest should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.BloodRequest().Create(c.Request.Context(), &createBloodRequest)
	if err != nil {
		h.handlerResponse(c, "storage.BloodRequest.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.BloodRequest().GetByID(c.Request.Context(), &models.BloodRequestPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodRequest.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodRequest resposne", http.StatusCreated, resp)
}

// GetByID blood_request godoc
// @ID get_by_id_blood_blood_request
// @Router /blood_request/{id} [GET]
// @Summary Get By ID BloodRequest
// @Description Get By ID BloodRequest
// @Tags BloodRequest
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdBloodRequest(c *gin.Context) {
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

	resp, err := h.strg.BloodRequest().GetByID(c.Request.Context(), &models.BloodRequestPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodRequest.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id BloodRequest resposne", http.StatusOK, resp)
}

// GetList blood_request godoc
// @ID get_list_blood_request
// @Router /blood_request [GET]
// @Summary Get List BloodRequest
// @Description Get List BloodRequest
// @Tags BloodRequest
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListBloodRequest(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list BloodRequest offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list BloodRequest limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.BloodRequest().GetList(c.Request.Context(), &models.BloodRequestGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.BloodRequest.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list BloodRequest resposne", http.StatusOK, resp)
}

// Update blood_request godoc
// @ID update_blood_request
// @Router /blood_request/{id} [PUT]
// @Summary Update BloodRequest
// @Description Update BloodRequest
// @Tags BloodRequest
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param BloodRequest body models.UpdateBloodRequest true "UpdateBloodRequestRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateBloodRequest(c *gin.Context) {

	var (
		id                 string = c.Param("id")
		updateBloodRequest models.UpdateBloodRequest
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateBloodRequest)
	if err != nil {
		h.handlerResponse(c, "error BloodRequest should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateBloodRequest.Id = id
	rowsAffected, err := h.strg.BloodRequest().Update(c.Request.Context(), &updateBloodRequest)
	if err != nil {
		h.handlerResponse(c, "storage.BloodRequest.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.BloodRequest.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.BloodRequest().GetByID(c.Request.Context(), &models.BloodRequestPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodRequest.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodRequest resposne", http.StatusAccepted, resp)
}

// Delete blood_request godoc
// @ID delete_blood_blood_request
// @Router /blood_request/{id} [DELETE]
// @Summary Delete BloodRequest
// @Description Delete BloodRequest
// @Tags BloodRequest
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteBloodRequest(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.BloodRequest().Delete(c.Request.Context(), &models.BloodRequestPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodRequest.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodRequest resposne", http.StatusNoContent, nil)
}
