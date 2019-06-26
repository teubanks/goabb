package design

import (
	. "goa.design/goa/v3/dsl"
)

var UserCreatePayload = Type("UserCreatePayload", func() {
	Attribute("first_name", String, "First Name")
	Attribute("last_name", String, "Last Name")
	Attribute("email", String, "Email")
	Attribute("age", Int, "How old")
	Attribute("birthday", String, "Date of Birth")

	Required("email")
})

var User = Type("user", func() {
	Attribute("first_name", String, "First Name")
	Attribute("last_name", String, "Last Name")
	Attribute("email", String, "Email")
	Attribute("age", Int, "How old")
	Attribute("birthday", String, "Date of Birth")
})

var _ = Service("users", func() {
	Description("CRUD users")

	Method("create", func() {
		Payload(UserCreatePayload)
		// Result describes the method result
		// Here the result is a simple integer value
		Result(User)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			POST("/users")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusCreated)
		})
	})

	Method("read_all", func() {
		// Result describes the method result
		// Here the result is a simple integer value
		Result(ArrayOf(User))
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/users")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})

	Method("read", func() {
		Payload(func() {
			Attribute("id", Int, "User's database ID")
		})
		// Result describes the method result
		// Here the result is a simple integer value
		Result(User)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			GET("/users/{id}")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(StatusOK)
		})
	})
})
