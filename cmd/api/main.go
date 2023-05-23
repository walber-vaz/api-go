package main

import "github.com/walber-vaz/api-go/cmd/api/routers"

func main() {
	app := routers.RouterHandler()

	app.Run(":9000")
}
