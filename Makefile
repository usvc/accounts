init:
	@$(MAKE) dev ARG="init"
build:
	@$(MAKE) dev ARG="build"
test:
	@$(MAKE) dev ARG="test"
test.once: build
	@go test -coverprofile c.out
start:
	@UID=$$(id -u) docker-compose up -V migrator app
run: build
	$(CURDIR)/app
shell:
	@$(MAKE) dev ARG="shell"
dev:
	@docker run -it --network host -u $$(id -u) -v "$(CURDIR):/go/src/app" zephinzer/golang-dev:latest ${ARG}
version.get:
	@docker run -v "$(CURDIR):/app" zephinzer/vtscripts:latest get-latest -q
version.bump:
	@docker run -v "$(CURDIR):/app" zephinzer/vtscripts:latest iterate ${VERSION} -i