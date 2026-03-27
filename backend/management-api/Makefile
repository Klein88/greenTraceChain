PLATFORM=$(shell uname -m)
DATETIME=$(shell date "+%Y%m%d%H%M%S")
VERSION=v3.0.0

chainmaker-management:
	@cd src && GOPATH=${GOPATH} go build -o ../bin/chainmaker-management.bin
	@mkdir -p ./release
	@rm -rf ./tmp/chainmaker-management/
	@mkdir -p ./tmp/chainmaker-management/
	@mkdir ./tmp/chainmaker-management/bin
	@mkdir ./tmp/chainmaker-management/log
	@mkdir ./tmp/chainmaker-management/configs
	@mkdir ./tmp/chainmaker-management/dependence
	@cp bin/chainmaker-management.bin ./tmp/chainmaker-management/bin
	@cp -r configs/* ./tmp/chainmaker-management/configs
	@cp -r dependence/* ./tmp/chainmaker-management/dependence
	@cp scripts/* ./tmp/chainmaker-management/bin
	@cd ./tmp;tar -zcvf chainmaker-management-$(VERSION).$(DATETIME).$(PLATFORM).tar.gz chainmaker-management; mv chainmaker-management-$(VERSION).$(DATETIME).$(PLATFORM).tar.gz ../release
	@rm -rf ./tmp/

build:
	@cd src go build -o ../bin/chainmaker-management.bin

run:
	@cd bin && ./chainmaker-management.bin -config ../configs/

clean:
	@rm -rf ./bin

docker:
	@docker build . -t management-backend:latest
	@docker tag management-backend:latest chainmakerofficial/management-backend:$(VERSION)

