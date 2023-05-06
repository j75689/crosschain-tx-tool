# crosschain-tx-tool

A testing tool for sending cross chain tx between greenfield and BSC

### Build
```shell script
make build
````


### Run

gnfd -> bsc replace with your values
```shell script
./build/crossd --from-chain gnfd --key 31d8f71a3f631bb90fc980a7bb631a7314b8bf9e2c78f3833f7afb9f1d91884d --chain-id greenfield_9000-121 --url http://127.0.0.1:26750 --to 0x76d244CE05c3De4BbC6fDd7F56379B145709ade9 --amount 100 --tx-count 1

````
bsc -> gnfd replace with your values
```shell script
./build/crossd --from-chain bsc --key 31d8f71a3f631bb90fc980a7bb631a7314b8bf9e2c78f3833f7afb9f1d91884d --chain-id 714 --url http://localhost:8502 --to 0x76d244CE05c3De4BbC6fDd7F56379B145709ade9 --amount 100000000000000000 --contract-addr 0x35dd738E306d32f2709824B6f744f188DA01D3C5 --tx-count 1
````