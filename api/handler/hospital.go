package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create hospital godoc
// @ID create_hospital
// @Router /hospital [POST]
// @Summary Create Hospital
// @Description Create Hospital
// @Tags Hospital
// @Accept json
// @Procedure json
// @Param Hospital body models.CreateHospital true "CreateHospitalRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateHospital(c *gin.Context) {
	var createHospital models.CreateHospital

	checker, err := h.strg.Hospital().GetList(c, &models.HospitalGetListRequest{})
	if err != nil {
		h.handlerResponse(c, "get list Hospital offset in create hospital", http.StatusBadRequest, "invalid offset")
		return
	}

	for _, a := range checker.Hospitals {
		if a.HeadNurse.Id == createHospital.HeadNurseID {

			h.handlerResponse(c, "Head nurse is already assigned to another hospital", http.StatusBadRequest, nil)
			return

		}
	}

	err = c.ShouldBindJSON(&createHospital)
	if err != nil {
		h.handlerResponse(c, "error Hospital should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Hospital().Create(c.Request.Context(), &createHospital)
	if err != nil {
		h.handlerResponse(c, "storage.Hospital.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Hospital().GetByID(c.Request.Context(), &models.HospitalPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Hospital.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Hospital resposne", http.StatusCreated, resp)
}

// GetByID hospital godoc
// @ID get_by_id_hospital
// @Router /hospital/{id} [GET]
// @Summary Get By ID Hospital
// @Description Get By ID Hospital
// @Tags Hospital
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdHospital(c *gin.Context) {
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

	resp, err := h.strg.Hospital().GetByID(c.Request.Context(), &models.HospitalPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Hospital.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Hospital resposne", http.StatusOK, resp)
}

// GetList hospital godoc
// @ID get_list_hospital
// @Router /hospital [GET]
// @Summary Get List Hospital
// @Description Get List Hospital
// @Tags Hospital
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListHospital(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Hospital offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Hospital limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Hospital().GetList(c.Request.Context(), &models.HospitalGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Hospital.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Hospital resposne", http.StatusOK, resp)
}

// Update hospital godoc
// @ID update_hospital
// @Router /hospital/{id} [PUT]
// @Summary Update Hospital
// @Description Update Hospital
// @Tags Hospital
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Hospital body models.UpdateHospital true "UpdateHospitalRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateHospital(c *gin.Context) {

	var (
		id             string = c.Param("id")
		updateHospital models.UpdateHospital
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateHospital)
	if err != nil {
		h.handlerResponse(c, "error Hospital should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateHospital.Id = id
	rowsAffected, err := h.strg.Hospital().Update(c.Request.Context(), &updateHospital)
	if err != nil {
		h.handlerResponse(c, "storage.Hospital.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Hospital.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Hospital().GetByID(c.Request.Context(), &models.HospitalPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Hospital.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Hospital resposne", http.StatusAccepted, resp)
}

// Delete hospital godoc
// @ID delete_hospital
// @Router /hospital/{id} [DELETE]
// @Summary Delete Hospital
// @Description Delete Hospital
// @Tags Hospital
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteHospital(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Hospital().Delete(c.Request.Context(), &models.HospitalPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Hospital.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Hospital resposne", http.StatusNoContent, nil)
}
