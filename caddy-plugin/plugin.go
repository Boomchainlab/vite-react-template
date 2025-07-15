package plugin

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(MyPlugin{})
}

type MyPlugin struct{}

func (MyPlugin) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.myplugin",
		New: func() caddy.Module { return new(MyPlugin) },
	}
}

func (m MyPlugin) Provision(ctx caddy.Context) error {
	return nil
}

func (m MyPlugin) Validate() error {
	return nil
}

func (m MyPlugin) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	w.Header().Set("X-Custom-Boomchain", "Boomchainlab.blog served by Caddy!")
	return next.ServeHTTP(w, r)
}
