REPO_HOME=${GOPATH}/src/github.com/OCRVblockchain/fabric-certstore
SDK_HOME=${GOPATH}/src/github.com/hyperledger/fabric-sdk-go
SDK_TAG=v1.0.0

fetch_sdk:
	@git clone --single-branch --branch $(SDK_TAG) https://github.com/hyperledger/fabric-sdk-go $(SDK_HOME)

patch_sdk:
	@cd $(SDK_HOME) && git apply $(REPO_HOME)/sdk.patch

build_image:
	@docker build -t patched-fabric-peer .

#run_test:
#	@docker run -t \
#		-v $(GOPATH):/go \
#		-v $(HOME)/fabric-samples/first-network/crypto-config:/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto \
#		-w /go/src/$(TRAVIS_GO_IMPORT_PATH)/test \
#		--network=net_byfn \
#		golang:latest bash -c \
#		'go mod download 2>&1 | awk "!/^go: (finding|downloading|extracting)/" && go test -v -failfast ./main_test.go'

clean:
	@rm -rf $(SDK_HOME)
	@docker rmi -f patched-fabric-peer
