package validators

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jangkartech/twin-branch-office/pkg/dto"
	"github.com/jangkartech/twin-branch-office/pkg/services"
	"github.com/jangkartech/twin-util/pkg/constant"
	"github.com/jangkartech/twin-util/pkg/errors"
	"github.com/jangkartech/twin-util/pkg/logger"
	"github.com/jangkartech/twin-util/pkg/util"
)

func ValidateGetBranchOfficeRequest(ctx *gin.Context) (*dto.GetBranchOfficeRequest, error) {
	validate := validator.New()
	var req dto.GetBranchOfficeRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		return nil, err
	}
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	if req.Page == nil {
		defaultPage := constant.PaginationPage
		req.Page = &defaultPage
	}
	if req.Limit == nil {
		defaultLimit := constant.PaginationLimit
		req.Limit = &defaultLimit
	}

	util.EnsureStatusAllowed(req.Status, []string{constant.StatusDeleted})
	return &req, nil
}

func ValidateCreateBranchOfficeRequest(ctx *gin.Context, branchOfficeService services.BranchOfficeServiceInterface) (*dto.CreateBranchOfficeRequest, error) {
	validate := validator.New()
	var req dto.CreateBranchOfficeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	if name := req.Name; name != "" {
		warehouseExists, err := branchOfficeService.ExistsBranchOfficeByField(ctx, services.ExistsBranchOfficeByFieldInput{
			Field: "name",
			Value: name,
		})
		if err != nil {
			logger.Log.Error(err.Error())
			return nil, err
		}
		if warehouseExists {
			return nil, &errors.DBValidationError{Field: "name", Tag: "exists"}
		}
	}

	return &req, nil
}

func ValidateUpdateBranchOfficeRequest(ctx *gin.Context, branchOfficeService services.BranchOfficeServiceInterface, id string) (*dto.UpdateBranchOfficeRequest, error) {
	var req dto.UpdateBranchOfficeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return nil, err
	}

	return &req, nil
}

func ValidateGetSimpleBranchOfficeRequest(ctx *gin.Context) (*dto.GetSimpleBranchOfficeRequest, error) {
	validate := validator.New()
	var req dto.GetSimpleBranchOfficeRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		return nil, err
	}

	if err := validate.Struct(req); err != nil {
		return nil, err
	}
	return &req, nil
}
