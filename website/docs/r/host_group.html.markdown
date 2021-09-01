---
layout: "zabbix"
page_title: "Zabbix: zabbix_host_group"
sidebar_current: "docs-zabbix-resource-host-group"
description: |-
  Provides a zabbix host group resource. This can be used to create and manage Zabbix Host Group.
---

# zabbix_host

An [host group](https://www.zabbix.com/documentation/current/manual/api/reference/hostgroup) is typically a group of hosts to monitor.

## Example Usage

Create a new host group

```hcl
resource "zabbix_host_group" "demo_group" {
  name = "Template demo group"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the host group.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `group_id` - The zabbix host group ID
