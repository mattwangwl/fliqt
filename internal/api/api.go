package api

import (
	"fliqt/internal/model"
	"fliqt/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func New() *Api {
	return &Api{}
}

type Api struct {
	engine *gin.Engine
}

func (a *Api) Run() {
	a.engine = gin.Default()

	a.router()

	a.engine.Run(":8801")
}

func (a *Api) router() {
	{
		impl := newEmployee()
		g := a.engine.Group("employee")
		g.POST("", impl.Create)
		g.GET("", impl.GetAll)
		g.GET(":id", impl.Get)
		g.PUT(":id", impl.Update)
		g.PATCH(":id/status", impl.UpdateStatus)
	}
}

func newEmployee() *employee {
	return &employee{
		employeeService: service.NewEmployee(),
	}
}

type employee struct {
	employeeService *service.Employee
}

func (e *employee) Create(ctx *gin.Context) {
	req := &struct {
		Name     string `json:"name"`
		Position string `json:"position"`
		Status   int    `json:"status"`
	}{}

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := e.employeeService.Create(ctx, model.Employee{
		Name:     req.Name,
		Position: req.Position,
		Status:   req.Status,
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (e *employee) GetAll(ctx *gin.Context) {
	req := &struct {
		Name     *string  `form:"name"`
		Position []string `form:"position"`
		Status   *int     `form:"status"`
	}{}

	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employees, err := e.employeeService.List(ctx, model.EmployeeListCond{
		Name:     req.Name,
		Position: req.Position,
		Status:   req.Status,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": employees})
}

func (e *employee) Get(ctx *gin.Context) {
	req := &struct {
		ID int64 `uri:"id"`
	}{}

	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := e.employeeService.Get(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (e *employee) Update(ctx *gin.Context) {
	req := &struct {
		ID       int64  `uri:"id"`
		Name     string `json:"name"`
		Position string `json:"position"`
		Status   int    `json:"status"`
	}{}

	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := e.employeeService.Update(ctx, model.EmployeeUpdateCond{
		ID:       req.ID,
		Name:     &req.Name,
		Position: &req.Position,
		Status:   &req.Status,
	}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (e *employee) UpdateStatus(ctx *gin.Context) {
	req := &struct {
		ID     int64 `uri:"id"`
		Status int   `json:"status"`
	}{}

	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := e.employeeService.UpdateStatus(ctx, req.ID, req.Status); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
