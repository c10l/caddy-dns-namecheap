Namecheap module for Caddy
===========================

This package contains a DNS provider module for [Caddy](https://github.com/caddyserver/caddy). It can be used to manage DNS records with Namecheap accounts.

## Caddy module name

```
dns.providers.namecheap
```

## Config examples

To use this module for the ACME DNS challenge, [configure the ACME issuer in your Caddy JSON](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuer/acme/) like so:

```json
{
	"module": "acme",
	"challenges": {
		"dns": {
			"provider": {
				"name": "namecheap",
				"api_key": "NAMECHEAP_API_KEY"
				"api_user": "NAMECHEAP_API_USER"
			}
		}
	}
}
```

or with the Caddyfile:

```
tls {
	dns namecheap {
    api_key env.NAMECHEAP_API_KEY
    api_user env.NAMECHEAP_API_USER
  }
}
```

You can replace `env.NAMECHEAP_API_KEY` and `env.NAMECHEAP_API_USER` with the actual auth token if you prefer to put it directly in your config instead of an environment variable.


## Authenticating

See [the associated README in the libdns package](https://github.com/c10l/libdns-namecheap) for important information about credentials.
