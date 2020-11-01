package main

import "sandbox-echo/router"

func main() {
	e := router.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
