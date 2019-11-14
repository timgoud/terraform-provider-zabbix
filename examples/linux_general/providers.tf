provider "zabbix" {
  user       = var.user
  password   = var.password
  server_url = "http://localhost/api_jsonrpc.php"
}


