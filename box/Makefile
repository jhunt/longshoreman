#
# Makefile for longshoreman Vagrant image
#

RECIPE := build.json

default: test build

test:
	@cat $(RECIPE) | jq . >/dev/null
	@packer validate $(RECIPE)

build:
	packer build $(RECIPE)

clean:
	rm -rf output-*/
