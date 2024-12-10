package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Create appointment godoc
// @ID create_appointment
// @Router /appointment [POST]
// @Summary Create Appointment
// @Description Create Appointment
// @Tags Appointment
// @Accept json
// @Procedure json
// @Param Appointment body models.CreateAppointment true "CreateAppointmentRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) CreateAppointment(c *gin.Context) {

	var createAppointment models.CreateAppointment
	err := c.ShouldBindJSON(&createAppointment)
	if err != nil {
		h.handlerResponse(c, "error Appointment should bind json", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.strg.Appointment().Create(c.Request.Context(), &createAppointment)
	if err != nil {
		h.handlerResponse(c, "storage.Appointment.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.strg.Appointment().GetByID(c.Request.Context(), &models.AppointmentPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Appointment.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Appointment resposne", http.StatusCreated, resp)
}

// GetByID appointment godoc
// @ID get_by_id_appointment
// @Router /appointment/{id} [GET]
// @Summary Get By ID Appointment
// @Description Get By ID Appointment
// @Tags Appointment
// @Accept json
// @Procedure json
// @Param id path string false "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetByIdAppointment(c *gin.Context) {
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

	resp, err := h.strg.Appointment().GetByID(c.Request.Context(), &models.AppointmentPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Appointment.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get by id Appointment resposne", http.StatusOK, resp)
}

// GetList appointment godoc
// @ID get_list_appointment
// @Router /appointment [GET]
// @Summary Get List Appointment
// @Description Get List Appointment
// @Tags Appointment
// @Accept json
// @Procedure json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) GetListAppointment(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list Appointment offset", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list Appointment limit", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.strg.Appointment().GetList(c.Request.Context(), &models.AppointmentGetListRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.Appointment.get_list", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list Appointment resposne", http.StatusOK, resp)
}

// Update appointment godoc
// @ID update_appointment
// @Router /appointment/{id} [PUT]
// @Summary Update Appointment
// @Description Update Appointment
// @Tags Appointment
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Param Appointment body models.UpdateAppointment true "UpdateAppointmentRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) UpdateAppointment(c *gin.Context) {

	var (
		id                string = c.Param("id")
		updateAppointment models.UpdateAppointment
	)

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := c.ShouldBindJSON(&updateAppointment)
	if err != nil {
		h.handlerResponse(c, "error Appointment should bind json", http.StatusBadRequest, err.Error())
		return
	}
	updateAppointment.Id = id
	rowsAffected, err := h.strg.Appointment().Update(c.Request.Context(), &updateAppointment)
	if err != nil {
		h.handlerResponse(c, "storage.Appointment.update", http.StatusInternalServerError, err.Error())
		return
	}

	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.Appointment.update", http.StatusBadRequest, "now rows affected")
		return
	}

	resp, err := h.strg.Appointment().GetByID(c.Request.Context(), &models.AppointmentPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Appointment.getById", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Appointment resposne", http.StatusAccepted, resp)
}

// Delete appointment godoc
// @ID delete_appointment
// @Router /appointment/{id} [DELETE]
// @Summary Delete Appointment
// @Description Delete Appointment
// @Tags Appointment
// @Accept json
// @Procedure json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DeleteAppointment(c *gin.Context) {

	var id string = c.Param("id")

	if !helper.IsValidUUID(id) {
		h.handlerResponse(c, "is valid uuid", http.StatusBadRequest, "invalid id")
		return
	}

	err := h.strg.Appointment().Delete(c.Request.Context(), &models.AppointmentPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.Appointment.delete", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create Appointment resposne", http.StatusNoContent, nil)
}
