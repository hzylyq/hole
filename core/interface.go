package core

type IEntity interface {
	CreateRepo() (string, error)
	CreateApi() (string, error)
	CreateService() (string, error)
}
