package docs

var userApi = map[string]Doc{
	"post": {
		Tags:        []string{"users"},
		Summary:     "signUp",
		OperationId: "create user",
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
						"nickname": {"string"},
						"email":    {"string"},
					},
				},
			},
		},
		Responses: map[string]Response{
			"201": {Description: "Created"},
			"400": {Description: "BadRequest"},
			"409": {Description: "Conflict"},
		},
		Deprecated: false,
	},
	"get": {
		Tags:        []string{"users"},
		Summary:     "info",
		OperationId: "get user info",
		Consumes:    []string{"application/json"},
		Produces:    []string{"*/*"},
		Parameters:  []Parameter{},
		Security: map[string]interface{}{
			"bearerAuth": []string{},
		},
		Responses: map[string]Response{
			"200": {
				Description: "Success", Schema: map[string]interface{}{
					"type": "object",
					"properties": map[string]struct{ Type string `json:"type"` }{
						"username": {"string"},
						"nickname": {"string"},
						"email":    {"string"},
					},
				},
			},
			"422": {Description: "Unprocessable Entity"},
		},
	},
}
