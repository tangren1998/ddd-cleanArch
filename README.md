# DDD + Clean Arch 探索与演示
## 请求处理流程
### 请求
```http request
GET /user/1
```
### infra 基础设施层（包含 web，UI，Device，DB，...）
```go
// web：接受请求,并处理数据
// gin example
r.GET('/user/:id', GetUserByID)

func GetUserByID(c *gin.Context) {
	id := c.Param("id") // 提取id
	controller := adapter.UserController{C: c}
	controller.GetUserByID(id)
}
```
```go
// db：实现 domain repo 接口
// mysql example
type MysqlUserRepo struct {}
func (repo MysqlUserRepo) GetByID(id string) (*domain.User, error) {
	var user domain.User
	if result := global.DB.Where("id = ?", id).First(&user); result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
```
### adapter 调度层
```go
// controller(控制器, 调用 UseCase)
type UserController struct {
	C *gin.Context
}
func (controller *UserController) GetUserByID(id string) {
	presenter := userPresenter{c: controller.C}
	uc := usecase.UserUseCase{Repo: db.MysqlUserRepo{}, Output: &presenter}
	uc.GetUserByID(id)
}
```
```go
// presenter（演示者, 实现 Use Case Output Port）
type userPresenter struct {
	c *gin.Context
}
func (presenter *userPresenter) Standard(user *domain.User, err error) {
	// todo: 处理 err
	presenter.c.JSON(http.StatusOK, gin.H{"id": user.ID})
}
```
### useCase 用例层
```go
// 用户用例(实现 Use Case input Port，调用 Use Case Output Port)
// 之所以用useCase命名而不用interactor，是因为Gland提示拼写错误，看着烦
type UserUseCase struct {
	Repo   domain.UserRepo
	Output UserOutput
}
// Use Case Output Port
type UserOutput interface {
	Standard(u *domainUser, err error)
}
func (u *UserUseCase) GetUserByID(id string) {
	user, err := u.Repo.GetByID(id)
	u.Output.Standard(user, err)
}
```

### domain 领域层
```go
type User struct {
	ID string
}
type UserRepo interface {
	GetByID(id string) (*User, error)
}
```