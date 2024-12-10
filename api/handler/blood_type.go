package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create blood_type godoc
// @ID create_blood_type
// @Router /blood_type [POST]
// @Summary Create BloodType
// @Description Create BloodType
// @Tags BloodType
// @Accept json
// @Procedure json
// @Param BloodType body models.CreateBloodType true "CreateBloodTypeRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateBloodType(c *gin.Context) {

	var createBloodType models.CreateBloodType
	err := c.ShouldBindJSON(&createBloodType)
	if err != nil {
		h.handlerResponse(c, "error BloodType should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.BloodType().Create(c.Request.Context(), &createBloodType)
	if err != nil {
		h.handlerResponse(c, "storage.BloodType.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.BloodType().GetByID(c.Request.Context(), &models.BloodTypePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodType.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodType resposne", http.StatusCreated, resp)
}

// GetByID blood_type godoc
// @ID get_by_id_blood_type
// @Router /blood_type/{id} [GET]
// @Summary Get By ID BloodType
// @Description Get By ID BloodType
// @Tags BloodType
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdBloodType(c *gin.Context) {
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

	resp, err := h.strg.BloodType().GetByID(c.Request.Context(), &models.BloodTypePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodType.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id BloodType resposne", http.StatusOK, resp)
}

// GetList blood_type godoc
// @ID get_list_blood_type
// @Router /blood_type [GET]
// @Summary Get List BloodType
// @Description Get List BloodType
// @Tags BloodType
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListBloodType(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list BloodType offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list BloodType limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.BloodType().GetList(c.Request.Context(), &models.BloodTypeGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.BloodType.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list BloodType resposne", http.StatusOK, resp)
}

// Update blood_type godoc
// @ID update_blood_type
// @Router /blood_type/{id} [PUT]
// @Summary Update BloodType
// @Description Update BloodType
// @Tags BloodType
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param BloodType body models.UpdateBloodType true "UpdateBloodTypeRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateBloodType(c *gin.Context) {

	var (
		id              string = c.Param("id")
		updateBloodType models.UpdateBloodType
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateBloodType)
	if err != nil {
		h.handlerResponse(c, "error BloodType should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateBloodType.Id = id
	rowsAffected, err := h.strg.BloodType().Update(c.Request.Context(), &updateBloodType)
	if err != nil {
		h.handlerResponse(c, "storage.BloodType.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.BloodType.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.BloodType().GetByID(c.Request.Context(), &models.BloodTypePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodType.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodType resposne", http.StatusAccepted, resp)
}

// Delete blood_type godoc
// @ID delete_blood_type
// @Router /blood_type/{id} [DELETE]
// @Summary Delete BloodType
// @Description Delete BloodType
// @Tags BloodType
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteBloodType(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.BloodType().Delete(c.Request.Context(), &models.BloodTypePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodType.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodType resposne", http.StatusNoContent, nil)
}
