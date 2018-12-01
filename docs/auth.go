package docs

var authApi = map[string]Doc{
	"post": {
		Tags:        []string{"auth"},
		Summary:     "signIn",
		OperationId: "authenticate user",
		Consumes:    []string{"application/json"},
		Produces:    []string{"*/*"},
		Parameters: []Parameter{
			{
				In:          "body",
				Name:        "request",
				Description: "request",
				Required:    true,
				Schema: map[string]interface{}{
					"type": "object",
					"properties": map[string]struct{ Type string `json:"type"` }{
						"username": {"string"},
						"password": {"string"},
					},
				},
			},
		},
		Responses: map[string]Response{
			"200": {Description: "Success"},
			"400": {Description: "BadRequest"},
			"404": {Description: "NotFound"},
			"422": {Description: "Unprocessable Entity"},
		},
	},
}

var refreshApi = map[string]Doc{
	"get": {
		Tags:        []string{"auth"},
		Summary:     "refresh",
		OperationId: "refresh token",
		Consumes:    []string{"application/json"},
		Produces:    []string{"*/*"},
		Responses: map[string]Response{
			"200": {Description: "Success", Schema: map[string]interface{}{
				"type": "object",
				"properties": map[string]struct{ Type string `json:"type"` }{
					"access":  {"string"},
					"refresh": {"string"},
				},
			}},
			"422": {Description: "Unproccessable Entity"},
		},
	},
}
