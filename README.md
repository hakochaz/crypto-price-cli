# Crypto Price CLI
Command line interface for getting reliable price data - using the CoinGecko API.

# Installation 
Install the program via the following command.

```
go get github.com/hakochaz/crypto-price-cli
```

## Price Command
Get current or historical price data for any crypocurrency. 

```
crypto-price-cli price --id=bitcoin --vc=usd --date=05-19-2018 --amount=2.57
```

### Required flags 

```id``` - identifer of the coin you wish to get price data for.

```vc``` - versus currency that the price will be denominated in.

### Optional flags

```date``` - get historical price data for a specific date in the past.

```amount``` - calculate value of a specific amount of crypto.

## Search Command 
Search for a supported coin by name to get the id for use in price commands.

```
crypto-price-cli search --name=Bitcoin
```

## VC Command
List all the supported versus currencies.

```
crypto-price-cli vc
```

## Coins Command
Get a list of all compatible coins.

There is no pagination, so it is recommended to use the search command to get the id of a coin.

```
crypto-price-cli coins
```
