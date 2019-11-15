TEST?=$$(go list ./... |grep -v 'vendor')
PKG_NAME=zabbix
DIR=~/.terraform.d/plugins

default: build


build:
	go install

install:
	mkdir -vp $(DIR)
	go build -o $(DIR)/terraform-provider-datadog

uninstall:
	@rm -vf $(DIR)/terraform-provider-datadog


test:
	go test -i $(TEST) || exit 1
	echo $(TEST) | \
		xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m

