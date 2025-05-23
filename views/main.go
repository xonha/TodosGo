package views

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humabunrouter"
	"github.com/uptrace/bunrouter"
)

var (
	Router = bunrouter.New()
	api    huma.API
)

func Init() {
	Router = bunrouter.New()
	config := huma.DefaultConfig("API Todos", "1.0.0")
	config.DocsPath = ""
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"Bearer": {
			Type: "oauth2",
			Flows: &huma.OAuthFlows{
				AuthorizationCode: &huma.OAuthFlow{
					AuthorizationURL: "https://example.com/oauth/authorize",
					TokenURL:         "https://example.com/oauth/token",
					Scopes: map[string]string{
						"scope1": "Scope 1 description...",
						"scope2": "Scope 2 description...",
					},
				},
			},
		},
	}
	Router.GET("/", bunrouter.HTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!doctype html>
		<html>
			<head>
				<title>API Todos</title>
				<link rel="icon" href="https://upload.wikimedia.org/wikipedia/commons/6/67/Microsoft_To-Do_icon.png" type="image/png">
				<meta charset="utf-8" />
				<meta name="viewport" content="width=device-width, initial-scale=1" />
			</head>
			<body>
				<script id="api-reference" data-url="/openapi.json"></script>
				<script>
					var configuration = { theme: 'purple' }
					document.getElementById('api-reference').dataset.configuration =
						JSON.stringify(configuration)
				</script>
				<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
			</body>
		</html>`))
	}))

	api = humabunrouter.New(Router, config)
	todos()
}
