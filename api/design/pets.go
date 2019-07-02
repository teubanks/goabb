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

	Method("create", func() {
		Payload(PetCreatePayload)
		Result(Pet)
		HTTP(func() {
			POST("pets")
			Response(StatusOK)
		})
	})

	Method("read", func() {
		Payload(func() {
			Attribute("id", Int, "Pet's database ID")
		})

		Result(Pet)
		HTTP(func() {
			GET("pets/{id}")
			Response(StatusOK)
		})
	})

	Method("read_all", func() {
		Result(ArrayOf(Pet))
		HTTP(func() {
			GET("pets")
			Response(StatusOK)
		})
	})

	Method("update", func() {
		Payload(PetUpdatePayload)
		HTTP(func() {
			PATCH("pets/{id}")
			Response(StatusNoContent)
		})
	})

	Method("delete", func() {
		Payload(func() {
			Attribute("id", Int, "Pet's database ID")
		})

		HTTP(func() {
			DELETE("pets/{id}")
			Response(StatusOK)
		})
	})
})
