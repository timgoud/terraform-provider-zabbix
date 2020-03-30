---
layout: "zabbix"
page_title: "Zabbix: zabbix_trigger"
sidebar_current: "docs-zabbix-resource-trigger"
description: |-
  Provides a zabbix trigger resource. This can be used to create and manage Zabbix trigger.
---

# zabbix_trigger

Provides a zabbix trigger resource. This can be used to create and manage Zabbix trigger.

## Example Usage

Create a new trigger

```hcl
resource "zabbix_host_group" "demo_group" {
  name = "Template demo group"
}

resource "zabbix_template" "demo_template" {
  host        = "template"
  name        = "template demo"
  description = "An exemple of template with item and trigger"
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

resource "zabbix_trigger" "demo_trigger" {
  description = "demo trigger"
  expression  = "{${zabbix_template.demo_template.host}:${zabbix_item.demo_item.key}.last()}=0"
  priority    = 5
 status      = 0
}
```

Create two trigger with one dependencies
```hcl
resource "zabbix_host_group" "demo_group" {
  name = "Template demo group"
}

resource "zabbix_template" "demo_template" {
  host        = "template"
  name        = "template demo"
  description = "An exemple of template with item and trigger"
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

resource "zabbix_trigger" "demo_trigger" {
  description = "demo trigger"
  expression  = "{${zabbix_template.demo_template.host}:${zabbix_item.demo_item.key}.last()}=0"
  priority    = 5
 status      = 0
}

resource "zabbix_trigger" "demo_trigger" {
  description = "demo trigger"
  expression  = "{${zabbix_template.demo_template.host}:${zabbix_item.demo_item.key}.last()}=0"
  dependencies = [
      zabbix_trigger.demo_trigger.id,
  ]
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) Name of the trigger.
* `expression` - (Required) Expand expression of the trigger.
* `comment` - (Optional) Additional description of ther trigger.
* `priority` - (Optional) Severity of the trigger.
Possible values are:
0 - (default) not classified
1 - information
2 - warning
3 - average
4 - high
5 - disaster.
* `status` - (Optional) Whether the trigger is enabled or disabled.
Possible values are:
0 - (default) enabled
1 - disabled.
* `dependencies` - (Optional) Triggers id that the trigger is dependent on.

## Import

trigger can be imported using their id, e.g.

```
$ terraform import zabbix_trigger.new_trigger 123456
```
