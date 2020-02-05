package zabbix

import (
	"fmt"
	"log"
	"strings"

	"github.com/claranet/go-zabbix-api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mcuadros/go-version"
)

func dataSourceZabbixServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceZabbixServerRead,
		Schema: map[string]*schema.Schema{
			"server_version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Version of the Zabbix server.",
			},
			"unit_time_days": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String representation of the days unit used by the Zabbix server (`d` for version 3.4+, empty string otherwise).",
			},
			"unit_time_hours": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String representation of the hours unit used by the Zabbix server (`h` for version 3.4+, empty string otherwise).",
			},
			"unit_time_minutes": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String representation of the minutes unit used by the Zabbix server (`m` for version 3.4+, empty string otherwise).",
			},
			"unit_time_seconds": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String representation of the seconds unit used by the Zabbix server (`s` for version 3.4+, empty string otherwise).",
			},
			"unit_time_weeks": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "String representation of the weeks unit used by the Zabbix server (`w` for version 3.4+, empty string otherwise).",
			},
		},
	}
}

func dataSourceZabbixServerRead(d *schema.ResourceData, meta interface{}) (err error) {
	var serverVersion string
	if v, ok := d.GetOkExists("server_version"); ok {
		serverVersion = v.(string)
		log.Printf("[DEBUG] Forcing Zabbix Server version to %s\n", serverVersion)
	} else {
		api := meta.(*zabbix.API)

		serverVersion, err = api.Version()
		if err != nil {
			return err
		}

		log.Printf("[DEBUG] Actual Zabbix Server version is %s\n", serverVersion)
	}

	d.SetId(fmt.Sprintf("zabbix_server_%s", strings.ReplaceAll(serverVersion, ".", "_")))
	d.Set("server_version", serverVersion)

	if version.Compare(serverVersion, "3.4.0", ">=") {
		d.Set("unit_time_days", "d")
		d.Set("unit_time_hours", "h")
		d.Set("unit_time_minutes", "m")
		d.Set("unit_time_seconds", "s")
		d.Set("unit_time_weeks", "w")
	} else {
		d.Set("unit_time_days", "")
		d.Set("unit_time_hours", "")
		d.Set("unit_time_minutes", "")
		d.Set("unit_time_seconds", "")
		d.Set("unit_time_weeks", "")
	}

	return nil
}
