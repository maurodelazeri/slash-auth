package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("authentication", func() {
	BasePath("/auth")
	NoSecurity()
	Action("login", func() {
		NoSecurity()
		Routing(
			POST("/login"),
		)
		Description("Sign a user in")
		Payload(LoginPayload)
		Response(OK, Token)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

	Action("register", func() {
		NoSecurity()
		Routing(
			POST("/register"),
		)
		Description("Create a new user")
		Payload(AccountPayload)
		Response(NoContent, func() {
			Headers(func() {
				Header("Authorization", String, "Generated JWT")
			})
		})
		Response(OK, Token)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

})

var _ = Resource("account", func() {
	DefaultMedia(Account)
	BasePath("/account")
	Security(JWT)

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all users")
		Response(OK, CollectionOf(Account))
		Response(Unauthorized)
		Response(InternalServerError)
	})

	Action("show", func() {
		Routing(
			GET("/:accountID"),
		)
		Description("Retrieve a user with a given id")
		Params(func() {
			Param("accountID", Integer, "Account ID", func() {
				Minimum(1)
			})
		})
		Response(OK, Account)
		Response(Unauthorized)
		Response(NotFound)
	})

	Action("update", func() {
		Routing(
			PUT("/:accountID"),
		)
		Description("Change user data")
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Payload(UpdateAccountPayload)
		Response(OK)
		Response(NotFound)
		Response(Unauthorized)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:accountID"),
		)
		Params(func() {
			Param("accountID", Integer, "Account ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(Unauthorized)
		Response(InternalServerError)
		Response(BadRequest, ErrorMedia)
	})
})
