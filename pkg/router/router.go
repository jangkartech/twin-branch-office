package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jangkartech/twin-branch-office/pkg/controllers"
	"github.com/jangkartech/twin-branch-office/pkg/repos"
	"github.com/jangkartech/twin-branch-office/pkg/services"
)

func Register(route *gin.Engine) {
	branchOfficeRepo := repos.NewBranchOfficeRepo()
	branchOfficeService := services.NewBranchOfficeService(branchOfficeRepo)
	branchOfficeController := controllers.NewBranchOfficeController(branchOfficeService)
	route.GET("/branch-offices", branchOfficeController.GetBranchOffices)
	route.POST("/branch-office", branchOfficeController.CreateBranchOffice)
	route.GET("/branch-office/:id", branchOfficeController.ShowBranchOffice)
	route.PUT("/branch-office/:id", branchOfficeController.UpdateBranchOffice)
	route.DELETE("/branch-office/:id", branchOfficeController.SoftDeleteBranchOffice)
	route.DELETE("/branch-office/hard-delete/:id", branchOfficeController.HardDeleteBranchOffice)
	route.PATCH("/branch-office/:id", branchOfficeController.RestoreBranchOffice)
	route.GET("/branch-offices/simple", branchOfficeController.GetSimpleBranchOffices)

}
