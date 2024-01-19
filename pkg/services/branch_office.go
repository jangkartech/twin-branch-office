package services

import (
	"context"
	"errors"

	"github.com/jangkartech/twin-branch-office/pkg/dto"
	"github.com/jangkartech/twin-branch-office/pkg/models"
	"github.com/jangkartech/twin-branch-office/pkg/repos"
	"github.com/jangkartech/twin-util/pkg/util"
	"gorm.io/gorm"
)

type BranchOfficeServiceInterface interface {
	ExistsBranchOfficeById(ctx context.Context, id string, withTrash bool) (bool, error)
	ExistsBranchOfficeByField(ctx context.Context, input ExistsBranchOfficeByFieldInput) (bool, error)
	GetBranchOfficeList(ctx context.Context, req dto.GetBranchOfficeRequest) ([]*models.BranchOffice, error)
	GetBranchOfficeById(ctx context.Context, id string) (*models.BranchOffice, error)
	CreateBranchOffice(ctx context.Context, req dto.CreateBranchOfficeRequest) (*models.BranchOffice, error)
	UpdateBranchOfficeById(ctx context.Context, id string, req dto.UpdateBranchOfficeRequest) (*models.BranchOffice, error)
	SoftDeleteBranchOfficeById(ctx context.Context, id string) error
	HardDeleteBranchOfficeById(ctx context.Context, id string) error
	RestoreBranchOfficeById(ctx context.Context, id string) error
	GetTotalRowsAndPages(ctx context.Context, req dto.GetBranchOfficeRequest) (int64, int64, error)

	GetSimpleBranchOfficeList(ctx context.Context, req dto.GetSimpleBranchOfficeRequest) ([]*models.BranchOffice, error)
}
type ExistsBranchOfficeByFieldInput struct {
	Field     string
	Value     string
	ExceptId  *string
	WithTrash *bool
}
type branchOfficeService struct {
	branchOfficeRepo repos.BranchOfficeRepoInterface
}

func NewBranchOfficeService(branchOfficeRepo repos.BranchOfficeRepoInterface) BranchOfficeServiceInterface {
	return &branchOfficeService{
		branchOfficeRepo: branchOfficeRepo,
	}
}

func (s *branchOfficeService) convertToBranchOfficeListFilter(req dto.GetBranchOfficeRequest) repos.GetBranchOfficeListFilter {
	if req.Fields != nil {
		fields := util.ClearInvalidFields(*req.Fields, []string{"name", "address"})
		req.Fields = &fields
	} else {
		req.Fields = &[]string{"name"}
	}

	if req.Keyword != nil {
		keyword := util.ClearInvalidKeyword(*req.Keyword)
		req.Keyword = &keyword
	}

	return repos.GetBranchOfficeListFilter{
		Fields:  req.Fields,
		Keyword: req.Keyword,
		Limit:   req.Limit,
		Page:    req.Page,
		Status:  req.Status,
	}
}

func (s *branchOfficeService) ExistsBranchOfficeById(ctx context.Context, id string, withTrash bool) (bool, error) {
	branchOffice, err := s.branchOfficeRepo.GetBranchOfficeById(ctx, id, withTrash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	exists := branchOffice != nil
	return exists, nil
}

func (s *branchOfficeService) ExistsBranchOfficeByField(ctx context.Context, input ExistsBranchOfficeByFieldInput) (bool, error) {
	withTrash := false
	if input.WithTrash != nil {
		withTrash = *input.WithTrash
	}
	branchOffice, err := s.branchOfficeRepo.GetBranchOfficeByField(ctx, input.Field, input.Value, withTrash)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	if branchOffice != nil && (input.ExceptId == nil || *input.ExceptId != branchOffice.Id) {
		return true, nil
	}
	return false, nil
}

func (s *branchOfficeService) GetBranchOfficeList(ctx context.Context, req dto.GetBranchOfficeRequest) ([]*models.BranchOffice, error) {
	filter := s.convertToBranchOfficeListFilter(req)

	res, err := s.branchOfficeRepo.GetBranchOfficeList(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *branchOfficeService) GetBranchOfficeById(ctx context.Context, id string) (*models.BranchOffice, error) {
	res, err := s.branchOfficeRepo.GetBranchOfficeById(ctx, id, false)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *branchOfficeService) CreateBranchOffice(ctx context.Context, req dto.CreateBranchOfficeRequest) (*models.BranchOffice, error) {
	branchOffice := models.BranchOffice{
		Id:          req.Id,
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		FaxNumber:   req.FaxNumber,
		City:        req.City,
	}
	res, err := s.branchOfficeRepo.CreateBranchOffice(ctx, branchOffice)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *branchOfficeService) UpdateBranchOfficeById(ctx context.Context, id string, req dto.UpdateBranchOfficeRequest) (*models.BranchOffice, error) {
	branchOffice := models.BranchOffice{}

	if req.Name != nil {
		branchOffice.Name = *req.Name
	}
	if req.Address != nil {
		branchOffice.Address = *req.Address
	}
	if req.PhoneNumber != nil {
		branchOffice.PhoneNumber = *req.PhoneNumber
	}
	if req.City != nil {
		branchOffice.City = *req.City
	}
	if req.FaxNumber != nil {
		branchOffice.FaxNumber = *req.FaxNumber
	}

	res, err := s.branchOfficeRepo.UpdateBranchOfficeById(ctx, id, branchOffice)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *branchOfficeService) SoftDeleteBranchOfficeById(ctx context.Context, id string) error {
	err := s.branchOfficeRepo.SoftDeleteBranchOfficeById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *branchOfficeService) HardDeleteBranchOfficeById(ctx context.Context, id string) error {
	err := s.branchOfficeRepo.HardDeleteBranchOfficeById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *branchOfficeService) RestoreBranchOfficeById(ctx context.Context, id string) error {
	err := s.branchOfficeRepo.RestoreBranchOfficeById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *branchOfficeService) GetTotalRowsAndPages(ctx context.Context, req dto.GetBranchOfficeRequest) (int64, int64, error) {
	filter := s.convertToBranchOfficeListFilter(req)
	res, err := s.branchOfficeRepo.GetBranchOfficeCount(ctx, filter)
	if err != nil {
		return 0, 0, err
	}
	totalRows := res
	totalPages := util.CalculateTotalPage(totalRows, *req.Limit)
	return totalRows, totalPages, nil
}

func (s *branchOfficeService) GetSimpleBranchOfficeList(ctx context.Context, req dto.GetSimpleBranchOfficeRequest) ([]*models.BranchOffice, error) {
	if req.Keyword != nil {
		keyword := util.ClearInvalidKeyword(*req.Keyword)
		req.Keyword = &keyword
	}

	var filter = repos.GetBranchOfficeListFilter{
		Fields:  &[]string{"name"},
		Keyword: req.Keyword,
	}

	res, err := s.branchOfficeRepo.GetBranchOfficeList(ctx, filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}
