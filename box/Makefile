#
# Makefile for longshoreman Vagrant image
#

RECIPE := build.json
BOX    := jhunt/longshoreman

default: test

test:
	@cat $(RECIPE) | jq . >/dev/null
	@packer validate $(RECIPE)

push:
	packer push -name $(BOX) $(RECIPE)

build:
	packer build $(RECIPE)

clean:
	rm -rf output-*/

import:
	vagrant box remove --force longshoreman || true
	vagrant box add --name longshoreman boxen/longshoreman.box
