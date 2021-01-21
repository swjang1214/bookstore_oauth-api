# bookstore_oauth-api

## Service

```go
type IRepository interface {
GetById(string) (*AccessToken, *errors.RestError)
}
type IService interface {
GetById(string) (*AccessToken, *errors.RestError)
}
type service struct {
repository IRepository
}
func NewService(repo IRepository) IService {
return &service{
repository: repo,
}
}
func (s *service) GetById(accessTokenId string) (*AccessToken, \*errors.RestError) {
s.repository.GetById(accessTokenId)
}
```

## DBDepository

```go
type IDBRepository interface {
GetById(string) (*access_token.AccessToken, *errors.RestError)
}
type dbRepository struct {
}
func NewRepository() IDBRepository {
return &dbRepository{}
}
func (*dbRepository) GetById(id string) (*access_token.AccessToken, \*errors.RestError) {
}
```

## Http

```go
type IAccessTokenHandler interface {
	GetById(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.IService
}
func NewHandler(service access_token.IService) IAccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}
func (handler *accessTokenHandler) GetById(c *gin.Context) {
	handler.service.GetById(accessTokenId)
}
```

## 사용

```go
    atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))
    router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
```
