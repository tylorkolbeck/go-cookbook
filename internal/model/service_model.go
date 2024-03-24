package model

type IService interface {
	Initialize()
}

type IHandler interface {
	SetupRouter()
}

type ServiceContainer struct {
	Service IService
	Handler IHandler
}
