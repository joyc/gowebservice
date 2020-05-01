package main

import (
	"net/http"

	"github.com/joyc/gowebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
