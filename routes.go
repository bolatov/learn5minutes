package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// Route{
	// 	"Index",
	// 	"GET",
	// 	"/",
	// 	Index,
	// },
	Route{
		"AlgoIndex",
		"GET",
		"/",
		AlgoIndex,
	},
	Route{
		"AlgoShow",
		"GET",
		"/algo/{algoId}",
		AlgoShow,
	},
	Route{
		"AlgoCreate",
		"POST",
		"/algo",
		AlgoCreate,
	},
	Route{
		"AlgoUpdate",
		"PUT",
		"/algo",
		AlgoUpdate,
	},
	Route{
		"AlgoRandom",
		"GET",
		"/random",
		AlgoRandom,
	},
}
