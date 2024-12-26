GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

FinalityRelayerManagerAbiPath := ./l2fp-contracts/out/FinalityRelayerManager.sol/FinalityRelayerManager.json
BLSApkRegistryAbiPath := ./l2fp-contracts/out/BLSApkRegistry.sol/BLSApkRegistry.json


build:
	env GO111MODULE=on go build -o l2fp-aggregator ./cmd

clean:
	rm l2fp-aggregator

test:
	go test -v ./...

lint:
	golangci-lint run ./...

compile:
	cd ./l2fp-contracts && forge install && forge build && cd ../

bindings: binding-bls binding-finality

binding-bls:
	$(eval temp := $(shell mktemp))

	cat $(BLSApkRegistryAbiPath) \
    	| jq -r .bytecode.object > $(temp)

	cat $(BLSApkRegistryAbiPath) \
		| jq .abi \
		| abigen --pkg bls \
		--abi - \
		--out bindings/bls/bls_apk_registry.go \
		--type BLSApkRegistry \
		--bin $(temp)

		rm $(temp)

binding-finality:
	$(eval temp := $(shell mktemp))

	cat $(FinalityRelayerManagerAbiPath) \
    	| jq -r .bytecode.object > $(temp)

	cat $(FinalityRelayerManagerAbiPath) \
		| jq .abi \
		| abigen --pkg finality \
		--abi - \
		--out bindings/finality/finality_relayer_manager.go \
		--type FinalityRelayerManager \
		--bin $(temp)

		rm $(temp)


.PHONY: \
	 build \
	 compile \
	 bindings \
	 binding-bls \
	 binding-finality \
	 clean \
	 test \
	 lint \