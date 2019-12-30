---
layout: "zabbix"
page_title: "Provider: Zabbix"
sidebar_current: "docs-zabbix-index"
description: |-
  The Zabbix provider is used to interact with zabbix resource. The provider needs to be configured with the proper credentials before it can be used.
---

# zabbix_item

Provides a zabbix item resource. This can be used to create and manage Zabbix Item.

## Example Usage

Create a new item

```hcl
resource "zabbix_host_group" "demo_group" {
  name = "Template demo group"
}

resource "zabbix_template" "demo_template" {
  host        = "template"
  name        = "template demo"
  description = "An exemple of template with item"
  groups      = [zabbix_host_group.demo_group.name]
}

resource "zabbix_item" "demo_item" {
  name        = "demo item"
  key         = "demo.key"
  delay       = "34"
  description = "Item for the demo template"
  trends      = "300"
  history     = "25"
  host_id     = zabbix_template.demo_template.id
}
```

## Argument Reference

The following arguments are supported:

* `host_id` - (Required) ID of the host or template that the item belongs to.
* `delay` - (Required) Update interval of the item. Accepts seconds or a time unit with suffix (30s,1m,2h,1d).
* `key` - (Required) Item key.
* `name` - (Required) Name of the item.
* `type` - (Required) Type of the item.
Possible values:
0 - Zabbix agent
1 - SNMPv1 agent
2 - Zabbix trapper
3 - simple check
4 - SNMPv2 agent
5 - Zabbix internal
6 - SNMPv3 agent
7 - Zabbix agent (active)
8 - Zabbix aggregate
9 - web item
10 - external check
11 - database monitor
12 - IPMI agent
13 - SSH agent
14 - TELNET agent
15 - calculated
16 - JMX agent
17 - SNMP trap
18 - Dependent item
19 - HTTP agent
* `value_type` - (Required) Type of information of the item.
Possible values:
0 - numeric float
1 - character
2 - log
3 - numeric unsigned
4 - text
* `interface_id` - (Optional)  ID of the item's host interface.
Not required for template items. Optional for internal, active agent, trapper, aggregate, calculated, dependent and database monitor items.
* `data_type` - (Optional, Remove in Zabbix 3.4) Data type of the item.
Possible values:
0 - (default) decimal
1 - octal
2 - hexadecimal
3 - boolean
* `delta` - (Optional, Remove in Zabbix 3.4) Value that will be stored.
Possible values:
0 - (default) as is
1 - Delta, speed per second
2 - Delta, simple change
* `description` - (Optional) Description of the item.
* `history` - (Optional) Number of days to keep item's history data. Default 90.
* `trends` - (Optional)	Number of days to keep item's trends data. Default: 365.
* `trapper_host` - (Optional) Allowed hosts. Used only by trapper items.




## Import

item can be imported using their id, e.g.

```
$ terraform import zabbix_item.new_item 123456
```
