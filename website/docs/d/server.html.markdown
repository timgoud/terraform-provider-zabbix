---
layout: "zabbix"
page_title: "Zabbix: zabbix_server"
sidebar_current: "docs-zabbix-data-source-server"
description: |-
  Provides a Zabbix Server data source. This can be used to get information about the Zabbix Server.
---

# zabbix_server

Provides a zabbix server data source. This can be used to get information about the Zabbix Server.

## Example Usage

Get the version of a Zabbix Server

```hcl
data "zabbix_server" "demo_server" {
}

output "version" {
  value = data.zabbix_server.demo_server.server_version
}
```

## Argument Reference

The following arguments are supported:

* `compare_version` - (Optional) Version to compare the Zabbix server version to.

## Attributes

* `server_version` - Version of the Zabbix server.
* `server_version_gt` - Returns true if the version of the Zabbix server is strictly greater than the version provided in `compare_version`.
* `server_version_lt` - Returns true if the version of the Zabbix server is strictly less than the version provided in `compare_version`.
* `server_version_ge` - Returns true if the version of the Zabbix server is greater or equal to the version provided in `compare_version`.
* `server_version_le` - Returns true if the version of the Zabbix server is less or equal to the version provided in `compare_version`.
* `unit_time_days` - String representation of the days unit used by the Zabbix server (`d` for version 3.4+, empty string otherwise).
* `unit_time_hours` - String representation of the hours unit used by the Zabbix server (`h` for version 3.4+, empty string otherwise).
* `unit_time_minutes` - String representation of the minutes unit used by the Zabbix server (`m` for version 3.4+, empty string otherwise).
* `unit_time_seconds` - String representation of the seconds unit used by the Zabbix server (`s` for version 3.4+, empty string otherwise).
* `unit_time_weeks` - String representation of the weeks unit used by the Zabbix server (`w` for version 3.4+, empty string otherwise).
