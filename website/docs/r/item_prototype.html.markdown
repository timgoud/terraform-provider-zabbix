---
layout: "zabbix"
page_title: "Zabbix: zabbix_item_prototype"
sidebar_current: "docs-zabbix-resource-item-prototype"
description: |-
  Provides a zabbix item_prototype resource. This can be used to create and manage Zabbix Item prototype.
---

# zabbix_item_prototype

Provides a zabbix item_prototype resource. This can be used to create and manage Zabbix Item prototype.

## Example Usage

Create a new item

```hcl
resource "zabbix_host_group" "demo_group" {
  name = "Template demo group"
}

resource "zabbix_template" "demo_template" {
  host        = "template"
  name        = "template demo"
  description = "An exemple of template with lld_rule"
  groups      = [zabbix_host_group.demo_group.name]
}

resource "zabbix_lld_rule" "demo_lld_rule" {
    delay = 300
    host_id = zabbix_template.demo_template.id
    interface_id = "0"
    key = "demo.lld.rule"
    name = "demo discovery rule"
    type = 0
    filter {
        condition {
            macro = "{#FSTYPE}"
            value = "@fs"
        }
        eval_type = 0
    }
}

resource "zabbix_item_prototype" "demo_item_prototype" {
  delay = 60
  host_id  = zabbix_template.demo_template.id
  rule_id = zabbix_lld_rule.demo_lld_rule.id
  interface_id = "0"
  key = "demo.key"
  name = "demo item prototype"
}
```

## Argument Reference

The following arguments are supported:

* `host_id` - (Required) ID of the host or template that the item belongs to.
* `delay` - (Required) Update interval of the item. Accepts seconds or a time unit with suffix (30s,1m,2h,1d).
* `key` - (Required) Item key.
* `name` - (Required) Name of the item.
* `rule_id` - (Required) ID of the LLD rule that the item belongs to.
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
* `status` - (Optional)Status of the item prototype.
Possible values:
0 - (default) enabled item prototype
1 - disabled item prototype
3 - unsupported item prototype.




## Import

item prototype can be imported using their id, e.g.

```
$ terraform import zabbix_item_prototype.new_item 123456
```
