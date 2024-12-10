package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create sis_admin godoc
// @ID create_sis_admin
// @Router /sis_admin [POST]
// @Summary Create SisAdmin
// @Description Create SisAdmin
// @Tags SisAdmin
// @Accept json
// @Procedure json
// @Param SisAdmin body models.CreateSisAdmin true "CreateSisAdminRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateSisAdmin(c *gin.Context) {

	var createSisAdmin models.CreateSisAdmin
	err := c.ShouldBindJSON(&createSisAdmin)
	if err != nil {
		h.handlerResponse(c, "error SisAdmin should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.SisAdmin().Create(c.Request.Context(), &createSisAdmin)
	if err != nil {
		h.handlerResponse(c, "storage.SisAdmin.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.SisAdmin().GetByID(c.Request.Context(), &models.SisAdminPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.SisAdmin.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create SisAdmin resposne", http.StatusCreated, resp)
}

// GetByID sis_admin godoc
// @ID get_by_id_sis_admin
// @Router /sis_admin/{id} [GET]
// @Summary Get By ID SisAdmin
// @Description Get By ID SisAdmin
// @Tags SisAdmin
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdSisAdmin(c *gin.Context) {
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

	resp, err := h.strg.SisAdmin().GetByID(c.Request.Context(), &models.SisAdminPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.SisAdmin.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id SisAdmin resposne", http.StatusOK, resp)
}

// GetList sis_admin godoc
// @ID get_list_sis_admin
// @Router /sis_admin [GET]
// @Summary Get List SisAdmin
// @Description Get List SisAdmin
// @Tags SisAdmin
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListSisAdmin(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list SisAdmin offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list SisAdmin limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.SisAdmin().GetList(c.Request.Context(), &models.SisAdminGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.SisAdmin.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list SisAdmin resposne", http.StatusOK, resp)
}

// Update sis_admin godoc
// @ID update_sis_admin
// @Router /sis_admin/{id} [PUT]
// @Summary Update SisAdmin
// @Description Update SisAdmin
// @Tags SisAdmin
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param SisAdmin body models.UpdateSisAdmin true "UpdateSisAdminRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateSisAdmin(c *gin.Context) {

	var (
		id             string = c.Param("id")
		updateSisAdmin models.UpdateSisAdmin
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateSisAdmin)
	if err != nil {
		h.handlerResponse(c, "error SisAdmin should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateSisAdmin.Id = id
	rowsAffected, err := h.strg.SisAdmin().Update(c.Request.Context(), &updateSisAdmin)
	if err != nil {
		h.handlerResponse(c, "storage.SisAdmin.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.SisAdmin.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.SisAdmin().GetByID(c.Request.Context(), &models.SisAdminPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.SisAdmin.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create SisAdmin resposne", http.StatusAccepted, resp)
}

// Delete sis_admin godoc
// @ID delete_sis_admin
// @Router /sis_admin/{id} [DELETE]
// @Summary Delete SisAdmin
// @Description Delete SisAdmin
// @Tags SisAdmin
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteSisAdmin(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.SisAdmin().Delete(c.Request.Context(), &models.SisAdminPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.SisAdmin.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create SisAdmin resposne", http.StatusNoContent, nil)
}
