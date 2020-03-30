---
layout: "zabbix"
page_title: "Zabbix: zabbix_template_link"
sidebar_current: "docs-zabbix-resource-template-link"
description: |-
  Provider a virtual resource to track template dependencies such as item, trigger and low level discovery rule.
---

# zabbix_template_link

Provider a virtual resource to track template dependencies such as item, trigger and low level discovery rule.

## Example Usage

Create a new template link

```hcl
resource "zabbix_template" "demo_template" {
  host        = "demo template"
  groups      = ["Discovered hosts"]
  description = "A basic template"
  macro = {
    EXAMPLE = "85"
  }
}

resource "zabbix_template_link" "demo_template_link" {
  template_id = zabbix_template.demo_template_link.id
}
```

Create a template link to track one item
```hcl
resource "zabbix_template" "demo_template" {
  host        = "demo template"
  groups      = ["Discovered hosts"]
  description = "A basic template"
  macro = {
    EXAMPLE = "85"
  }
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

resource "zabbix_template_link" "demo_template_link" {
  template_id = zabbix_template.demo_template_link.id
  item {
      item_id = zabbix_item.
  }
}
```

## Argument Reference

The following arguments are supported:

* `template_id` - (Required) Id of the template.
* `item` - (Optional) Use to track template's item. Item can be used multiple time.
    * `item_id` - (Required) id of the track item.
* `trigger` - (Optional) Use to track template's trigger. Trigger can be used multiple time.
    * `trigger_id` - (Required) id of the track trigger.
* `lld_rule` - (Optional) Use to track template's low level discovery rule.
    * `lld_rule_id` - (Required) id of the track lld rule. lld_rule can be used multiple time.

## Import

template can be imported using the dependencies id, e.g.
```
$ terraform import zabbix_template_link.new_template_link TEMPLATEID_ITEMID_TRIGGERID_LLDRULEID
```

```
$ terraform import zabbix_template_link.new_template_link 123456
```
