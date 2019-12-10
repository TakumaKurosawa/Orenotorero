# docker #
DOCKER = docker
DOCKER_BUILD = $(DOCKER) build -t client .
DOCKER_RUN = $(DOCKER) run -d -p 80:3000 --name orenotorero-client client
DOCKER_LOGS = $(DOCKER) logs orenotorero-client
DOCKER_STOP = $(DOCKER) stop orenotorero-client
DOCKER_CONTAINER_PRUNE = $(DOCKER) container prune -f


help: ## 使い方
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

start: ## DBコンテナを立ち上げるコマンド
	$(DOCKER_BUILD)
	$(DOCKER_RUN)

stop: ## 立ち上がっているDBコンテナをクローズするコマンド
	$(DOCKER_STOP)
	$(DOCKER_CONTAINER_PRUNE)

log: ## 立ち上げているDBコンテナのログを表示するコマンド
	$(DOCKER_LOGS)
