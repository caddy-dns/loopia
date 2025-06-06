package loopia

import (
	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	libdnsloopia "github.com/libdns/loopia"
)

type Provider struct {
	*libdnsloopia.Provider
	Logging bool `json:"logging,omitempty"` // Enable logging
}

func init() {
	caddy.RegisterModule(Provider{})
}

// CaddyModule returns the Caddy module information.
func (Provider) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "dns.providers.loopia",
		New: func() caddy.Module { return &Provider{Provider: new(libdnsloopia.Provider)} },
	}
}

// Provision sets up the module. Implements caddy.Provisioner.
func (p *Provider) Provision(ctx caddy.Context) error {
	repl := caddy.NewReplacer()

	if p.Logging {
		p.Provider.SetLogger(ctx.Logger(p).Sugar())
	}
	p.Provider.Username = repl.ReplaceAll(p.Provider.Username, "")
	p.Provider.Password = repl.ReplaceAll(p.Provider.Password, "")
	p.Provider.Customer = repl.ReplaceAll(p.Provider.Customer, "")
	return nil
}

// UnmarshalCaddyfile sets up the DNS provider from Caddyfile tokens. Syntax:
//
//	loopia [<username> <password>] {
//	    username <username>
//	    password <password>
//	    customer <customer no>
//	}
func (p *Provider) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			p.Provider.Username = d.Val()
		}
		if d.NextArg() {
			p.Provider.Password = d.Val()
		}
		if d.NextArg() {
			return d.ArgErr()
		}
		for nesting := d.Nesting(); d.NextBlock(nesting); {
			switch d.Val() {
			case "username":
				if p.Provider.Username != "" {
					return d.Err("User already set")
				}
				if d.NextArg() {
					p.Provider.Username = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "password":
				if p.Provider.Password != "" {
					return d.Err("User already set")
				}
				if d.NextArg() {
					p.Provider.Password = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "customer":
				if p.Provider.Customer != "" {
					return d.Err("Customer already set")
				}
				if d.NextArg() {
					p.Provider.Customer = d.Val()
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			case "logging":
				if p.Logging {
					return d.Err("Logging already set")
				}
				if d.NextArg() {
					// Check if the value is a boolean
					// and set the Logging field accordingly
					// If not, return an error
					switch d.Val() {
					case "true", "yes", "on", "1":
						p.Logging = true
					case "false", "no", "off", "0":
						p.Logging = false
					default:
						return d.Errf("unrecognized logging value '%s'", d.Val())
					}
				}
				if d.NextArg() {
					return d.ArgErr()
				}
			default:
				return d.Errf("unrecognized subdirective '%s'", d.Val())
			}
		}
	}
	if p.Provider.Username == "" {
		return d.Err("missing user")
	}
	if p.Provider.Password == "" {
		return d.Err("missing password")
	}
	return nil
}

// Interface guards
var (
	_ caddyfile.Unmarshaler = (*Provider)(nil)
	_ caddy.Provisioner     = (*Provider)(nil)
)
