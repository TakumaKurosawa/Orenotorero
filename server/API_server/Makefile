# docker #
DOCKER = docker
DOCKER_BUILD = $(DOCKER) build -t api .
DOCKER_RUN = $(DOCKER) run -d -p 3000:3000 --name orenotorero-api api
DOCKER_LOGS = $(DOCKER) logs orenotorero-api
DOCKER_STOP = $(DOCKER) stop orenotorero-api
DOCKER_CONTAINER_PRUNE = $(DOCKER) container prune -f


help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

start: ## APIを立ち上げるコマンド
	$(DOCKER_BUILD)
	$(DOCKER_RUN)

stop: ## 立ち上がっているAPIをクローズするコマンド
	$(DOCKER_STOP)
	$(DOCKER_CONTAINER_PRUNE)

log: ## 立ち上げているAPIコンテナのログを表示するコマンド
	$(DOCKER_LOGS)

