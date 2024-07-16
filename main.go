package main

import (
	"net/http"
	"projeto/routes"

)





func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
