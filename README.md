# Pikamon
A Pokemon Discord bot inspired by the late Pokecord.

[![CircleCI](https://circleci.com/gh/Jac0bDeal/pikamon.svg?style=shield)](https://circleci.com/gh/Jac0bDeal/pikamon)
[![Go Report Card](https://goreportcard.com/badge/github.com/Jac0bDeal/pikamon)](https://goreportcard.com/report/github.com/Jac0bDeal/pikamon)

## Installation
Pikamon is built on Go 1.14+, please make sure this version is used to build and test the application.

Build the Pikamon binary using the make target `build` (`rebuild` if you already have built it before):
```shell script
make build
make rebuild
```

Once it is built, then the bot is can be run at minimum with the `--token` or `-t` flag:
```shell script
./bin/pikamon --token <TOKEN_HERE>
./bin/pikamon -t <TOKEN_HERE>
```
Alternatively, if the token is exported via the appropriate env variable, the `--token` flag is not needed:
```shell script
export PIKAMON_TOKEN <TOKEN_HERE>
./bin/pikamon
```
