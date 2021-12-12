package main

import "clean_architecutre_demo/adapters/api"

func main() {
	apiAdapter := api.NewAPIAdapater()
	apiAdapter.InitHandlers()
	apiAdapter.RunService(9000)
}
