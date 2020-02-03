variable "password" {
  type = string
}

variable "user" {
  type = string
}

variable "server_url" {
  type    = string
  default = "http://localhost/api_jsonrpc.php"
}
