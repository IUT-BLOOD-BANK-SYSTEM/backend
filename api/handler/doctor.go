package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create doctor godoc
// @ID create_doctor
// @Router /doctor [POST]
// @Summary Create Doctor
// @Description Create Doctor
// @Tags Doctor
// @Accept json
// @Procedure json
// @Param Doctor body models.CreateDoctor true "CreateDoctorRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateDoctor(c *gin.Context) {

	var createDoctor models.CreateDoctor
	err := c.ShouldBindJSON(&createDoctor)
	if err != nil {
		h.handlerResponse(c, "error Doctor should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Doctor().Create(c.Request.Context(), &createDoctor)
	if err != nil {
		h.handlerResponse(c, "storage.Doctor.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Doctor().GetByID(c.Request.Context(), &models.DoctorPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Doctor.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Doctor resposne", http.StatusCreated, resp)
}

// GetByID doctor godoc
// @ID get_by_id_doctor
// @Router /doctor/{id} [GET]
// @Summary Get By ID Doctor
// @Description Get By ID Doctor
// @Tags Doctor
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdDoctor(c *gin.Context) {
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

	resp, err := h.strg.Doctor().GetByID(c.Request.Context(), &models.DoctorPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Doctor.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Doctor resposne", http.StatusOK, resp)
}

// GetList doctor godoc
// @ID get_list_doctor
// @Router /doctor [GET]
// @Summary Get List Doctor
// @Description Get List Doctor
// @Tags Doctor
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListDoctor(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Doctor offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Doctor limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Doctor().GetList(c.Request.Context(), &models.DoctorGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Doctor.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Doctor resposne", http.StatusOK, resp)
}

// Update doctor godoc
// @ID update_doctor
// @Router /doctor/{id} [PUT]
// @Summary Update Doctor
// @Description Update Doctor
// @Tags Doctor
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Doctor body models.UpdateDoctor true "UpdateDoctorRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateDoctor(c *gin.Context) {

	var (
		id           string = c.Param("id")
		updateDoctor models.UpdateDoctor
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}
	err := c.ShouldBindJSON(&updateDoctor)
	if err != nil {
		h.handlerResponse(c, "error Doctor should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateDoctor.Id = id
	rowsAffected, err := h.strg.Doctor().Update(c.Request.Context(), &updateDoctor)
	if err != nil {
		h.handlerResponse(c, "storage.Doctor.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Doctor.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Doctor().GetByID(c.Request.Context(), &models.DoctorPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Doctor.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Doctor resposne", http.StatusAccepted, resp)
}

// Delete doctor godoc
// @ID delete_doctor
// @Router /doctor/{id} [DELETE]
// @Summary Delete Doctor
// @Description Delete Doctor
// @Tags Doctor
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteDoctor(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Doctor().Delete(c.Request.Context(), &models.DoctorPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Doctor.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Doctor resposne", http.StatusNoContent, nil)
}
