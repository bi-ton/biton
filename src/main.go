package main

import (
	"github.com/msw-x/moon/app"
)

func main() {
	app.Run[Conf](version, run)
}
