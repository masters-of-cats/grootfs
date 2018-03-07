.PHONY: all \
	test \
	go-vet concourse-go-vet go-generate \
	image push-image \
	update-deps unit integration

all:
	GOOS=linux go build -o grootfs .
	GOOS=linux go build -o tardis ./store/filesystems/overlayxfs/tardis

cf: all
	GOOS=linux go build -tags cloudfoundry -o tardis ./store/filesystems/overlayxfs/tardis

###### Help ###################################################################

help:
	@echo '    all ................................. builds the grootfs cli'
	@echo '    deps ................................ installs dependencies'
	@echo '    update-deps ......................... updates dependencies'
	@echo '    unit ................................ run unit tests'
	@echo '    integration ......................... run integration tests'
	@echo '    test ................................ runs tests in concourse-lite'
	@echo '    compile-tests ....................... checks that tests can be compiled'
	@echo '    go-vet .............................. runs go vet in grootfs source code'
	@echo '    concourse-go-vet .................... runs go vet in concourse-lite'
	@echo '    go-generate ......................... runs go generate in grootfs source code'
	@echo '    image ............................... builds a docker image'
	@echo '    push-image .......................... pushes image to docker-hub'

###### Dependencies ###########################################################

deps:
	git submodule update --init --recursive

###### Testing ################################################################

compile-tests:
	ginkgo build -r .; find . -name '*.test' | xargs rm

unit:
	./script/test -u

integration:
	./script/test -i

test: concourse-go-vet
	./script/test -a

###### Go tools ###############################################################

go-vet:
	GOOS=linux go vet `go list ./... | grep -v vendor`

concourse-go-vet:
	fly -t garden-ci e -c ci/tasks/go-vet.yml -i grootfs-git-repo=${PWD}

go-generate:
	GOOS=linux go generate `go list ./... | grep -v vendor`

###### Docker #################################################################

image:
	docker build -t cfgarden/grootfs-ci .

push-image:
	docker push cfgarden/grootfs-ci
