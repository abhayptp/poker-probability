# Poker Probability

[![Actions Status](https://github.com/abhayptp/poker-probability/actions/workflows/go.yml/badge.svg)](https://github.com/abhayptp/poker-probability/actions)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A Golang CLI app which calculates probability of winning poker round based on known information ie open community cards, self cards and other player !
cards (if you know them :p).

## Build
```
go build
```

## Run
```
./poker-probability -community_cards="<community-cards>" -simulation_rounds=<simulation-rounds> -players_count=<players-count> <space-separated-player-cards>
```

Example:
```
./poker-probability -community_cards="AS" -simulation_rounds=1000 -players_count=5 "AH,1D"
```
![2022-11-29-014041_2560x1440_scrot](https://user-images.githubusercontent.com/22256898/204371908-f4c3a664-b210-4ca3-9db9-b75297eba7e3.png)
