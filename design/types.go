package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var LoginPayload = Type("LoginPayload", func() {
	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("jamesbond@gmail.com")
	})

	Attribute("password", String, func() {
		MinLength(5)
		MaxLength(100)
		Example("abcd1234")
	})
	Required("email", "password")
})

var AccountPayload = Type("AccountPayload", func() {
	Attribute("first_name", String, func() {
		MinLength(2)
		MaxLength(200)
		Example("James")
	})

	Attribute("last_name", String, func() {
		MinLength(2)
		MaxLength(200)
		Example("Bond")
	})

	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("Bond")
	})

	Attribute("password", String, func() {
		MinLength(5)
		MaxLength(100)
		Example("abcd1234")
	})
	Required("first_name", "last_name", "email", "password")
})

var UpdateAccountPayload = Type("UpdateAccountPayload", func() {
	Attribute("first_name", String, func() {
		MinLength(2)
		MaxLength(200)
		Example("James")
	})

	Attribute("last_name", String, func() {
		MinLength(2)
		MaxLength(200)
		Example("Bond")
	})

	Attribute("email", String, func() {
		MinLength(6)
		MaxLength(400)
		Format("email")
		Example("Bond")
	})
})
