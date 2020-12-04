REGISTRY ?= "" 

## Go compiles elector for linux os
compile:
	GOOS=linux GOARCH=amd64 go build -o elector main.go

.PHONY: docker
## Docker build and push to your registry, needs REGISTRY env var to be set
docker: checkarg compile
	docker build . -t elector-test
	docker tag elector-test:latest $(REGISTRY)/elector-test
	docker push $(REGISTRY)/elector-test

.PHONY: check
## Check whichs pod is elected leader
check:
	@echo "Running services:"
	@kubectl get pod | grep elector | cut -f1 -d' '
	@echo "\nElected leader:"
	@kubectl get lease elector-lock -n default -o yaml | grep holderIdentity | cut -f2 -d':' | cut -c2-

.PHONY: deploy
## Deploy serviceaccount and deployment to k8s, requires REGISTRY env var to be set
deploy: checkarg
	@cp deployment.yaml dep.yaml
	@sed -i.bak "s#__REGISTRY__#$(REGISTRY)#" dep.yaml
	@kubectl apply -f dep.yaml
	@rm -rf dep.yaml dep.yaml.bak

## Deploys service account with access to coordination api, only needed once 
deploy-sa:
	@kubectl apply -f sa.yaml

.PHONY: checkarg
checkarg:
	@if [[ $(REGISTRY) = "" ]]; then \
		echo "Please set REGISTRY env var to URL to your registry"; exit 1; \
	fi

# COLORS
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
RESET  := $(shell tput -Txterm sgr0)

.PHONY: help
# Help target stolen from this comment: https://gist.github.com/prwhite/8168133#gistcomment-2278355
## Show help
help:
	@echo 'Usage:'
	@echo '  ${YELLOW}make${RESET} ${GREEN}<target>${RESET}'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${YELLOW}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)
