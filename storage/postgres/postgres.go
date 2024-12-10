package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type store struct {
	db               *pgxpool.Pool
	user             *userRepo
	blood_type       *blood_typeRepo
	region           *regionRepo
	hospital         *hospitalRepo
	head_nurse       *head_nurseRepo
	doctor           *doctorRepo
	blood_donation   *blood_donationRepo
	blood_request    *blood_requestRepo
	blood_inventory  *blood_inventoryRepo
	appointment      *appointmentRepo
	blood_acceptance *blood_acceptanceRepo
	sis_admin        *sis_adminRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) User() storage.UserRepoI {

	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}
func (s *store) BloodType() storage.BloodTypeRepoI {

	if s.blood_type == nil {
		s.blood_type = NewBloodTypeRepo(s.db)
	}

	return s.blood_type
}

func (s *store) Region() storage.RegionRepoI {

	if s.region == nil {
		s.region = NewRegionRepo(s.db)
	}

	return s.region
}

func (s *store) Hospital() storage.HospitalRepoI {

	if s.hospital == nil {
		s.hospital = NewHospitalRepo(s.db)
	}

	return s.hospital
}

func (s *store) HeadNurse() storage.HeadNurseRepoI {

	if s.head_nurse == nil {
		s.head_nurse = NewHeadNurseRepo(s.db)
	}

	return s.head_nurse
}

func (s *store) Doctor() storage.DoctorRepoI {

	if s.doctor == nil {
		s.doctor = NewDoctorRepo(s.db)
	}

	return s.doctor
}

func (s *store) BloodDonation() storage.BloodDonationRepoI {

	if s.blood_donation == nil {
		s.blood_donation = NewBloodDonationRepo(s.db)
	}

	return s.blood_donation
}

func (s *store) BloodRequest() storage.BloodRequestRepoI {

	if s.blood_request == nil {
		s.blood_request = NewBloodRequestRepo(s.db)
	}

	return s.blood_request
}

func (s *store) BloodInventory() storage.BloodInventoryRepoI {

	if s.blood_inventory == nil {
		s.blood_inventory = NewBloodInventoryRepo(s.db)
	}

	return s.blood_inventory
}

func (s *store) Appointment() storage.AppointmentRepoI {

	if s.appointment == nil {
		s.appointment = NewAppointmentRepo(s.db)
	}

	return s.appointment
}

func (s *store) BloodAcceptance() storage.BloodAcceptanceRepoI {

	if s.blood_acceptance == nil {
		s.blood_acceptance = NewBloodAcceptanceRepo(s.db)
	}

	return s.blood_acceptance
}

func (s *store) SisAdmin() storage.SisAdminRepoI {

	if s.sis_admin == nil {
		s.sis_admin = NewSisAdminRepo(s.db)
	}

	return s.sis_admin
}
