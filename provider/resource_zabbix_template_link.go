package provider

import (
	"fmt"
	"strings"

	"github.com/claranet/go-zabbix-api"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceZabbixTemplateLink() *schema.Resource {
	return &schema.Resource{
		Create: resourceZabbixTemplateLinkCreate,
		Read:   resourceZabbixTemplateLinkRead,
		Update: resourceZabbixTemplateLinkUpdate,
		Delete: resourceZabbixTemplateLinkDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				templateID, itemsID, triggersID, err := resourceZabbixTemplateLinkParseID(d.Id())
				if err != nil {
					return nil, err
				}
				d.Set("item", itemsID)
				d.Set("trigger", triggersID)
				d.Set("template_id", templateID)
				d.SetId(randStringNumber(5))
				return []*schema.ResourceData{d}, nil
			}},
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"server_item": &schema.Schema{ // Use to detect change in server item this shouldn't be used by the user
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"trigger": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Required: true,
			},
			"server_trigger": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func resourceZabbixTemplateLinkCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(randStringNumber(5))
	return resourceZabbixTemplateLinkReadLocal(d, meta)
}

func resourceZabbixTemplateLinkRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*zabbix.API)

	itemsTerraform, err := getTerraformTemplateItems(d, api)
	if err != nil {
		return err
	}
	localItems := d.Get("item").(*schema.Set).List()

	var serverItems []string
	for _, item := range itemsTerraform {
		present := false
		for _, localItem := range localItems {
			if item == localItem.(string) {
				present = true
				break
			}
		}

		if !present {
			serverItems = append(serverItems, item)
		}
	}
	d.Set("server_item", serverItems)

	triggersTerraform, err := getTerraformTemplateTriggers(d, api)
	if err != nil {
		return err
	}
	localTriggers := d.Get("trigger").(*schema.Set).List()

	var serverTriggers []string
	for _, trigger := range triggersTerraform {
		present := false
		for _, localTrigger := range localTriggers {
			if trigger == localTrigger.(string) {
				present = true
				break
			}
		}

		if !present {
			serverTriggers = append(serverTriggers, trigger)
		}
	}
	d.Set("server_trigger", serverTriggers)

	return nil
}

func resourceZabbixTemplateLinkReadLocal(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*zabbix.API)

	itemsTerraform, err := getTerraformTemplateItems(d, api)
	if err != nil {
		return err
	}
	d.Set("item", itemsTerraform)
	d.Set("server_item", []string{})

	triggersTerraform, err := getTerraformTemplateTriggers(d, api)
	if err != nil {
		return err
	}
	d.Set("trigger", triggersTerraform)
	d.Set("server_trigger", []string{})

	return nil
}

func resourceZabbixTemplateLinkUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*zabbix.API)

	if err := updateZabbixTemplateItem(d, api); err != nil {
		return err
	}
	if err := updateZabbixTemplateTrigger(d, api); err != nil {
		return err
	}
	return resourceZabbixTemplateLinkReadLocal(d, meta)
}

func resourceZabbixTemplateLinkDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func getTerraformTemplateItems(d *schema.ResourceData, api *zabbix.API) ([]string, error) {
	params := zabbix.Params{
		"output": "extend",
		"templateids": []string{
			d.Get("template_id").(string),
		},
		"inherited": false,
	}
	items, err := api.ItemsGet(params)
	if err != nil {
		return nil, err
	}

	itemsTerraform := make([]string, len(items))
	for i, item := range items {
		itemsTerraform[i] = item.ItemID
	}
	return itemsTerraform, nil
}

func getTerraformTemplateTriggers(d *schema.ResourceData, api *zabbix.API) ([]string, error) {
	params := zabbix.Params{
		"output": "extend",
		"templateids": []string{
			d.Get("template_id").(string),
		},
		"inherited": false,
	}
	triggers, err := api.TriggersGet(params)
	if err != nil {
		return nil, err
	}

	TriggersTerraform := make([]string, len(triggers))
	for i, trigger := range triggers {
		TriggersTerraform[i] = trigger.TriggerID
	}
	return TriggersTerraform, nil
}

func updateZabbixTemplateItem(d *schema.ResourceData, api *zabbix.API) error {
	if d.HasChange("server_item") {
		oldI, _ := d.GetChange("server_item")
		localItemsTerraform := oldI.(*schema.Set).List()
		var deletedItems []string

		for _, item := range localItemsTerraform {
			deletedItems = append(deletedItems, item.(string))
		}
		if err := api.ItemsDeleteByIds(deletedItems); err != nil {
			return err
		}
	}
	return nil
}

func updateZabbixTemplateTrigger(d *schema.ResourceData, api *zabbix.API) error {
	if d.HasChange("server_trigger") {
		oldT, _ := d.GetChange("server_trigger")
		localTriggersTerraform := oldT.(*schema.Set).List()
		var deletedTriggers []string

		for _, trigger := range localTriggersTerraform {
			deletedTriggers = append(deletedTriggers, trigger.(string))
		}
		if err := api.TriggersDeleteByIds(deletedTriggers); err != nil {
			return err
		}
	}
	return nil
}

func resourceZabbixTemplateLinkParseID(ID string) (templateID string, itemID []string, triggerID []string, err error) {
	parseID := strings.Split(ID, "_")
	if len(parseID) != 3 {
		err = fmt.Errorf(`Expected id format TEMPLATEID_ITEMID_TRIGGERID,
		if you have multiple ITEMID and TRIGGERID use "." to separate the id`)
		return
	}
	templateID = parseID[0]
	itemID = strings.Split(parseID[1], ".")
	triggerID = strings.Split(parseID[2], ".")
	return
}
