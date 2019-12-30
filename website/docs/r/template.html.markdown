---
layout: "zabbix"
page_title: "Provider: ZAbbix"
sidebar_current: "docs-zabbix-template"
description: |-
  The Zabbix provider is used to interact with zabbix resource. The provider needs to be configured with the proper credentials before it can be used.
---

# zabbix_template

Provides a zabbix template resource. This can be used to create and manage Zabbix Template.

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

template can be imported using their id, e.g.

```
$ terraform import zabbix_template.new_template 123456
```
