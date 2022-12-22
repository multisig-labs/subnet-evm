# Autoload a .env if one exists
set dotenv-load

export SUBNET_EVM_VERSION := "v0.4.6-ggp1"

build:
	mkdir -p ./build
	./scripts/build.sh build/ggp_evm.{{SUBNET_EVM_VERSION}}.bin

rebuild: build
	find ~/.avalanche-cli -name "nYKSQqw15js3YmQxX1Z8aC13jr3EA39V9x1DQpep9yKz9TYPc" -depth -exec rm {} \;
	cp -f build/ggp_evm.{{SUBNET_EVM_VERSION}}.bin ~/.avalanche-cli/bin/avalanchego/avalanchego-v1.9.5/plugins/nYKSQqw15js3YmQxX1Z8aC13jr3EA39V9x1DQpep9yKz9TYPc

create: build
	avalanche subnet create ggpevm --force --custom --vm  build/ggp_evm.{{SUBNET_EVM_VERSION}}.bin --genesis ./genesis.json
	jq '.TokenName="GGP"' ~/.avalanche-cli/subnets/ggpevm/sidecar.json > ~/.avalanche-cli/subnets/ggpevm/sidecar-fixed.json
	mv ~/.avalanche-cli/subnets/ggpevm/sidecar-fixed.json ~/.avalanche-cli/subnets/ggpevm/sidecar.json
	cat ~/.avalanche-cli/subnets/ggpevm/*

deploy:
  avalanche subnet deploy ggpevm --local

start:
	avalanche network start

stop:
	avalanche network stop 

restart: stop rebuild start

run-script file:
	cd contract-examples && npx hardhat run --network local scripts/{{file}}

test file="./test/ExampleMulticall.ts":
  cd contract-examples && npx hardhat test --network local {{file}}

test-e2e:
	E2E=true ./scripts/run.sh

ping:
  cast call 0x0300000000000000000000000000000000000000 `cast sig "getCurrentBlockNumber()"`
  # cast call --rpc-url http://127.0.0.1:9654/ext/bc/jtd1894n4pQ2Ph3bdR22vWcEiBvbvWf6si67N9GDUr6hmHBTt/rpc 0x0200000000000000000000000000000000000001 `cast calldata "mintNativeCoin(address,uint256)"`


env:
	env | sort

clean:
	find ~/.avalanche-cli -name ".DS_Store" -depth -exec rm {} \;
	find ~/.avalanche-cli -name "nYKSQqw15js3YmQxX1Z8aC13jr3EA39V9x1DQpep9yKz9TYPc" -depth -exec rm {} \;
	# rm -f ~/.avalanche-cli/bin/avalanchego/avalanchego-v1.9.4/plugins/nYKSQqw15js3YmQxX1Z8aC13jr3EA39V9x1DQpep9yKz9TYPc
	rm -rf ~/.avalanche-cli/runs
	rm -rf ~/.avalanche-cli/logs
	rm -rf ~/.avalanche-cli/snapshots
	rm -rf ~/.avalanche-cli/subnets
	rm -rf ~/.avalanche-cli/vms

	

# status:
# 	subnet-cli status blockchain \
# 		--private-uri=http://localhost:9650 \
# 		--blockchain-id=arRzKZKieSmwPGPNnGHb6kMEKyLKzjYrUCeF2gcikHyidRouX \
# 		--check-bootstrapped
# VMID nZ7XwSmMoCqmy9icAuerHW9Apv1UK6f7kfW6x1et9q7k7DUFd from gogopool
# $AVALANCHEGO_PLUGIN_PATH/nZ7XwSmMoCqmy9icAuerHW9Apv1UK6f7kfW6x1et9q7k7DUFd
# cp $AVALANCHEGO_PLUGIN_PATH/nZ7XwSmMoCqmy9icAuerHW9Apv1UK6f7kfW6x1et9q7k7DUFd ~/.avalanche-cli/vms/gogopool

# x: 
# 	avalanche subnet create testnet --genesis ./genesis.json --evm --latest
# add-validator:
# 	subnet-cli add subnet-validator --enable-prompt=false \
# 		--public-uri=http://localhost:9650 \
# 		--subnet-id=p433wpuXyJiDhyazPYyZMJeaoPSW76CBZ2x7wrVPLgvokotXz \
# 		--node-ids=NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5

# create-blockchain:
# 	subnet-cli create blockchain --enable-prompt=false  \
# 		--public-uri=http://localhost:9650 \
# 		--subnet-id=p433wpuXyJiDhyazPYyZMJeaoPSW76CBZ2x7wrVPLgvokotXz \
# 		--chain-name=subnetevm \
# 		--vm-id=srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy \
# 		--vm-genesis-path="/tmp/genesis.json"



# ava-start:
# 	$AVALANCHEGO_EXEC_PATH --network-id=local --staking-enabled=false --log-level=debug
	
# anr-boot: 
# 	avalanche-network-runner server \
# 		--log-level debug \
# 		--port=":8080" \
# 		--grpc-gateway-port=":8081"

# anr-start:
# 	avalanche-network-runner control start \
# 		--log-level debug \
# 		--endpoint="0.0.0.0:8080" \
# 		--number-of-nodes=5 \
# 		--avalanchego-path $AVALANCHEGO_EXEC_PATH \
# 		--plugin-dir $AVALANCHEGO_PLUGIN_PATH

# # --blockchain-specs '[{"vm_name": "subnetevm", "genesis": "/tmp/genesis.json"}]'		

# start:
# 	# subnet-cli status blockchain --blockchain-id=2CA6j5zYzasynPsFeNoqWkmTCt3VScMvXUZHbfDJ8k3oGzAPtU  --private-uri=http://localhost:9650 
# 	# srEXiWaHuhNyGwPUi444Tu47ZEDwxTWrbQiuD7FmgSAQ6X7Dy
# 	# subnet-cli create VMID subnetevm
# 	subnet-cli create subnet --public-uri=http://localhost:9650 --enable-prompt=false
