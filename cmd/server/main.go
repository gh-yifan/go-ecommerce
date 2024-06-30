package server

import "github.com/yifan-guo/go-ecommerce/pkg/api"

func main() {
	router := api.SetupRouter()
	router.Run()
}
