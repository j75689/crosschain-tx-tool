# crosschain-tx-tool

A testing tool for sending cross chain tx between greenfield and BSC

### Build
```shell script
make build
````


### Run

gnfd -> bsc replace with your values
```shell script
./build/crossd --from-chain gnfd --key 31d8f71a3f631bb90fc980a7bb631a7314b8bf9e2c78f3833f7afb9f1d91884d --chain-id greenfield_9000-1741 --url k8s-gnfd-gnfdvali-05c593a3c7-c5502bc443c42322.elb.ap-northeast-1.amazonaws.com:9090 --to 0x76d244CE05c3De4BbC6fDd7F56379B145709ade9 --amount 100

````
bsc -> gnfd replace with your values
```shell script
./build/crossd --from-chain bsc --key 31d8f71a3f631bb90fc980a7bb631a7314b8bf9e2c78f3833f7afb9f1d91884d --chain-id 714 --url https://qa-gnfd-bsc.bnbchain.world --to 0x76d244CE05c3De4BbC6fDd7F56379B145709ade9 --amount 100000000000000000 --contract-addr 0xc8CB5439c767A63aca1c01862252B2F3495fDcFE 
````