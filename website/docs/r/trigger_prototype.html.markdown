---
layout: "zabbix"
page_title: "Zabbix: zabbix_trigger_prototype"
sidebar_current: "docs-zabbix-resource-trigger-prototype"
description: |-
  Provides a zabbix trigger_prototype resource. This can be used to create and manage Zabbix trigger prototype.
---

# zabbix_trigger_prototype

Provides a zabbix trigger_prototype resource. This can be used to create and manage Zabbix trigger prototype.

## Example Usage

Create a new trigger

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

resource "zabbix_trigger_prototype" "trigger_prototype_demo" {
  description = "trigger prototype demo"
  expression = "demo.trigger.prototype"
  priority = 5
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) Name of the trigger.
* `expression` - (Required) Expand expression of the trigger.
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

trigger prototype can be imported using their id, e.g.

```
$ terraform import zabbix_trigger_prototype.new_trigger 123456
```
