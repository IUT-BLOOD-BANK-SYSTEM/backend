package api

import (
	_ "app/api/docs"
	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewApi(r *gin.Engine, cfg *config.Config, storage storage.StorageI, logger logger.LoggerI) {

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	handler := handler.NewHandler(cfg, storage, logger)

	r.Use(customCORSMiddleware())

	// User Register Api
	r.POST("/register", handler.Register)
	// User Login Api
	r.POST("/login", handler.Login)
	//HeadNurseRegister
	r.POST("/head_nurse_register", handler.HeadNurseRegister)
	// HeadNurse Login Api
	r.POST("/head_nurse_login", handler.HeadNurseLogin)
	//Doctor Register
	r.POST("/doctor_register", handler.DoctorRegister)
	// Doctor Login Api
	r.POST("/doctor_login", handler.DoctorLogin)
	// SisAdmin Login Api
	r.POST("/sis_admin_login", handler.SisAdminLogin)
	//BloodDonationHistory
	r.GET("/blood_donation_history", handler.GetListBloodDonation)
	//BloodAcceptanceHistory
	r.GET("/blood_acceptance_history", handler.GetListBloodAcceptance)
	//BloodRequesHistory
	r.GET("/blood_request_history", handler.GetListBloodRequest)

	//User
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	//BloodType
	r.POST("/blood_type", handler.CreateBloodType)
	r.GET("/blood_type/:id", handler.GetByIdBloodType)
	r.GET("/blood_type", handler.GetListBloodType)
	r.PUT("/blood_type/:id", handler.UpdateBloodType)
	r.DELETE("/blood_type/:id", handler.DeleteBloodType)
	//Region
	r.POST("/region", handler.CreateRegion)
	r.GET("/region/:id", handler.GetByIdRegion)
	r.GET("/region", handler.GetListRegion)
	r.PUT("/region/:id", handler.UpdateRegion)
	r.DELETE("/region/:id", handler.DeleteRegion)
	//Hospital
	r.POST("/hospital", handler.CreateHospital)
	r.GET("/hospital/:id", handler.GetByIdHospital)
	r.GET("/hospital", handler.GetListHospital)
	r.PUT("/hospital/:id", handler.UpdateHospital)
	r.DELETE("/hospital/:id", handler.DeleteHospital)
	//HeadNurse
	r.POST("/head_nurse", handler.CreateHeadNurse)
	r.GET("/head_nurse/:id", handler.GetByIdHeadNurse)
	r.GET("/head_nurse", handler.GetListHeadNurse)
	r.PUT("/head_nurse/:id", handler.UpdateHeadNurse)
	r.DELETE("/head_nurse/:id", handler.DeleteHeadNurse)
	//Doctor
	r.POST("/doctor", handler.CreateDoctor)
	r.GET("/doctor/:id", handler.GetByIdDoctor)
	r.GET("/doctor", handler.GetListDoctor)
	r.PUT("/doctor/:id", handler.UpdateDoctor)
	r.DELETE("/doctor/:id", handler.DeleteDoctor)
	//BloodDonation
	r.POST("/blood_donation", handler.CreateBloodDonation)
	r.GET("/blood_donation/:id", handler.GetByIdBloodDonation)
	r.GET("/blood_donation", handler.GetListBloodDonation)
	r.PUT("/blood_donation/:id", handler.UpdateBloodDonation)
	r.DELETE("/blood_donation/:id", handler.DeleteBloodDonation)
	//BloodRequest
	r.POST("/blood_request", handler.CreateBloodRequest)
	r.GET("/blood_request/:id", handler.GetByIdBloodRequest)
	r.GET("/blood_request", handler.GetListBloodRequest)
	r.PUT("/blood_request/:id", handler.UpdateBloodRequest)
	r.DELETE("/blood_request/:id", handler.DeleteBloodRequest)
	///BloodInventory
	r.POST("/blood_inventory", handler.CreateBloodInventory)
	r.GET("/blood_inventory/:id", handler.GetByIdBloodInventory)
	r.GET("/blood_inventory", handler.GetListBloodInventory)
	r.PUT("/blood_inventory/:id", handler.UpdateBloodInventory)
	r.DELETE("/blood_inventory/:id", handler.DeleteBloodInventory)
	//Appointment
	r.POST("/appointment", handler.CreateAppointment)
	r.GET("/appointment/:id", handler.GetByIdAppointment)
	r.GET("/appointment", handler.GetListAppointment)
	r.PUT("/appointment/:id", handler.UpdateAppointment)
	r.DELETE("/appointment/:id", handler.DeleteAppointment)
	//BloodAcceptance
	r.POST("/blood_acceptance", handler.CreateBloodAcceptance)
	r.GET("/blood_acceptance/:id", handler.GetByIdBloodAcceptance)
	r.GET("/blood_acceptance", handler.GetListBloodAcceptance)
	r.PUT("/blood_acceptance/:id", handler.UpdateBloodAcceptance)
	r.DELETE("/blood_acceptance/:id", handler.DeleteBloodAcceptance)
	//SisAmin
	r.POST("/sis_admin", handler.CreateSisAdmin)
	r.GET("/sis_admin/:id", handler.GetByIdSisAdmin)
	r.GET("/sis_admin", handler.GetListSisAdmin)
	r.PUT("/sis_admin/:id", handler.UpdateSisAdmin)
	r.DELETE("/sis_admin/:id", handler.DeleteSisAdmin)
	///////////////////////////////////////////////////

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Accesp-Encoding, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
