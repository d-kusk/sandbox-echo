package main

import "github.com/d-kusk/sandbox-echo/route"

func main() {
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
