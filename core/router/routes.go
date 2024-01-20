package router

var Routes []Route

var (
	Get    = Create("GET")
	Post   = Create("POST")
	Patch  = Create("PATCH")
	Put    = Create("PUT")
	Delete = Create("DELETE")
)
