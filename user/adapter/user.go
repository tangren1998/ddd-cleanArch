package adapter

import (
	"github.com/gin-gonic/gin"
	"github.com/tangren1998/ddd-cheanArch/user/domain"
	"github.com/tangren1998/ddd-cheanArch/user/infra/db"
	"github.com/tangren1998/ddd-cheanArch/user/usecase"
	"net/http"
)

// UserController (控制器, 调用 UseCase)
type UserController struct {
	C *gin.Context
}

func (controller *UserController) GetUserByID(id string) {
	presenter := userPresenter{c: controller.C}
	uc := usecase.UserUseCase{Repo: db.MysqlUserRepo{}, Output: &presenter}
	uc.GetUserByID(id)
}

// UserPresenter -（演示者, 实现 Use Case Output Port）
type userPresenter struct {
	c *gin.Context
}

func (presenter *userPresenter) Standard(user *domain.User, err error) {
	// todo: 处理 err
	presenter.c.JSON(http.StatusOK, gin.H{"id": user.ID, "username": user.Username})
}