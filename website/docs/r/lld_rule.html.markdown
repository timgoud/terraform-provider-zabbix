---
layout: "zabbix"
page_title: "Zabbix: zabbix_lld_rule"
sidebar_current: "docs-zabbix-resource-lld-rule"
description: |-
  Provides a zabbix lld_rule resource. This can be used to create and manage Zabbix low level discovery rule.
---

# zabbix_lld_rule

[Low-level discovery](https://www.zabbix.com/documentation/current/manual/api/reference/discoveryrule) discovery provides a way to automatically create items, triggers, and graphs for different entities on a host.

## Example Usage

Create a new low level discovery rule

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
```

## Argument Reference

The following arguments are supported:

* `delay` - (Required) Update interval of the LLD rule in seconds.
* `host_id` - (Required) ID of the host that the LLD rule belongs to.
* `interface_id` - (Required) ID of the LLD rule's host interface. Used only for host LLD rules. Optional for Zabbix agent (active), Zabbix internal, Zabbix trapper and database monitor LLD rules.
* `key` - (Required) LLD rule key.
* `name` - (Required) Name of the LLD rule.
* `type` - (Required) Type of the LLD rule. Can be `0` (Zabbix agent), `1` (SNMPv1 agent), `2` (Zabbix trapper), `3` (simple check), `4` (SNMPv2 agent), `5` (Zabbix internal), `6` (SNMPv3 agent), `7` (Zabbix agent active), `8` (Zabbix aggregate), `9` (web item), `10` (external check), `11` (database monitor), `12` (IPMI agent), `13` (SSH agent), `14` (TELNET agent), `15` (calculated), `16` (JMX agent).
* `filter` - (Required) LLD rule filter object for the LLD rule.
    * `condition` - (Required) Set of filter conditions to use for filtering results. Multiple `condition` are allowed.
        * `macro` - (Required) LLD macro to perform the check on.
        * `value` - (Required) Value to compare with.
        * `operator` - (Optional) Condition operator.
Possible values:
8 - (default) matches regular expression.
    * `eval_type` - (Required) Filter condition evaluation method.
Possible values:
0 - and/or
1 - and
2 - or
3 - custom expression.
    * `formula` - (Optional) User-defined expression to be used for evaluating conditions of filters with a custom expression. The expression must contain IDs that reference specific filter conditions by its formulaid. The IDs used in the expression must exactly match the ones defined in the filter conditions: no condition can remainunused or omitted.
Required for custom expression filters.

## Import

LLD rules can be imported using their id, e.g.

```
$ terraform import zabbix_lld_rule.new_lld_rule 123456
```
