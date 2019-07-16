package design

import (
	. "goa.design/goa/v3/dsl"
)

var PetCreatePayload = Type("PetCreatePayload", func() {
	Attribute("name", String, "Name")
	Attribute("species", String, "Species")
	Attribute("user_id", String, "User ID")

	Required("user_id", "name")
})

var PetUpdatePayload = Type("PetUpdatePayload", func() {
	Attribute("id", Int, "Pet's ID")
	Attribute("name", String, "Name")
	Attribute("species", String, "Species")

	Required("id")
})

var Pet = Type("pet", func() {
	Attribute("name", String, "Name")
	Attribute("species", String, "Species")
	Attribute("user_id", String, "User ID")
})

var _ = Service("pets", func() {
	Description("CRUD Pets")
	HTTP(func() {
		Path("/pets")
	})

	Method("create", func() {
		Payload(PetCreatePayload)
		Result(Pet)
		HTTP(func() {
			POST("/")
			Response(StatusOK)
		})
	})

	Method("read", func() {
		Payload(func() {
			Attribute("id", Int, "Pet's database ID")
		})

		Result(Pet)
		HTTP(func() {
			GET("/{id}")
			Response(StatusOK)
		})
	})

	Method("read_all", func() {
		Result(ArrayOf(Pet))
		HTTP(func() {
			GET("/")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Payload(PetUpdatePayload)
		HTTP(func() {
			PATCH("/{id}")
			Response(StatusNoContent)
		})
	})

	Method("delete", func() {
		Payload(func() {
			Attribute("id", Int, "Pet's database ID")
		})

		HTTP(func() {
			DELETE("/{id}")
			Response(StatusOK)
		})
	})
})
