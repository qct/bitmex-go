# Golang SDK for [bitmex](https://www.bitmex.com)

inspired by [https://github.com/BitMEX/api-connectors](https://github.com/BitMEX/api-connectors), [https://www.bitmex.com/api/explorer](https://www.bitmex.com/api/explorer) and [https://github.com/jxc6698/bitcoin-exchange-api](https://github.com/jxc6698/bitcoin-exchange-api)

## why
the generated connectors by bitmex have too many mistakes to use, this SDK fix these bugs to ensure bitmex API can get right results.

## how
all structs and APIs are from the bitmex official api connectors, based on that, add authentication and fix bugs.

## next
I only tested some API I will use, in the near future, I will keep adding and fixing to make more API available.

## installation
`go get github.com/qct/bitmex-go`

## how to use
what? need I tell you this? Look at `main.go`