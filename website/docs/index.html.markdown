---
layout: "zabbix"
page_title: "Provider: Zabbix"
sidebar_current: "docs-zabbix-index"
description: |-
  The Zabbix provider is used to interact with the resources supported by Zabbix API. The provider needs to be configured with the proper credentials before it can be used.
---

# Zabbix Provider

The [Zabbix](https://www.zabbix.com) provider is used to interact with the resources supported
by Zabbix API. The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Datadog provider
provider "zabbix" {
  user       = var.user
  password   = var.password
  server_url = var.server_url
}

# Create a new host
resource "zabbix_host" "default" {
  # ...
}

# Create a new timeboard
resource "zabbix_template" "default" {
  # ...
}
```

## Argument Reference

The following arguments are supported:

* `user` - (Required) Zabbix username. This can also be set via the `ZABBIX_USER` environment variable.
* `password` - (Required) Zabbix user password. This can also be set via the `ZABBIX_PASSWORD` environment variable.
* `server_url` - (Required) The API Url. This can be also be set via the `ZABBIX_SERVER_URL` environment variable. Note that this URL must point to `api_jsonrpc.php`. For example `http://localhost/api_jsonrpc.php`.
* `tls_insecure` - (Optional) Set to `true` for skipping verification of TLS certificates. Also can be set via `ZABBIX_TLS_INSECURE`.