package storage

import (
	"app/api/models"
	"context"
)

type StorageI interface {
	Close()
	User() UserRepoI
	BloodType() BloodTypeRepoI
	Region() RegionRepoI
	Hospital() HospitalRepoI
	HeadNurse() HeadNurseRepoI
	Doctor() DoctorRepoI
	BloodDonation() BloodDonationRepoI
	BloodRequest() BloodRequestRepoI
	BloodInventory() BloodInventoryRepoI
	Appointment() AppointmentRepoI
	BloodAcceptance() BloodAcceptanceRepoI
	SisAdmin() SisAdminRepoI
}

type UserRepoI interface {
	Create(context.Context, *models.CreateUser) (string, error)
	GetByID(context.Context, *models.UserPrimaryKey) (*models.User, error)
	GetList(context.Context, *models.UserGetListRequest) (*models.UserGetListResponse, error)
	Update(context.Context, *models.UpdateUser) (int64, error)
	Delete(context.Context, *models.UserPrimaryKey) error
}

type BloodTypeRepoI interface {
	Create(context.Context, *models.CreateBloodType) (string, error)
	GetByID(context.Context, *models.BloodTypePrimaryKey) (*models.BloodType, error)
	GetList(context.Context, *models.BloodTypeGetListRequest) (*models.BloodTypeGetListResponse, error)
	Update(context.Context, *models.UpdateBloodType) (int64, error)
	Delete(context.Context, *models.BloodTypePrimaryKey) error
}

type RegionRepoI interface {
	Create(context.Context, *models.CreateRegion) (string, error)
	GetByID(context.Context, *models.RegionPrimaryKey) (*models.Region, error)
	GetList(context.Context, *models.RegionGetListRequest) (*models.RegionGetListResponse, error)
	Update(context.Context, *models.UpdateRegion) (int64, error)
	Delete(context.Context, *models.RegionPrimaryKey) error
}

type HospitalRepoI interface {
	Create(context.Context, *models.CreateHospital) (string, error)
	GetByID(context.Context, *models.HospitalPrimaryKey) (*models.Hospital, error)
	GetList(context.Context, *models.HospitalGetListRequest) (*models.HospitalGetListResponse, error)
	Update(context.Context, *models.UpdateHospital) (int64, error)
	Delete(context.Context, *models.HospitalPrimaryKey) error
}

type HeadNurseRepoI interface {
	Create(context.Context, *models.CreateHeadNurse) (string, error)
	GetByID(context.Context, *models.HeadNursePrimaryKey) (*models.HeadNurse, error)
	GetList(context.Context, *models.HeadNurseGetListRequest) (*models.HeadNurseGetListResponse, error)
	Update(context.Context, *models.UpdateHeadNurse) (int64, error)
	Delete(context.Context, *models.HeadNursePrimaryKey) error
}

type DoctorRepoI interface {
	Create(context.Context, *models.CreateDoctor) (string, error)
	GetByID(context.Context, *models.DoctorPrimaryKey) (*models.Doctor, error)
	GetList(context.Context, *models.DoctorGetListRequest) (*models.DoctorGetListResponse, error)
	Update(context.Context, *models.UpdateDoctor) (int64, error)
	Delete(context.Context, *models.DoctorPrimaryKey) error
}

type BloodDonationRepoI interface {
	Create(context.Context, *models.CreateBloodDonation) (string, error)
	GetByID(context.Context, *models.BloodDonationPrimaryKey) (*models.BloodDonation, error)
	GetList(context.Context, *models.BloodDonationGetListRequest) (*models.BloodDonationGetListResponse, error)
	Update(context.Context, *models.UpdateBloodDonation) (int64, error)
	Delete(context.Context, *models.BloodDonationPrimaryKey) error
}

type BloodRequestRepoI interface {
	Create(context.Context, *models.CreateBloodRequest) (string, error)
	GetByID(context.Context, *models.BloodRequestPrimaryKey) (*models.BloodRequest, error)
	GetList(context.Context, *models.BloodRequestGetListRequest) (*models.BloodRequestGetListResponse, error)
	Update(context.Context, *models.UpdateBloodRequest) (int64, error)
	Delete(context.Context, *models.BloodRequestPrimaryKey) error
}

type BloodInventoryRepoI interface {
	Create(context.Context, *models.CreateBloodInventory) (string, error)
	GetByID(context.Context, *models.BloodInventoryPrimaryKey) (*models.BloodInventory, error)
	GetList(context.Context, *models.BloodInventoryGetListRequest) (*models.BloodInventoryGetListResponse, error)
	Update(context.Context, *models.UpdateBloodInventory) (int64, error)
	Delete(context.Context, *models.BloodInventoryPrimaryKey) error
}

type AppointmentRepoI interface {
	Create(context.Context, *models.CreateAppointment) (string, error)
	GetByID(context.Context, *models.AppointmentPrimaryKey) (*models.Appointment, error)
	GetList(context.Context, *models.AppointmentGetListRequest) (*models.AppointmentGetListResponse, error)
	Update(context.Context, *models.UpdateAppointment) (int64, error)
	Delete(context.Context, *models.AppointmentPrimaryKey) error
}

type BloodAcceptanceRepoI interface {
	Create(context.Context, *models.CreateBloodAcceptance) (string, error)
	GetByID(context.Context, *models.BloodAcceptancePrimaryKey) (*models.BloodAcceptance, error)
	GetList(context.Context, *models.BloodAcceptanceGetListRequest) (*models.BloodAcceptanceGetListResponse, error)
	Update(context.Context, *models.UpdateBloodAcceptance) (int64, error)
	Delete(context.Context, *models.BloodAcceptancePrimaryKey) error
}

type SisAdminRepoI interface {
	Create(context.Context, *models.CreateSisAdmin) (string, error)
	GetByID(context.Context, *models.SisAdminPrimaryKey) (*models.SisAdmin, error)
	GetList(context.Context, *models.SisAdminGetListRequest) (*models.SisAdminGetListResponse, error)
	Update(context.Context, *models.UpdateSisAdmin) (int64, error)
	Delete(context.Context, *models.SisAdminPrimaryKey) error
}
