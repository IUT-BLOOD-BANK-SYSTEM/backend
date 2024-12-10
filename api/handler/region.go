package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create region godoc
// @ID create_region
// @Router /region [POST]
// @Summary Create Region
// @Description Create Region
// @Tags Region
// @Accept json
// @Procedure json
// @Param Region body models.CreateRegion true "CreateRegionRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateRegion(c *gin.Context) {

	var createRegion models.CreateRegion
	err := c.ShouldBindJSON(&createRegion)
	if err != nil {
		h.handlerResponse(c, "error Region should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Region().Create(c.Request.Context(), &createRegion)
	if err != nil {
		h.handlerResponse(c, "storage.Region.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Region().GetByID(c.Request.Context(), &models.RegionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Region.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Region resposne", http.StatusCreated, resp)
}

// GetByID region godoc
// @ID get_by_id_region
// @Router /region/{id} [GET]
// @Summary Get By ID Region
// @Description Get By ID Region
// @Tags Region
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdRegion(c *gin.Context) {
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

	resp, err := h.strg.Region().GetByID(c.Request.Context(), &models.RegionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Region.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Region resposne", http.StatusOK, resp)
}

// GetList region godoc
// @ID get_list_region
// @Router /region [GET]
// @Summary Get List Region
// @Description Get List Region
// @Tags Region
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListRegion(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Region offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Region limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Region().GetList(c.Request.Context(), &models.RegionGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Region.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Region resposne", http.StatusOK, resp)
}

// Update region godoc
// @ID update_region
// @Router /region/{id} [PUT]
// @Summary Update Region
// @Description Update Region
// @Tags Region
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Region body models.UpdateRegion true "UpdateRegionRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateRegion(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateRegion models.UpdateRegion
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateRegion)
	if err != nil {
		h.handlerResponse(c, "error Region should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateRegion.Id = id
	rowsAffected, err := h.strg.Region().Update(c.Request.Context(), &updateRegion)
	if err != nil {
		h.handlerResponse(c, "storage.Region.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Region.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Region().GetByID(c.Request.Context(), &models.RegionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Region.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Region resposne", http.StatusAccepted, resp)
}

// Delete region godoc
// @ID delete_region
// @Router /region/{id} [DELETE]
// @Summary Delete Region
// @Description Delete Region
// @Tags Region
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteRegion(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Region().Delete(c.Request.Context(), &models.RegionPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Region.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Region resposne", http.StatusNoContent, nil)
}
