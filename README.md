# Redirector

Redirect is just a simple Go web application that will redirect clients to a new location based on their host header. It could be used to have a misspelled domain redirect to the correct spelling, provide a vanity URL for a longer URL, etc.

## Installation

```
go get github.com/yzguy/redirector
```

## Usage

A config file (`config.yaml`) is used to configure the server, as well as domain redirects. See below for an example

```yaml
---
server:
    address: "0.0.0.0"
    port: 8080

redirects:
  site1.domain.com:
    to: https://www.yahoo.com/
    with: 302

  site2.domain.com:
    to: https://www.google.com/
    with: 301
```

A domain is specified, eg. `site1.domain.com`, with `to` and `with` keys, specify where to redirect the user to, and with what HTTP Response code (`301` or `302`)
