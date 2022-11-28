# Poker Probability Calculator

[![Actions Status](https://github.com/abhayptp/poker-probability/actions/workflows/go.yml/badge.svg)](https://github.com/abhayptp/poker-probability/actions)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A Golang CLI app which calculates probability of winning poker round based on known information ie open community cards, self cards and other player cards (if you know them :p).

## Build
```
go build
```

## Run
```
./poker-probability -community_cards="<community-cards>" -simulation_rounds=<simulation-rounds> <space-separated-player-cards>
```
Example:
```
./poker-probability -community_cards="AS" -simulation_rounds=10000 -player_count=5 "AH,1D"
```