package design

import (
	. "github.com/goadesign/goa/design" // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl"
)

// API(apiname, anonFxntoDefineProperties)
var _ = API("cellar", func() {
	Description("The Wine Review Service")
	Host("localhost:8080")
})

// BottlePayload is defining input types, (request to API)
var BottlePayload = Type("BottlePayload", func() {
	Description("BottlePayload is the type used")

	Attribute("name", String, "Name of bottle", func() {
		MinLength(2)
	})

	Attribute("vintage", Integer, "Vintage of bottle", func() {
		Minimum(1900)
	})

	Attribute("rating", Integer, "Critic's rating of bottle", func() {
		Minimum(1)
		Maximum(5)
	})
	Required("name", "vintage", "rating")
})

// BottleMedia defining the output type (response of API).
// Can also define the view on it.
var BottleMedia = MediaType("application/vnd.gophercon.goa.bottle", func() {
	TypeName("bottle")
	Reference(BottlePayload) // some of the media type attrs are reused from payload above

	Attributes(func() {
		Attribute("ID", Integer, "Unique bottle ID") // specific to this
		Attribute("name")                            // non-specific
		Attribute("vintage")
		Attribute("rating")
		// any rendering of resource must have non-zero values for these
		Required("ID", "name", "vintage", "rating")
	})

	View("default", func() {
		Attribute("ID")
		Attribute("name")
		Attribute("vintage")
		Attribute("rating")
	})
})

// actions are endpoints

var _ = Resource("bottle", func() {
	Description("A wine bottle")
	BasePath("/bottles")

	Action("create", func() {
		Description("creates a bottle")
		Routing(POST("/"))
		Payload(BottlePayload)
		Response(Created)
	})

	Action("show", func() {
		Description("shows a bottle")
		Routing(GET("/:id"))
		Params(func() {
			Param("id", Integer)
		})
		Response(OK, BottleMedia)
	})
})
