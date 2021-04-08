# chora-chain

**chora-chain** is an experimental blockchain built with [Cosmos SDK](https://github.com/cosmos/cosmos-sdk).

## Prerequisites

- [Go >= 1.15](https://golang.org/doc/install)
- [Starport >= 0.15](https://github.com/tendermint/starport)

## Quickstart

```
starport serve
```

## Testing IBC

```
starport serve -c config/chora-chora/starport/chora-1.yml
```

```
starport serve -c config/chora-chora/starport/chora-2.yml
```

```
starport relayer configure --advanced --source-rpc "http://0.0.0.0:26657" --source-faucet "http://0.0.0.0:4500" --source-port "transfer" --source-version "ics20-1" --target-rpc "http://0.0.0.0:26659" --target-faucet "http://0.0.0.0:4501" --target-port "transfer" --target-version "ics20-1"
```

```
starport relayer connect
```

## Testing Ecodex

```
starport serve -c config/chora-chora/starport/chora-1.yml
```

```
starport serve -c config/chora-chora/starport/chora-2.yml
```

```
starport relayer configure --advanced --source-rpc "http://0.0.0.0:26657" --source-faucet "http://0.0.0.0:4500" --source-port "ecodex" --source-version "ecodex-1" --target-rpc "http://0.0.0.0:26659" --target-faucet "http://0.0.0.0:4501" --target-port "ecodex" --target-version "ecodex-1"
```

```
starport relayer connect
```

## Learn More

- [Chora Documentation](https://docs.chora.io)
- [Cosmos SDK Documentation](https://docs.cosmos.network)
- [Cosmos SDK Tutorials](https://tutorials.cosmos.network)
