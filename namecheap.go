package namecheap

import (
	namecheap "github.com/c10l/libdns-namecheap"
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
)

// Provider wraps the provider implementation as a Caddy module.
type Provider struct{ *namecheap.Provider }

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "dns.providers.namecheap",
		New: func() caddy.Module {
			return &Provider{new(namecheap.Provider)}
		},
	}
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
// namecheap {
//     api_key <api_key>
//     api_user <api_user>
// }
//
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	repl := caddy.NewReplacer()
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "api_key":
				if d.NextArg() {
					p.Provider.APIKey = repl.ReplaceAll(d.Val(), "")
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "api_user":
				if d.NextArg() {
					p.Provider.APIUser = repl.ReplaceAll(d.Val(), "")
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.APIKey == "" {
		return d.Err("missing API key")
	}
	if p.Provider.APIUser == "" {
		return d.Err("missing API user")
	}
	return nil
}

// Interface guard
var _ caddyfile.Unmarshaler = (*Provider)(nil)
