package zabbix

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccZabbixDataSourceServer_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccZabbixDataSourceServerConfig_basic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.zabbix_server.test", "server_version"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_days", testZabbixServerUnitDays()),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_hours", testZabbixServerUnitHours()),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_minutes", testZabbixServerUnitMinutes()),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_seconds", testZabbixServerUnitSeconds()),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_weeks", testZabbixServerUnitWeeks()),
				),
			},
		},
	})
}

func TestAccZabbixDataSourceServer_force_32(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccZabbixDataSourceServerConfig_force_version("3.2.0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.zabbix_server.test", "server_version", "3.2.0"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_days", ""),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_hours", ""),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_minutes", ""),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_seconds", ""),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_weeks", ""),
				),
			},
			{
				Config: testAccZabbixDataSourceServerConfig_force_version("3.2.11"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.zabbix_server.test", "server_version", "3.2.11"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_days", ""),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_hours", ""),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_minutes", ""),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_seconds", ""),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_weeks", ""),
				),
			},
		},
	})
}

func TestAccZabbixDataSourceServer_force_34(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccZabbixDataSourceServerConfig_force_version("3.4.0"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.zabbix_server.test", "server_version", "3.4.0"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_days", "d"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_hours", "h"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_minutes", "m"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_seconds", "s"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_weeks", "w"),
				),
			},
			{
				Config: testAccZabbixDataSourceServerConfig_force_version("3.4.15"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.zabbix_server.test", "server_version", "3.4.15"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_days", "d"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_hours", "h"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_minutes", "m"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_seconds", "s"),
					resource.TestCheckResourceAttr("data.zabbix_server.test", "unit_time_weeks", "w"),
				),
			},
		},
	})
}

func testAccZabbixDataSourceServerConfig_basic() string {
	return `
data "zabbix_server" "test" {
}
`
}

func testAccZabbixDataSourceServerConfig_force_version(version string) string {
	return fmt.Sprintf(`
data "zabbix_server" "test" {
	server_version = "%s"
}
`, version)
}
