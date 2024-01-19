package models

import (
	"time"

	"github.com/jangkartech/twin-branch-office/pkg/dto"
	"gorm.io/gorm"
)

type BranchOffice struct {
	Id          string         `gorm:"type:varchar(36);primaryKey;" json:"id"`
	Name        string         `gorm:"type:varchar(100);" json:"name"`
	Address     string         `gorm:"type:varchar(100);" json:"address"`
	PhoneNumber string         `gorm:"type:varchar(100);" json:"phone_number"`
	FaxNumber   string         `gorm:"type:varchar(100);" json:"fax_number"`
	City        string         `gorm:"type:varchar(100);" json:"city"`
	CreatedAt   time.Time      `gorm:"default:CURRENT_TIMESTAMP;" json:"created_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (m *BranchOffice) ToDtoResponse() *dto.BranchOfficeResource {
	return &dto.BranchOfficeResource{
		Id:          m.Id,
		Name:        m.Name,
		Address:     m.Address,
		PhoneNumber: m.PhoneNumber,
		FaxNumber:   m.FaxNumber,
		City:        m.City,
		CreatedAt:   m.CreatedAt.Unix(),
	}
}

func (m *BranchOffice) ToDtoSimpleResponse() *dto.SimpleBranchOfficeResource {
	return &dto.SimpleBranchOfficeResource{
		Id:   m.Id,
		Name: m.Name,
	}
}
