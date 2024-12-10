package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create head_nurse godoc
// @ID create_head_nurse
// @Router /head_nurse [POST]
// @Summary Create HeadNurse
// @Description Create HeadNurse
// @Tags HeadNurse
// @Accept json
// @Procedure json
// @Param HeadNurse body models.CreateHeadNurse true "CreateHeadNurseRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateHeadNurse(c *gin.Context) {

	var createHeadNurse models.CreateHeadNurse
	err := c.ShouldBindJSON(&createHeadNurse)
	if err != nil {
		h.handlerResponse(c, "error HeadNurse should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.HeadNurse().Create(c.Request.Context(), &createHeadNurse)
	if err != nil {
		h.handlerResponse(c, "storage.HeadNurse.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.HeadNurse().GetByID(c.Request.Context(), &models.HeadNursePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.HeadNurse.getById", http.StatusInternalServerError, err.Error())
		return
	}

	hospital, err := h.strg.Hospital().GetByID(c.Request.Context(), &models.HospitalPrimaryKey{Id: resp.Hospital.Id})
	if err != nil {
		h.handlerResponse(c, "storage.HeadNurse(hospital get by id).getById", http.StatusInternalServerError, err.Error())
		return
	}
	if hospital.Id == resp.Hospital.Id {
		_, err := h.strg.Hospital().Update(c, &models.UpdateHospital{
			Name:        hospital.Name,
			Id:          hospital.Id,
			RegionID:    hospital.Region.Id,
			Address:     hospital.Address,
			PhoneNumber: hospital.PhoneNumber,
			HeadNurseID: id,
		})
		if err != nil {
			h.handlerResponse(c, "error HeadNurse should bind json", http.StatusBadRequest, err.Error())
			return
		}
	}
	h.handlerResponse(c, "create HeadNurse resposne", http.StatusCreated, resp)
}

// GetByID head_nurse godoc
// @ID get_by_id_head_nurse
// @Router /head_nurse/{id} [GET]
// @Summary Get By ID HeadNurse
// @Description Get By ID HeadNurse
// @Tags HeadNurse
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdHeadNurse(c *gin.Context) {
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

	resp, err := h.strg.HeadNurse().GetByID(c.Request.Context(), &models.HeadNursePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.HeadNurse.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id HeadNurse resposne", http.StatusOK, resp)
}

// GetList head_nurse godoc
// @ID get_list_head_nurse
// @Router /head_nurse [GET]
// @Summary Get List HeadNurse
// @Description Get List HeadNurse
// @Tags HeadNurse
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListHeadNurse(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list HeadNurse offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list HeadNurse limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.HeadNurse().GetList(c.Request.Context(), &models.HeadNurseGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.HeadNurse.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list HeadNurse resposne", http.StatusOK, resp)
}

// Update head_nurse godoc
// @ID update_head_nurse
// @Router /head_nurse/{id} [PUT]
// @Summary Update HeadNurse
// @Description Update HeadNurse
// @Tags HeadNurse
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param HeadNurse body models.UpdateHeadNurse true "UpdateHeadNurseRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateHeadNurse(c *gin.Context) {

	var (
		id              string = c.Param("id")
		updateHeadNurse models.UpdateHeadNurse
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}
	err := c.ShouldBindJSON(&updateHeadNurse)
	if err != nil {
		h.handlerResponse(c, "error HeadNurse should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateHeadNurse.Id = id
	rowsAffected, err := h.strg.HeadNurse().Update(c.Request.Context(), &updateHeadNurse)
	if err != nil {
		h.handlerResponse(c, "storage.HeadNurse.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.HeadNurse.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.HeadNurse().GetByID(c.Request.Context(), &models.HeadNursePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.HeadNurse.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create HeadNurse resposne", http.StatusAccepted, resp)
}

// Delete head_nurse godoc
// @ID delete_head_nurse
// @Router /head_nurse/{id} [DELETE]
// @Summary Delete HeadNurse
// @Description Delete HeadNurse
// @Tags HeadNurse
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteHeadNurse(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.HeadNurse().Delete(c.Request.Context(), &models.HeadNursePrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.HeadNurse.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create HeadNurse resposne", http.StatusNoContent, nil)
}
