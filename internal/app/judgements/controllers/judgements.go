package controllers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/infinity-oj/server-v2/internal/pkg/sessions"

	"github.com/gin-gonic/gin"
	"github.com/infinity-oj/server-v2/internal/app/judgements/services"
	"go.uber.org/zap"
)

type Controller interface {
	CreateJudgement(c *gin.Context)
	GetJudgements(c *gin.Context)
	GetJudgement(c *gin.Context)

	GetTasks(c *gin.Context)
	GetTask(c *gin.Context)
	UpdateTask(c *gin.Context)
	ReserveTask(c *gin.Context)
}

type DefaultController struct {
	logger  *zap.Logger
	service services.JudgementsService
}

func (d *DefaultController) CreateJudgement(c *gin.Context) {
	d.logger.Debug("create judgement")
	session := sessions.GetSession(c)
	if session == nil {
		d.logger.Debug("get principal failed")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	d.logger.Debug("create judgement", zap.Uint64("account id", session.AccountId))

	request := struct {
		ProcessId    uint64 `json:"processId" binding:"required"`
		SubmissionId uint64 `json:"submissionId" binding:"required"`
	}{}

	if err := c.ShouldBind(&request); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": errs.Error(),
		})
		return
	}

	d.logger.Debug("create judgement",
		zap.Uint64("processes id", request.ProcessId),
		zap.Uint64("submission id", request.SubmissionId),
	)

	judgement, err := d.service.CreateJudgement(session.AccountId, request.ProcessId, request.SubmissionId)
	if err != nil {
		d.logger.Error("create judgement", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	d.logger.Debug("create judgement",
		zap.String("new judgement id", judgement.JudgementId),
	)
	c.JSON(200, judgement)
}

func (d *DefaultController) GetJudgements(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (d *DefaultController) GetJudgement(c *gin.Context) {
	c.AbortWithStatus(http.StatusNotImplemented)
}

func (d *DefaultController) GetTasks(c *gin.Context) {
	request := struct {
		Type string `form:"type" binding:"required,gt=0"`
	}{}

	if err := c.ShouldBindQuery(&request); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": errs.Error(),
		})
		return
	}

	d.logger.Debug("get tasks",
		zap.String("page", request.Type),
	)

	tasks, err := d.service.GetTasks(request.Type)
	if err != nil {
		d.logger.Error("get tasks", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (d *DefaultController) GetTask(c *gin.Context) {
	taskId := c.Param("taskId")

	d.logger.Debug("get task", zap.String("task taskId", taskId))

	task, err := d.service.GetTask(taskId)
	if err != nil {
		d.logger.Error("get task", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	if task == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, task)
}

func (d *DefaultController) UpdateTask(c *gin.Context) {
	//session := sessions.GetSession(c)
	//if session == nil {
	//	d.logger.Debug("get principal failed")
	//	c.AbortWithStatus(http.StatusUnauthorized)
	//	return
	//}

	taskId := c.Param("taskId")

	//d.logger.Debug("update task",
	//	zap.Uint64("account id", session.AccountId),
	//	zap.String("task id", taskId),
	//)

	request := struct {
		Token   string `json:"token" binding:"required"`
		Outputs string `json:"outputs" binding:"required"`
	}{}

	if err := c.ShouldBind(&request); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": errs.Error(),
		})
		return
	}

	d.logger.Debug("update task",
		zap.String("token", request.Token),
	)

	task, err := d.service.UpdateTask(request.Token, taskId, request.Outputs)
	if err != nil {
		d.logger.Error("update task", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		return
	}
	c.JSON(200, task)
}

func (d *DefaultController) ReserveTask(c *gin.Context) {
	taskId := c.Param("taskId")

	d.logger.Debug("get task", zap.String("task taskId", taskId))

	token, err := d.service.ReserveTask(taskId)
	if err != nil {
		d.logger.Error("get task", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func New(logger *zap.Logger, s services.JudgementsService) Controller {
	return &DefaultController{
		logger:  logger,
		service: s,
	}
}