GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

FinalityRelayerManagerAbiPath := ./manta-fp-contracts/out/FinalityRelayerManager.sol/FinalityRelayerManager.json
BLSApkRegistryAbiPath := ./manta-fp-contracts/out/BLSApkRegistry.sol/BLSApkRegistry.json
SymbioticVaultAbiPath := ./bindings/sfp/abi.json


build:
	env GO111MODULE=on go build -o manta-fp-aggregator ./cmd

clean:
	rm manta-fp-aggregator

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

binding-sv:
	$(eval temp := $(shell mktemp))

	cat $(SymbioticVaultAbiPath) \
    	| jq -r .bytecode.object > $(temp)

	cat $(SymbioticVaultAbiPath) \
		| jq .abi \
		| abigen --pkg sfp \
		--abi - \
		--out bindings/sfp/symbiotic_vault.go \
		--type SymbioticVault \
		--bin $(temp)

		rm $(temp)

.PHONY: \
	 build \
	 compile \
	 bindings \
	 binding-bls \
	 binding-finality \
	 binding-sv \
	 clean \
	 test \
	 lint \