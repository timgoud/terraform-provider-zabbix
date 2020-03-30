---
layout: "zabbix"
page_title: "Zabbix: zabbix_template"
sidebar_current: "docs-zabbix-resource-template"
description: |-
  Provides a zabbix template resource. This can be used to create and manage Zabbix Template.
---

# zabbix_template

A [template](https://www.zabbix.com/documentation/current/manual/api/reference/template) is a set of entities that can be conveniently applied to multiple hosts. 

## Example Usage

Create a new template

```hcl
resource "zabbix_template" "demo_template" {
  host        = "demo template"
  groups      = ["Discovered hosts"]
  description = "A basic template"
  macro = {
    EXAMPLE = "85"
  }
}
```

## Argument Reference

The following arguments are supported:

* `host` - (Required) Technical name of the template.
* `group` - (Required) Host group list of the template.
* `name` - (Optional) Display name of the template.
* `description` - (Optional) Description of the template.
* `macro` - (Optional) Template macro list .

## Import

Templates can be imported using their id, e.g.

```
$ terraform import zabbix_template.new_template 123456
```
