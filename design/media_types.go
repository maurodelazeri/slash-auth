package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Account is the account resource media type.
var Account = MediaType("application/ft.account+json", func() {
	Description("A user account")
	Attributes(func() {
		Attribute("id", Integer, "ID of the user", func() {
			Example(1)
		})
		Attribute("first_name", String, "First name of the user", func() {
			Example("James")
		})
		Attribute("last_name", String, "Last name of the user", func() {
			Example("Bond")
		})
		Attribute("email", String, "Email of the user", func() {
			Example("james@bond.com")
		})
		Required("id", "first_name", "last_name", "email")
	})

	View("default", func() {
		Attribute("id")
		Attribute("first_name")
		Attribute("last_name")
		Attribute("email")
	})
})

var Token = MediaType("application/ft.token+json", func() {
	Description("A JWT for a user")
	Attributes(func() {
		Attribute("token", String, "A JWT token")
		Required("token")
	})

	View("default", func() {
		Attribute("token")
	})
})
