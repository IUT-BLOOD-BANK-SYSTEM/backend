package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create blood_inventory godoc
// @ID create_blood_inventory
// @Router /blood_inventory [POST]
// @Summary Create BloodInventory
// @Description Create BloodInventory
// @Tags BloodInventory
// @Accept json
// @Procedure json
// @Param BloodInventory body models.CreateBloodInventory true "CreateBloodInventoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateBloodInventory(c *gin.Context) {

	var createBloodInventory models.CreateBloodInventory
	err := c.ShouldBindJSON(&createBloodInventory)
	if err != nil {
		h.handlerResponse(c, "error BloodInventory should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.BloodInventory().Create(c.Request.Context(), &createBloodInventory)
	if err != nil {
		h.handlerResponse(c, "storage.BloodInventory.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.BloodInventory().GetByID(c.Request.Context(), &models.BloodInventoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodInventory.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodInventory resposne", http.StatusCreated, resp)
}

// GetByID blood_inventory godoc
// @ID get_by_id_blood_inventory
// @Router /blood_inventory/{id} [GET]
// @Summary Get By ID BloodInventory
// @Description Get By ID BloodInventory
// @Tags BloodInventory
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdBloodInventory(c *gin.Context) {
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

	resp, err := h.strg.BloodInventory().GetByID(c.Request.Context(), &models.BloodInventoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodInventory.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id BloodInventory resposne", http.StatusOK, resp)
}

// GetList blood_inventory godoc
// @ID get_list_blood_inventory
// @Router /blood_inventory [GET]
// @Summary Get List BloodInventory
// @Description Get List BloodInventory
// @Tags BloodInventory
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListBloodInventory(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list BloodInventory offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list BloodInventory limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.BloodInventory().GetList(c.Request.Context(), &models.BloodInventoryGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.BloodInventory.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list BloodInventory resposne", http.StatusOK, resp)
}

// Update blood_inventory godoc
// @ID update_blood_inventory
// @Router /blood_inventory/{id} [PUT]
// @Summary Update BloodInventory
// @Description Update BloodInventory
// @Tags BloodInventory
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param BloodInventory body models.UpdateBloodInventory true "UpdateBloodInventoryRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateBloodInventory(c *gin.Context) {

	var (
		id                   string = c.Param("id")
		updateBloodInventory models.UpdateBloodInventory
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateBloodInventory)
	if err != nil {
		h.handlerResponse(c, "error BloodInventory should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateBloodInventory.Id = id
	rowsAffected, err := h.strg.BloodInventory().Update(c.Request.Context(), &updateBloodInventory)
	if err != nil {
		h.handlerResponse(c, "storage.BloodInventory.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.BloodInventory.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.BloodInventory().GetByID(c.Request.Context(), &models.BloodInventoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodInventory.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodInventory resposne", http.StatusAccepted, resp)
}

// Delete blood_inventory godoc
// @ID delete_blood_inventory
// @Router /blood_inventory/{id} [DELETE]
// @Summary Delete BloodInventory
// @Description Delete BloodInventory
// @Tags BloodInventory
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteBloodInventory(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.BloodInventory().Delete(c.Request.Context(), &models.BloodInventoryPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.BloodInventory.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create BloodInventory resposne", http.StatusNoContent, nil)
}
