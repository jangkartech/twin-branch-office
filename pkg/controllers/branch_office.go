package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jangkartech/twin-branch-office/pkg/dto"
	"github.com/jangkartech/twin-branch-office/pkg/services"
	"github.com/jangkartech/twin-branch-office/pkg/validators"
	"github.com/jangkartech/twin-util/pkg/util"
)

type BranchOfficeControllerInterface interface {
	GetBranchOffices(ctx *gin.Context)
	ShowBranchOffice(ctx *gin.Context)
	CreateBranchOffice(ctx *gin.Context)
	UpdateBranchOffice(ctx *gin.Context)
	SoftDeleteBranchOffice(ctx *gin.Context)
	RestoreBranchOffice(ctx *gin.Context)
	HardDeleteBranchOffice(ctx *gin.Context)

	GetSimpleBranchOffices(ctx *gin.Context)
}

type branchOfficeController struct {
	branchOfficeService services.BranchOfficeServiceInterface
}

func NewBranchOfficeController(branchOfficeService services.BranchOfficeServiceInterface) BranchOfficeControllerInterface {
	return &branchOfficeController{
		branchOfficeService: branchOfficeService,
	}
}

// GetBranchOffices godoc
// @Summary       Retrieve a list of branch offices
// @Description   Fetches a filtered list of branch offices and returns the results in JSON format.
// @Tags          Branch Offices
// @Produce       json
// @Param         branch_office body dto.GetBranchOfficeRequest true "JSON payload for branch office filtering"
// @Success       200 {object} dto.GetBranchOfficeResponse
// @Failure       500 {object} dto.InternalServerErrorResponse
// @Failure       422 {object} dto.UnprocessableEntityResponse{error=dto.GetBranchOfficeValidationResponse}
// @Failure       400 {object} dto.BadRequestResponse{error=dto.GetBranchOfficeValidationResponse}
// @Router        /branch-offices [get]
func (c *branchOfficeController) GetBranchOffices(ctx *gin.Context) {
	req, err := validators.ValidateGetBranchOfficeRequest(ctx)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	totalRows, totalPages, err := c.branchOfficeService.GetTotalRowsAndPages(ctx, *req)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	data, err := c.branchOfficeService.GetBranchOfficeList(ctx, *req)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	var responseData []*dto.BranchOfficeResource
	for _, item := range data {
		responseData = append(responseData, item.ToDtoResponse())
	}

	ctx.JSON(http.StatusOK, dto.GetBranchOfficeResponse{
		Data: responseData,
		Meta: &dto.BranchOfficeMeta{
			Pagination: &dto.Pagination{
				Limit:      *req.Limit,
				Page:       *req.Page,
				TotalRows:  totalRows,
				TotalPages: totalPages,
			},
		},
		Message: util.ResponseMessage(http.StatusOK),
	})
	return
}

// ShowBranchOffice godoc
// @Summary       Retrieve detailed information about a specific branch office by ID
// @Description   Retrieves and presents detailed information about a specific branch office in JSON format.
// @Tags          Branch Offices
// @Produce       json
// @Param         id  path  string  true "Unique identifier for the branch office"
// @Success       200 {object} dto.ShowBranchOfficeResponse
// @Failure       500 {object} dto.InternalServerErrorResponse
// @Failure       404 {object} dto.NotFoundResponse
// @Router        /branch-office/{id} [get]
func (c *branchOfficeController) ShowBranchOffice(ctx *gin.Context) {
	branchExists, err := c.branchOfficeService.ExistsBranchOfficeById(ctx, ctx.Param("id"), false)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	} else if !branchExists {
		util.HandleErrorResponse(ctx, http.StatusNotFound, err)
		return
	}

	data, err := c.branchOfficeService.GetBranchOfficeById(ctx, ctx.Param("id"))
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.ShowBranchOfficeResponse{
		Data:    data.ToDtoResponse(),
		Message: util.ResponseMessage(http.StatusOK),
	})
	return
}

// CreateBranchOffice godoc
// @Summary       Create a new branch office
// @Description   Creates a new branch office based on the provided data and returns the newly created branch office details in JSON format.
// @Tags          Branch Offices
// @Produce       json
// @Param         branch_office  body  dto.CreateBranchOfficeRequest  true  "JSON object containing branch office data"
// @Success       201 {object} dto.CreateBranchOfficeResponse
// @Failure       500 {object} dto.InternalServerErrorResponse
// @Failure       422 {object} dto.UnprocessableEntityResponse{error=dto.CreateBranchOfficeValidationResponse}
// @Failure       400 {object} dto.BadRequestResponse{error=dto.CreateBranchOfficeValidationResponse}
// @Router        /branch-office [post]
func (c *branchOfficeController) CreateBranchOffice(ctx *gin.Context) {
	req, err := validators.ValidateCreateBranchOfficeRequest(ctx, c.branchOfficeService)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}
	data, err := c.branchOfficeService.CreateBranchOffice(ctx, *req)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, dto.CreateBranchOfficeResponse{
		Data:    data.ToDtoResponse(),
		Message: util.ResponseMessage(http.StatusCreated),
	})
	return
}

// UpdateBranchOffice godoc
// @Summary       Update information of a specific branch office by ID
// @Description   Updates the information of a specific branch office based on the provided data and returns the updated branch office details in JSON format.
// @Tags          Branch Offices
// @Produce       json
// @Param         id  path  string  true  "ID of the branch office to be updated"
// @Param         branch_office  body  dto.UpdateBranchOfficeRequest  true  "JSON object containing updated branch office data"
// @Success       200 {object} dto.UpdateBranchOfficeResponse
// @Failure       500 {object} dto.InternalServerErrorResponse
// @Failure       422 {object} dto.UnprocessableEntityResponse{error=dto.UpdateBranchOfficeValidationResponse}
// @Failure       400 {object} dto.BadRequestResponse{error=dto.UpdateBranchOfficeValidationResponse}
// @Failure       404 {object} dto.NotFoundResponse
// @Router        /branch-office/{id} [put]
func (c *branchOfficeController) UpdateBranchOffice(ctx *gin.Context) {
	id := ctx.Param("id")
	branchExists, err := c.branchOfficeService.ExistsBranchOfficeById(ctx, id, false)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	} else if !branchExists {
		util.HandleErrorResponse(ctx, http.StatusNotFound, err)
		return
	}

	req, err := validators.ValidateUpdateBranchOfficeRequest(ctx, c.branchOfficeService, id)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	data, err := c.branchOfficeService.UpdateBranchOfficeById(ctx, id, *req)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.UpdateBranchOfficeResponse{
		Data:    data.ToDtoResponse(),
		Message: util.ResponseMessage(http.StatusOK),
	})
	return
}

// SoftDeleteBranchOffice godoc
// @Summary       Soft Delete a branch office by ID
// @Description   Performs a soft deletion of a specific branch office based on the provided ID and returns a confirmation message in JSON format.
// @Tags          Branch Offices
// @Produce       json
// @Param         id  path  string  true  "ID of the branch office to be soft deleted"
// @Success       200 {object} dto.DeleteBranchOfficeResponse
// @Failure       500 {object} dto.InternalServerErrorResponse
// @Failure       404 {object} dto.NotFoundResponse
// @Router        /branch-office/{id} [delete]
func (c *branchOfficeController) SoftDeleteBranchOffice(ctx *gin.Context) {
	branchExists, err := c.branchOfficeService.ExistsBranchOfficeById(ctx, ctx.Param("id"), false)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	} else if !branchExists {
		util.HandleErrorResponse(ctx, http.StatusNotFound, err)
		return
	}

	err = c.branchOfficeService.SoftDeleteBranchOfficeById(ctx, ctx.Param("id"))
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.DeleteBranchOfficeResponse{
		Message: util.ResponseMessage(http.StatusOK),
	})
	return
}

// HardDeleteBranchOffice godoc
// @Summary       Permanently Delete a branch office by ID
// @Description   Performs a permanent deletion of a specific branch office based on the provided ID.
// @Tags          Branch Offices
// @Produce       json
// @Param         id  path  string  true "ID of the branch office to be permanently deleted"
// @Success       200 {object} dto.HardDeleteBranchOfficeResponse
// @Failure       500 {object} dto.InternalServerErrorResponse
// @Failure       404 {object} dto.NotFoundResponse
// @Router        /branch-office/hard-delete/{id} [delete]
func (c *branchOfficeController) HardDeleteBranchOffice(ctx *gin.Context) {
	branchExists, err := c.branchOfficeService.ExistsBranchOfficeById(ctx, ctx.Param("id"), true)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	} else if !branchExists {
		util.HandleErrorResponse(ctx, http.StatusNotFound, err)
		return
	}

	err = c.branchOfficeService.HardDeleteBranchOfficeById(ctx, ctx.Param("id"))
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.HardDeleteBranchOfficeResponse{
		Message: util.ResponseMessage(http.StatusOK),
	})
	return
}

// RestoreBranchOffice godoc
// @Summary       Restore a branch office by ID
// @Description   Restores a previously soft-deleted branch office based on the provided ID and returns a confirmation message in JSON format.
// @Tags          Branch Offices
// @Produce       json
// @Param         id  path  string  true "ID of the branch office to be restored"
// @Success       200 {object} dto.RestoreBranchOfficeResponse
// @Failure       500 {object} dto.InternalServerErrorResponse
// @Failure       404 {object} dto.NotFoundResponse
// @Router        /branch-office/{id} [patch]
func (c *branchOfficeController) RestoreBranchOffice(ctx *gin.Context) {
	branchExists, err := c.branchOfficeService.ExistsBranchOfficeById(ctx, ctx.Param("id"), true)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	} else if !branchExists {
		util.HandleErrorResponse(ctx, http.StatusNotFound, err)
		return
	}

	err = c.branchOfficeService.RestoreBranchOfficeById(ctx, ctx.Param("id"))
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dto.RestoreBranchOfficeResponse{
		Message: util.ResponseMessage(http.StatusOK),
	})
	return
}

// GetSimpleBranchOffices	godoc
// @Summary       			Retrieve a list of simple branch offices
// @Description   			Fetches a filtered list of simple branch offices and returns the results in JSON format.
// @Tags          			Branch Offices
// @Produce       			json
// @Param         			branch_office query dto.GetSimpleBranchOfficeRequest true "JSON payload for simple branch office filtering"
// @Success       			200 {object} dto.GetSimpleBranchOfficeResponse
// @Failure       			500 {object} dto.InternalServerErrorResponse
// @Failure       			422 {object} dto.UnprocessableEntityResponse{error=dto.GetSimpleBranchOfficeValidationResponse}
// @Failure       			400 {object} dto.BadRequestResponse{error=dto.GetSimpleBranchOfficeValidationResponse}
// @Router        			/branch-offices/simple [get]
func (c *branchOfficeController) GetSimpleBranchOffices(ctx *gin.Context) {
	req, err := validators.ValidateGetSimpleBranchOfficeRequest(ctx)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusUnprocessableEntity, err)
		return
	}

	data, err := c.branchOfficeService.GetSimpleBranchOfficeList(ctx, *req)
	if err != nil {
		util.HandleErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	var responseData []*dto.SimpleBranchOfficeResource
	for _, item := range data {
		responseData = append(responseData, item.ToDtoSimpleResponse())
	}

	ctx.JSON(http.StatusOK, dto.GetSimpleBranchOfficeResponse{
		Data:    responseData,
		Message: util.ResponseMessage(http.StatusOK),
	})
	return
}
