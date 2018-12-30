CURRENT_TIMESTAMP=$(shell date +'%Y%m%d%H%M')
GOLANG_DEV_VERSION=0.1.5
IMAGE_REGISTRY_URL=docker.io
IMAGE_NAMESPACE=usvc
IMAGE_NAME=accounts

build: # builds the application - outputs an `app` binary
	@$(MAKE) _dev ARG="build"
build.docker: # builds the application into a docker image
	docker build \
		-t $(IMAGE_REGISTRY_URL)/$(IMAGE_NAMESPACE)/$(IMAGE_NAME):latest \
		.
test: build # runs tests in watch-mode
	@$(MAKE) _dev RUNARGS="-it" ARG="test"
test.once: build # runs tests once
	@$(MAKE) _go ARG="test -coverprofile c.out"
start: # starts the development environment
	@UID=$$(id -u) docker-compose up ${ARGS} app
start.once: build # runs the application on the host network
	@$(MAKE) _dev RUNARGS="--network host" ARG="start"
stop: # stops the development environment
	@UID=$$(id -u) docker-compose down ${ARGS}
migrate: # starts the migrator in the development enviornment
	@UID=$$(id -u) docker-compose run migrator
add.migration: # stubs a new migration
	@if [ "${NAME}" = "" ]; then \
		printf -- 'NAME variable was not specified.\n'; \
	else \
		$(MAKE) log.info MSG="Creating '$(CURRENT_TIMESTAMP)_$(NAME).up.sql'..."; \
		touch $(CURDIR)/migrations/$(CURRENT_TIMESTAMP)_$(NAME).up.sql; \
		$(MAKE) log.info MSG="Creating '$(CURRENT_TIMESTAMP)_$(NAME).down.sql'..."; \
		touch $(CURDIR)/migrations/$(CURRENT_TIMESTAMP)_$(NAME).down.sql; \
	fi
verify.migrations: # verifies that the migrations are up to date
	@if [ "$$(docker exec $$(docker ps | grep $$(basename $$(pwd)) | grep database | cut -f 1 -d ' ') bash -c "echo 'SELECT * FROM schema_migrations' | mysql -uroot -ptoor database 2>/dev/null | tail -n 1 | cut -f 1")" != "$$(ls -1 ./migrations | cut -f 1 -d '_' | sort -r | head -n 1)" ]; then \
		$(MAKE) log.error MSG="migrations are not up to date"; \
		exit 1; \
	else \
		$(MAKE) log.info MSG="migrations are in sync at $$(ls -1 ./migrations | cut -f 1 -d '_' | sort -r | head -n 1)"; \
	fi
shell: # creates a shell into the application container (requires application to be running)
	@docker exec -it $$(docker ps | grep $$(basename $$(pwd)) | grep app | cut -f 1 -d ' ') /bin/bash -l
db: # creates a shell into the database (requires database to be running)
	@docker exec -it $$(docker ps | grep $$(basename $$(pwd)) | grep database | cut -f 1 -d ' ') mysql -uroot -ptoor
logs: # displays the application logs
	@docker logs -f $$(docker ps | grep $$(basename $$(pwd)) | grep app | cut -f 1 -d ' ')
docs.add: # adds a document to the knowledge repository
	knowledge_repo --repo $(CURDIR) add ${DOC}
docs.update: # updates a document to the knowledge repository
	knowledge_repo --repo $(CURDIR) add --update ${DOC}
docs.view: # starts the knowledge repository server
	knowledge_repo --repo $(CURDIR) runserver
version.get: # retrieves the latest version we are at
	@docker run -v "$(CURDIR):/app" zephinzer/vtscripts:latest get-latest -q
version.bump: # bumps the version by 1: specify VERSION as "patch", "minor", or "major", to be specific about things
	@docker run -v "$(CURDIR):/app" zephinzer/vtscripts:latest iterate ${VERSION} -i
_dev: # base command to run (do not use)
	@docker run \
		${RUNARGS} \
		-u $$(id -u) \
		-v "$(CURDIR)/.cache/pkg:/go/pkg" \
		-v "$(CURDIR):/go/src/app" \
		zephinzer/golang-dev:$(GOLANG_DEV_VERSION) \
		${ARG}
_go: # base command to run (do not use)
	@docker run \
		-u $$(id -u) \
		-v "$(CURDIR)/.cache/pkg:/go/pkg" \
		-v "$(CURDIR):/go/src/app" \
		--entrypoint "go" \
		zephinzer/golang-dev:$(GOLANG_DEV_VERSION) \
		${ARG}
log.debug: # blue logs
	-@printf -- "\033[36m\033[1m_ [DEBUG] ${MSG}\033[0m\n"
log.info: # green logs
	-@printf -- "\033[32m\033[1m>  [INFO] ${MSG}\033[0m\n"
log.warn: # orange logs
	-@printf -- "\033[33m\033[1m?  [WARN] ${MSG}\033[0m\n"
log.error: # red logs
	-@printf -- "\033[31m\033[1m! [ERROR] ${MSG}\033[0m\n"