package main

import "github.com/demirbey05/golang-starter/internals/routers"

func main() {
	e := routers.InitRoutes()
	e.Logger.Fatal(e.Start(":8080"))
}
