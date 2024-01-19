package repos

import (
	"context"
	"fmt"

	"github.com/jangkartech/twin-branch-office/pkg/models"
	"github.com/jangkartech/twin-util/pkg/constant"
	"github.com/jangkartech/twin-util/pkg/db"
	"github.com/jangkartech/twin-util/pkg/util"
)

type BranchOfficeRepoInterface interface {
	GetBranchOfficeList(ctx context.Context, filter GetBranchOfficeListFilter) ([]*models.BranchOffice, error)
	CreateBranchOffice(ctx context.Context, BranchOffice models.BranchOffice) (*models.BranchOffice, error)
	UpdateBranchOfficeById(ctx context.Context, id string, BranchOffice models.BranchOffice) (*models.BranchOffice, error)
	SoftDeleteBranchOfficeById(ctx context.Context, id string) error
	HardDeleteBranchOfficeById(ctx context.Context, id string) error
	RestoreBranchOfficeById(ctx context.Context, id string) error
	GetBranchOfficeById(ctx context.Context, id string, withTrash bool) (*models.BranchOffice, error)
	GetBranchOfficeByField(ctx context.Context, field string, value string, withTrash bool) (*models.BranchOffice, error)
	GetBranchOfficeCount(ctx context.Context, filter GetBranchOfficeListFilter) (int64, error)
}

type branchOfficeRepo struct{}

type GetBranchOfficeListFilter struct {
	Fields  *[]string
	Keyword *string
	Limit   *int
	Page    *int
	Status  *string
}

func NewBranchOfficeRepo() BranchOfficeRepoInterface {
	return &branchOfficeRepo{}
}

func (r *branchOfficeRepo) GetBranchOfficeList(ctx context.Context, filter GetBranchOfficeListFilter) ([]*models.BranchOffice, error) {
	var list []*models.BranchOffice
	res := db.DB.Model(&models.BranchOffice{})

	if filter.Fields != nil && filter.Keyword != nil {
		if len(*filter.Fields) > 0 && *filter.Keyword != "" {
			subQuery := db.DB
			for _, field := range *filter.Fields {
				subQuery = res.Or(fmt.Sprintf("%s ILIKE ?", field), "%"+*filter.Keyword+"%")
			}
			res.Where(subQuery)
		}
	}

	if filter.Status != nil {
		if *filter.Status == constant.StatusDeleted {
			res.Unscoped().Where("deleted_at is not null")
		}
	}

	if filter.Limit != nil && filter.Page != nil {
		util.Paginate(res, *filter.Limit, *filter.Page)
	}

	if err := res.Order("name ASC").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *branchOfficeRepo) GetBranchOfficeById(ctx context.Context, id string, withTrash bool) (*models.BranchOffice, error) {
	var BranchOffice models.BranchOffice
	res := db.DB.Model(&models.BranchOffice{}).Where("id = ?", id)
	if withTrash {
		res.Unscoped()
	}
	if err := res.First(&BranchOffice).Error; err != nil {
		return nil, err
	}
	return &BranchOffice, nil
}

func (r *branchOfficeRepo) GetBranchOfficeByField(ctx context.Context, field string, value string, withTrash bool) (*models.BranchOffice, error) {
	var BranchOffice models.BranchOffice
	res := db.DB.Model(&models.BranchOffice{}).Where(fmt.Sprintf("%s = ?", field), value)
	if withTrash {
		res.Unscoped()
	}
	if err := res.First(&BranchOffice).Error; err != nil {
		return nil, err
	}
	return &BranchOffice, nil
}

func (r *branchOfficeRepo) CreateBranchOffice(ctx context.Context, BranchOffice models.BranchOffice) (*models.BranchOffice, error) {
	res := db.DB.Create(&BranchOffice)
	if err := res.Error; err != nil {
		return nil, err
	}
	return &BranchOffice, nil
}

func (r *branchOfficeRepo) UpdateBranchOfficeById(ctx context.Context, id string, BranchOffice models.BranchOffice) (*models.BranchOffice, error) {
	res := db.DB.Model(&models.BranchOffice{}).Where("id = ?", id).Updates(&BranchOffice).First(&BranchOffice)
	if res.Error != nil {
		return nil, res.Error
	}
	return &BranchOffice, nil
}

func (r *branchOfficeRepo) SoftDeleteBranchOfficeById(ctx context.Context, id string) error {
	BranchOffice := &models.BranchOffice{Id: id}
	res := db.DB.Delete(BranchOffice)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r *branchOfficeRepo) HardDeleteBranchOfficeById(ctx context.Context, id string) error {
	BranchOffice := &models.BranchOffice{Id: id}
	res := db.DB.Unscoped().Delete(BranchOffice)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r *branchOfficeRepo) RestoreBranchOfficeById(ctx context.Context, id string) error {
	res := db.DB.Model(&models.BranchOffice{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r *branchOfficeRepo) GetBranchOfficeCount(ctx context.Context, filter GetBranchOfficeListFilter) (int64, error) {
	var res int64

	query := db.DB.Model(&models.BranchOffice{})
	if filter.Fields != nil && filter.Keyword != nil {
		if len(*filter.Fields) > 0 && *filter.Keyword != "" {
			subQuery := db.DB
			for _, field := range *filter.Fields {
				subQuery = query.Or(fmt.Sprintf("%s ILIKE ?", field), "%"+*filter.Keyword+"%")
			}
			query.Where(subQuery)
		}
	}

	if filter.Status != nil {
		if *filter.Status == constant.StatusDeleted {
			query.Unscoped().Where("deleted_at is not null")
		}
	}
	err := query.Count(&res).Error
	if err != nil {
		return 0, err
	}
	return res, nil
}
