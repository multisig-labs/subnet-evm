cast call --rpc-url http://127.0.0.1:9652/ext/bc/2EXjigcFHF4wVASaghLLJ9pzBocfstoPUsR7nAo5zpfiVt3BLE/rpc 0x0300000000000000000000000000000000000000 `cast calldata "getCurrentBlockNumber()"`

cast call --rpc-url http://127.0.0.1:9652/ext/bc/2EXjigcFHF4wVASaghLLJ9pzBocfstoPUsR7nAo5zpfiVt3BLE/rpc 0x5DB9A7629912EBF95876228C24A848de0bfB43A9 `cast calldata "getCurrentBlockNumber()"`

252dba42
cast sig "aggregate((address,bytes)[])"
[["0x0300000000000000000000000000000000000000","0x11"]]

cast call 0x0300000000000000000000000000000000000000 0x252dba420000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000200000000000000000000000000300000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000011100000000000000000000000000000000000000000000000000000000000000

https://abi.hashex.org/

// NewDefaultNetwork returns a new network using a pre-defined
// network configuration.
// The following addresses are pre-funded:
// X-Chain Address 1: X-custom18jma8ppw3nhx5r4ap8clazz0dps7rv5u9xde7p
// X-Chain Address 1 Key: PrivateKey-ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN
// X-Chain Address 2: X-custom16045mxr3s2cjycqe2xfluk304xv3ezhkhsvkpr
// X-Chain Address 2 Key: PrivateKey-2fzYBh3bbWemKxQmMfX6DSuL2BFmDSLQWTvma57xwjQjtf8gFq
// P-Chain Address 1: P-custom18jma8ppw3nhx5r4ap8clazz0dps7rv5u9xde7p
// P-Chain Address 1 Key: PrivateKey-ewoqjP7PxY4yr3iLTpLisriqt94hdyDFNgchSxGGztUrTXtNN
// P-Chain Address 2: P-custom16045mxr3s2cjycqe2xfluk304xv3ezhkhsvkpr
// P-Chain Address 2 Key: PrivateKey-2fzYBh3bbWemKxQmMfX6DSuL2BFmDSLQWTvma57xwjQjtf8gFq
// C-Chain Address: 0x8db97C7cEcE249c2b98bDC0226Cc4C2A57BF52FC
// C-Chain Address Key: 56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027
// The following nodes are validators:
// _ NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg
// _ NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ
// _ NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN
// _ NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu
// \* NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5
