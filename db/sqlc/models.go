// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type Gender string

const (
	GenderM Gender = "M"
	GenderF Gender = "F"
)

func (e *Gender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Gender(s)
	case string:
		*e = Gender(s)
	default:
		return fmt.Errorf("unsupported scan type for Gender: %T", src)
	}
	return nil
}

type NullGender struct {
	Gender Gender `json:"gender"`
	Valid  bool   `json:"valid"` // Valid is true if Gender is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullGender) Scan(value interface{}) error {
	if value == nil {
		ns.Gender, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Gender.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullGender) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Gender), nil
}

type Roles string

const (
	RolesADMIN    Roles = "ADMIN"
	RolesGENERAL  Roles = "GENERAL"
	RolesASISTANT Roles = "ASISTANT"
)

func (e *Roles) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Roles(s)
	case string:
		*e = Roles(s)
	default:
		return fmt.Errorf("unsupported scan type for Roles: %T", src)
	}
	return nil
}

type NullRoles struct {
	Roles Roles `json:"roles"`
	Valid bool  `json:"valid"` // Valid is true if Roles is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRoles) Scan(value interface{}) error {
	if value == nil {
		ns.Roles, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Roles.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRoles) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Roles), nil
}

type Status string

const (
	StatusACTIVE   Status = "ACTIVE"
	StatusINACTIVE Status = "INACTIVE"
)

func (e *Status) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Status(s)
	case string:
		*e = Status(s)
	default:
		return fmt.Errorf("unsupported scan type for Status: %T", src)
	}
	return nil
}

type NullStatus struct {
	Status Status `json:"status"`
	Valid  bool   `json:"valid"` // Valid is true if Status is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullStatus) Scan(value interface{}) error {
	if value == nil {
		ns.Status, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Status.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullStatus) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Status), nil
}

type Clinic struct {
	ID           int64              `json:"id"`
	Name         string             `json:"name"`
	Phone        pgtype.Text        `json:"phone"`
	IDDepartment pgtype.Int8        `json:"id_department"`
	IDTown       pgtype.Int8        `json:"id_town"`
	Direction    pgtype.Text        `json:"direction"`
	Status       Status             `json:"status"`
	CreateAt     pgtype.Timestamptz `json:"create_at"`
}

type Department struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Town struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	IDDepartment int64  `json:"id_department"`
}

type User struct {
	ID           int64              `json:"id"`
	FullName     string             `json:"full_name"`
	Email        string             `json:"email"`
	Gender       Gender             `json:"gender"`
	Status       Status             `json:"status"`
	IDDepartment pgtype.Int8        `json:"id_department"`
	IDTown       pgtype.Int8        `json:"id_town"`
	Direction    pgtype.Text        `json:"direction"`
	CreateAt     pgtype.Timestamptz `json:"create_at"`
	IDClinic     int64              `json:"id_clinic"`
	Rol          Roles              `json:"rol"`
}
