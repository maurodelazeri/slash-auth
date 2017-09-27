package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Slash API", func() {
	Title("The Freshworks Hitback")
	Description("Authentication API")
	Contact(func() {
		Name("Mauro Delazeri")
		Email("mauro@hitback.us")
	})
	Host("localhost:8080")
	Scheme("http")
	BasePath("/")
	Consumes("application/json")
	Produces("application/json")
})
