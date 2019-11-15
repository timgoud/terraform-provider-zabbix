# Terraform zabbix provider

Terraform provider for zabbix that allow you to manage : host, template, trigger and item. 

## Example
***


### Host example:
```hcl
provider "zabbix" {
  user = "Admin"
  password = "zabbix"
  server_url = "http://localhost/api_jsonrpc.php"
}

resource "zabbix_host" "zabbix1" {
  host = "127.0.0.1"
  name = "the best name"
  interfaces {
    ip = "127.0.0.1"
    main = true
  }
  groups = ["Linux servers", "${zabbix_host_group.zabbix.name}"]
  templates = ["Template ICMP Ping"]
}

resource "zabbix_host_group" "zabbix" {
  name = "something"
}
```
### Template example:
The template link resource is required if you want to track your template item and trigger
```hcl
resource "zabbix_host_group" "demo_group" {
  name = "Template demo group"
}

resource "zabbix_template" "demo_template" {
  host        = "template"
  name        = "template demo"
  description = "An exemple of template with item and trigger"
  groups      = [zabbix_host_group.demo_group.name]
  macro = {
    MACRO_TEMPLATE = "12"
  }
}

# This virtual resource is responsible of ensuring no other items are associated to the template
resource "zabbix_template_link" "demo_template_link" {
  template_id = zabbix_template.demo_template.id
}
```
### Template with item and trigger example:
```hcl
provider "zabbix" {
  user       = var.user
  password   = var.password
  server_url = "http://localhost/api_jsonrpc.php"
}

resource "zabbix_host_group" "demo_group" {
  name = "Template demo group"
}

resource "zabbix_template" "demo_template" {
  host        = "template"
  name        = "template demo"
  description = "An exemple of template with item and trigger"
  groups      = [zabbix_host_group.demo_group.name]
  macro = {
    MACRO_TEMPLATE = "12"
  }
}

resource "zabbix_item" "demo_item" {
  name        = "demo item"
  key         = "demo.key"
  delay       = "34"
  description = "Item for the demo template"
  trends      = "300"
  history     = "25"
  host_id     = zabbix_template.demo_template.template_id
}

resource "zabbix_trigger" "demo_trigger" {
  description = "demo trigger"
  expression  = "{${zabbix_template.demo_template.host}:${zabbix_item.demo_item.key}.last()}={$MACRO_TEMPLATE}"
  priority    = 5
  status      = 0
}

# This virtual resource is responsible of ensuring no other items are associated to the template
resource "zabbix_template_link" "demo_template_link" {
  template_id = zabbix_template.demo_template.id
  item {
    item_id = zabbix_item.demo_item.id
  }
  trigger {
    trigger_id = zabbix_trigger.demo_trigger.id
  }
}
```
### Template dependencies example:
```hcl
provider "zabbix" {
  user       = var.user
  password   = var.password
  server_url = "http://localhost/api_jsonrpc.php"
}

resource "zabbix_host_group" "demo_group" {
  name = "Template demo group"
}

resource "zabbix_template" "template_1" {
  host        = "template_1"
  groups      = [zabbix_host_group.demo_group.name]
}

resource "zabbix_template_link" "demo_template_1_link" {
  template_id = zabbix_template.template_1.id
}

resource "zabbix_template" "template_2" {
  host = "template_2"
  groups = [zabbix_host_group.demo_group.name]
  linked_template = [ # use the template link template_id value to be sure that all template_1 dependencies has been updated
    zabbix_template.demo_template_1_link.template_id
  ]
}

resource "zabbix_template_link" "demo_template_2_link" {
  template_id = zabbix_template.template_2.id
}
```