Loopia DNS module for Caddy
===========================
This module contains a DNS provider package for Caddy.
It can be used to manage DNS records with 
the [Loopia API](https://www.loopia.se/api/).


[![Loopia API](https://static.loopia.se/loopiaweb/images/logos/loopia-api-logo.png)](https://www.loopia.se/api/)

__WARNING:__ This will only properly work if you set `propagation_timeout`.
Loopia can use __up to 15 minutes__, but usually less, to propagate 
the changes so a high enough timeout is needed, default is 2 minutes.
You can set the [propagation_timeout](https://caddyserver.com/docs/caddyfile/directives/tls#acme) value
to compensate for this.

### Loopia API account requirements
You will need to [create a API user account](https://www.loopia.se/api/authentication/) 
that is separate from the normal Loopia user account. 

The API user will need to have access to the following methods.
- getSubdomains
- addSubdomain
- removeSubdomain
- getZoneRecords
- addZoneRecord
- removeZoneRecord

## Caddy module name
```
dns.providers.loopia
```

## Config examples
This can be used together with the ACME DNS challenge.

### Caddyfile
You can put the config in a [Caddyfile tls block](https://caddyserver.com/docs/caddyfile/directives/tls).
Note that caddyserver/caddy#4177 must be solved before `propagation_timeout` is supported in Caddyfile.

```
tls {
    issuer acme {
        propagation_timeout "15m0s"
        dns loopia {
                username "<your user>@loopiaapi"
                password "<your password>"
        }
    }
}
```
### Full example Caddyfile
```
somehost.example.org
tls {
        issuer acme {
                email "<your email for acme notifications>"
                propagation_timeout "15m0s"
                dns loopia {
                        username "<api user>@loopiaapi"
                        password "<api password>"
                }
        }
}
respond "Hello, world!"
```


### JSON
Using JSON notation with the [ACME module](https://caddyserver.com/docs/json/apps/tls/automation/policies/issuers/acme/).
```json
{
    "module": "acme",
    "challenges": {
        "dns": {
            "propagation_timeout": "15m0s",
            "provider": {
                "name": "loopia",
                "username": "<your user>@loopiaapi",
                "password": "<your password>"
            }
        }
    }
}
```

