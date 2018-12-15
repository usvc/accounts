GOLANG_DEV_VERSION=0.1.5

init: # initialises this directory - use once only
	@$(MAKE) _dev ARG="init"
build: # builds the application - outputs an `app` binary
	@$(MAKE) _dev ARG="build"
test: build # runs tests in watch-mode
	@$(MAKE) _dev ARG="test"
test.once: build # runs tests once
	@$(MAKE) _dev ARG="test -coverprofile c.out"
start: # starts the development environment
	@UID=$$(id -u) docker-compose up ${ARGS} app
start.once: build # runs the application on the host network
	@$(MAKE) _dev ARG="start"
stop: # stops the development environment
	@UID=$$(id -u) docker-compose down ${ARGS}
migrate: # starts the migrator in the development enviornment
	@UID=$$(id -u) docker-compose run migrator
shell: # creates a shell into the application container (requires application to be running)
	@docker exec -it $$(docker ps | grep $$(basename $$(pwd)) | grep app | cut -f 1 -d ' ') /bin/bash -l
db: # creates a shell into the database (requires database to be running)
	@docker exec -it $$(docker ps | grep $$(basename $$(pwd)) | grep database | cut -f 1 -d ' ') mysql -uroot -ptoor
logs: # displays the application logs
	@docker logs -f $$(docker ps | grep $$(basename $$(pwd)) | grep app | cut -f 1 -d ' ')
version.get: # retrieves the latest version we are at
	@docker run -v "$(CURDIR):/app" zephinzer/vtscripts:latest get-latest -q
version.bump: # bumps the version by 1: specify VERSION as "patch", "minor", or "major", to be specific about things
	@docker run -v "$(CURDIR):/app" zephinzer/vtscripts:latest iterate ${VERSION} -i
_dev: # base command to run (do not use)
	@docker run -it --network host -u $$(id -u) -v "$(CURDIR)/.cache/pkg:/go/pkg" -v "$(CURDIR):/go/src/app" zephinzer/golang-dev:$(GOLANG_DEV_VERSION) ${ARG}