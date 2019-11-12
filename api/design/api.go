package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("GoaBB", func() {
	Title("An example of an API")
	Server("goabb", func() {
		Services("pets", "users")
		Host("localhost", func() {
			URI("http://localhost:8000")
			URI("grpc://localhost:8080")
		})
	})
})
