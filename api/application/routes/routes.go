package routes

import ( 
	"net/http"
	"github.com/dicknaniel/Chillowance/api/application/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"CREATEFAMILY",
		"POST",
		"/family",
		Family_Create,
	},
}
