package docs

var components = map[string]interface{}{
	"securitySchemes": map[string]interface{}{
		"bearerAuth": map[string]interface{}{
			"type":         "http",
			"scheme":       "bearer",
			"bearerFormat": "JWT",
		},
	},
}
