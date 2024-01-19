package dto

import "github.com/jangkartech/twin-util/pkg/dto"

type BranchOfficeResource struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phone_number"`
	FaxNumber   string `json:"fax_number"`
	City        string `json:"city"`
	CreatedAt   int64  `json:"created_at"`
}

type BranchOfficeMeta struct {
	Pagination *dto.Pagination `json:"pagination"`
}

type ShowBranchOfficeResponse struct {
	Data    *BranchOfficeResource `json:"data"`
	Message string                `json:"message"`
}

type GetBranchOfficeRequest struct {
	Fields  *[]string `validate:"omitempty" form:"fields"`
	Keyword *string   `validate:"omitempty" form:"keyword"`
	Limit   *int      `validate:"omitempty" form:"limit"`
	Page    *int      `validate:"omitempty" form:"page"`
	Status  *string   `validate:"omitempty" form:"status"`
}

type GetBranchOfficeValidationResponse struct {
	Fields  *string `json:"fields"`
	Keyword *string `json:"keyword"`
	Limit   *string `json:"limit"`
	Page    *string `json:"page"`
	Status  *string `json:"status"`
}

type GetBranchOfficeResponse struct {
	Data    []*BranchOfficeResource `json:"data"`
	Meta    *BranchOfficeMeta       `json:"meta"`
	Message string                  `json:"message"`
}

type CreateBranchOfficeRequest struct {
	Id          string `validate:"required" json:"id"`
	Name        string `validate:"required" json:"name"`
	Address     string `validate:"required" json:"address"`
	PhoneNumber string `validate:"required" json:"phone_number"`
	City        string `validate:"required" json:"city"`
	FaxNumber   string `validate:"required" json:"fax_number"`
}

type CreateBranchOfficeValidationResponse struct {
	Id          *string `json:"id"`
	Name        *string `json:"name"`
	Address     *string `json:"address"`
	PhoneNumber *string `json:"phone_number"`
	City        *string `json:"city"`
	FaxNumber   *string `json:"fax_number"`
}

type CreateBranchOfficeResponse struct {
	Data    *BranchOfficeResource `json:"data"`
	Message string                `json:"message"`
}

type UpdateBranchOfficeRequest struct {
	Name        *string `validate:"omitempty" json:"name"`
	Address     *string `validate:"omitempty" json:"address"`
	PhoneNumber *string `validate:"omitempty" json:"phone_number"`
	City        *string `validate:"omitempty" json:"city"`
	FaxNumber   *string `validate:"omitempty" json:"fax_number"`
}

type UpdateBranchOfficeValidationResponse struct {
	Name        *string `json:"name"`
	Address     *string `json:"address"`
	PhoneNumber *string `json:"phone_number"`
	City        *string `json:"city"`
	FaxNumber   *string `json:"fax_number"`
}

type UpdateBranchOfficeResponse struct {
	Data    *BranchOfficeResource `json:"data"`
	Message string                `json:"message"`
}

type DeleteBranchOfficeResponse struct {
	Message string `json:"message"`
}

type HardDeleteBranchOfficeResponse struct {
	Message string `json:"message"`
}

type RestoreBranchOfficeResponse struct {
	Message string `json:"message"`
}

type SimpleBranchOfficeResource struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetSimpleBranchOfficeRequest struct {
	Keyword *string `validate:"omitempty" form:"keyword"`
}

type GetSimpleBranchOfficeValidationResponse struct {
	Keyword *string `json:"keyword"`
}

type GetSimpleBranchOfficeResponse struct {
	Data    []*SimpleBranchOfficeResource `json:"data"`
	Message string                        `json:"message"`
}
