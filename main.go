package main

import "github.com/dsk52/sandbox-echo/router"

func main() {
	e := router.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
