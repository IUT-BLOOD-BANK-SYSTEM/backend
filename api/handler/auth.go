package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"app/api/models"
	"app/pkg/helper"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Login
// @Description Login
// @Tags Login
// @Accept json
// @Procedure json
// @Param login body models.LoginInfo true "LoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Login(c *gin.Context) {
	var login models.LoginInfo

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create user", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: login.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "User does not exist", http.StatusBadRequest, "User does not exist")
			return
		}
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(resp)

	if resp.Password != login.Password {
		h.handlerResponse(c, "Wrong password", http.StatusBadRequest, "Wrong password")
		return
	}

	token, err := helper.GenerateJWT(map[string]interface{}{
		"user_id": resp.Id,
	}, time.Hour*360, h.cfg.SecretKey)

	h.handlerResponse(c, "token", http.StatusCreated, token)
}

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Register
// @Description Register
// @Tags Register
// @Accept json
// @Procedure json
// @Param register body models.CreateUser true "CreateUserRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) Register(c *gin.Context) {

	var createUser models.CreateUser
	var id string
	err := c.ShouldBindJSON(&createUser)
	if err != nil {
		h.handlerResponse(c, "error user should bind json", http.StatusBadRequest, err.Error())
		return
	}

	if len(createUser.Password) < 7 {
		h.handlerResponse(c, "Password should inculude more than 7 elements", http.StatusBadRequest, errors.New("Password len should inculude more than 8 elements"))
		return
	}

	resp, err := h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Email: createUser.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			id, err = h.strg.User().Create(context.Background(), &createUser)
			if err != nil {
				h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			h.handlerResponse(c, "User already exist", http.StatusInternalServerError, err.Error())
			return
		}
	} else if err == nil {
		h.handlerResponse(c, "User already exist", http.StatusBadRequest, nil)
		return
	}
	resp, _ = h.strg.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})

	h.handlerResponse(c, "create user resposne", http.StatusCreated, resp)
}

// HeadNurseLogin godoc
// @ID head_nurse_login
// @Router /head_nurse_login [POST]
// @Summary HeadNurseLogin
// @Description HeadNurseLogin
// @Tags HeadNurseLogin
// @Accept json
// @Procedure json
// @Param head_nurse_login body models.LoginInfo true "HeadNurseLoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) HeadNurseLogin(c *gin.Context) {
	var login models.LoginInfo

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create HeadNurse", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.HeadNurse().GetByID(context.Background(), &models.HeadNursePrimaryKey{Email: login.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "HeadNurse does not exist", http.StatusBadRequest, "HeadNurse does not exist")
			return
		}
		h.handlerResponse(c, "storage.HeadNurse.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(resp)

	if resp.Password != login.Password {
		h.handlerResponse(c, "Wrong password", http.StatusBadRequest, "Wrong password")
		return
	}

	token, _ := helper.GenerateJWT(map[string]interface{}{
		"head_nurse_id": resp.Id,
	}, time.Hour*360, h.cfg.SecretKey)

	h.handlerResponse(c, "token", http.StatusCreated, token)
}

// HeadNurseRegister godoc
// @ID head_nurse_register
// @Router /head_nurse_register [POST]
// @Summary HeadNurseRegister
// @Description HeadNurseRegister
// @Tags HeadNurseRegister
// @Accept json
// @Procedure json
// @Param head_nurse_register body models.CreateHeadNurse true "CreateHeadNurseRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) HeadNurseRegister(c *gin.Context) {

	var createHeadNurse models.CreateHeadNurse
	var id string
	err := c.ShouldBindJSON(&createHeadNurse)
	if err != nil {
		h.handlerResponse(c, "error headNurse should bind json", http.StatusBadRequest, err.Error())
		return
	}

	if len(createHeadNurse.Password) < 7 {
		h.handlerResponse(c, "Password should inculude more than 7 elements", http.StatusBadRequest, errors.New("Password len should inculude more than 8 elements"))
		return
	}

	resp, err := h.strg.HeadNurse().GetByID(context.Background(), &models.HeadNursePrimaryKey{Email: createHeadNurse.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			id, err = h.strg.HeadNurse().Create(context.Background(), &createHeadNurse)
			if err != nil {
				h.handlerResponse(c, "storage.headNurse.create", http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			h.handlerResponse(c, "headNurse already exist", http.StatusInternalServerError, err.Error())
			return
		}
	} else if err == nil {
		h.handlerResponse(c, "headNurse already exist", http.StatusBadRequest, nil)
		return
	}
	resp, _ = h.strg.HeadNurse().GetByID(context.Background(), &models.HeadNursePrimaryKey{Id: id})

	h.handlerResponse(c, "create headNurse resposne", http.StatusCreated, resp)
}

// DoctorLogin godoc
// @ID doctor_login
// @Router /doctor_login [POST]
// @Summary DoctorLogin
// @Description DoctorLogin
// @Tags DoctorLogin
// @Accept json
// @Procedure json
// @Param doctor_login body models.LoginInfo true "DoctorLoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DoctorLogin(c *gin.Context) {
	var login models.LoginInfo

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create Doctor", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.Doctor().GetByID(context.Background(), &models.DoctorPrimaryKey{Email: login.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "Doctor does not exist", http.StatusBadRequest, "Doctor does not exist")
			return
		}
		h.handlerResponse(c, "storage.Doctor.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(resp)

	if resp.Password != login.Password {
		h.handlerResponse(c, "Wrong password", http.StatusBadRequest, "Wrong password")
		return
	}

	token, _ := helper.GenerateJWT(map[string]interface{}{
		"doctor_id": resp.Id,
	}, time.Hour*360, h.cfg.SecretKey)

	h.handlerResponse(c, "token", http.StatusCreated, token)
}

// DoctorRegister godoc
// @ID doctor_register
// @Router /doctor_register [POST]
// @Summary DoctorRegister
// @Description DoctorRegister
// @Tags DoctorRegister
// @Accept json
// @Procedure json
// @Param doctor_register body models.CreateDoctor true "CreateDoctorRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) DoctorRegister(c *gin.Context) {

	var createDoctor models.CreateDoctor
	var id string
	err := c.ShouldBindJSON(&createDoctor)
	if err != nil {
		h.handlerResponse(c, "error doctor should bind json", http.StatusBadRequest, err.Error())
		return
	}

	if len(createDoctor.Password) < 7 {
		h.handlerResponse(c, "Password should inculude more than 7 elements", http.StatusBadRequest, errors.New("Password len should inculude more than 8 elements"))
		return
	}

	resp, err := h.strg.Doctor().GetByID(context.Background(), &models.DoctorPrimaryKey{Email: createDoctor.Email})

	if err != nil {
		if err.Error() == "no rows in result set" {
			id, err = h.strg.Doctor().Create(context.Background(), &createDoctor)
			if err != nil {
				h.handlerResponse(c, "storage.Doctor.create", http.StatusInternalServerError, err.Error())
				return
			}
		} else {
			h.handlerResponse(c, "Doctor already exist", http.StatusInternalServerError, err.Error())
			return
		}
	} else if err == nil {
		h.handlerResponse(c, "Doctor already exist", http.StatusBadRequest, nil)
		return
	}
	resp, _ = h.strg.Doctor().GetByID(context.Background(), &models.DoctorPrimaryKey{Id: id})

	h.handlerResponse(c, "create Doctor resposne", http.StatusCreated, resp)
}

// SisAdminLogin godoc
// @ID sis_admin_login
// @Router /sis_admin_login [POST]
// @Summary SisAdminLogin
// @Description SisAdminLogin
// @Tags SisAdminLogin
// @Accept json
// @Procedure json
// @Param sis_admin_login body models.LoginInfo true "LoginRequest"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server error"
func (h *handler) SisAdminLogin(c *gin.Context) {
	var sis_admin_login models.LoginInfo

	err := c.ShouldBindJSON(&sis_admin_login) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create sis_admin", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.strg.SisAdmin().GetByID(context.Background(), &models.SisAdminPrimaryKey{Email: sis_admin_login.Email})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "sis_admin does not exist", http.StatusBadRequest, "User does not exist")
			return
		}
		h.handlerResponse(c, "storage.sis_admin.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println(resp)

	if resp.Password != sis_admin_login.Password {
		h.handlerResponse(c, "Wrong password", http.StatusBadRequest, "Wrong password")
		return
	}

	token, err := helper.GenerateJWT(map[string]interface{}{
		"user_id": resp.Id,
	}, time.Hour*360, h.cfg.SecretKey)

	h.handlerResponse(c, "token", http.StatusCreated, token)
}
