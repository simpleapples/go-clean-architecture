package api

type IAPIAdapter interface {
	RunService(port int) error
	InitHandlers()
}

func NewAPIAdapater() IAPIAdapter {
	return newHTTPAPIAdapter()
}
