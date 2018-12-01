package docs

var postApi = map[string]Doc{
	"post": {
		Tags:        []string{"posts"},
		Summary:     "create post",
		OperationId: "create post",
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
						"title":   {"string"},
						"content": {"string"},
					},
				},
			},
		},
		Responses: map[string]Response{
			"201": {Description: "Created"},
			"400": {Description: "Bad Request"},
			"401": {Description: "Unauthorized"},
			"422": {Description: "Unprocessable Entity"},
		},
	},
	"get": {
		Tags:        []string{"posts"},
		Summary:     "getAllPosts",
		OperationId: "get all posts",
		Consumes:    []string{"application/json"},
		Produces:    []string{"*/*"},
		Parameters:  []Parameter{},
		Responses: map[string]Response{
			"200": {
				Description: "Success", Schema: map[string]interface{}{
					"type": "array",
					"items": map[string]interface{}{
						"type": "object",
						"properties": map[string]map[string]interface{}{
							"post_id":   {"type": "integer"},
							"author":    {"type": "string"},
							"title":     {"type": "string"},
							"content":   {"type": "string"},
							"create_at": {"type": "string"},
							"update_at": {"type": "string"},
							"comments": {
								"type": "array",
								"items": map[string]interface{}{
									"type": "object",
									"properties": map[string]struct{ Type string `json:"type"` }{
										"id":        {"integer"},
										"content":   {"string"},
										"author":    {"string"},
										"create_at": {"string"},
										"update_at": {"string"},
									},
								},
							},
						},
					},
				},
			},
		},
	},
}
