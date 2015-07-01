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
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ChildrenIndex",
		"GET",
		"/children",
		ChildrenIndex,
	},
	Route{
		"ChildCreate",
		"POST",
		"/children",
		ChildCreate,
	},
	Route{
		"ChildShow",
		"GET",
		"/children/{childId}",
		ChildShow,
	},
	Route{
		"ChildShowChores",
		"GET",
		"/children/{childId}/chores",
		ChildShowChores,
	},
	Route{
		"ChildAddChore",
		"GET",
		"/children/{childId}/chores/{choreId}",
		ChildAddChore,
	},
	Route{
		"ChoreIndex",
		"GET",
		"/chores",
		ChoreIndex,
	},
	Route{
		"ChoreCreate",
		"POST",
		"/chores",
		ChoreCreate,
	},
	Route{
		"ChoreShow",
		"GET",
		"/chores/{choreId}",
		ChoreShow,
	},
}
