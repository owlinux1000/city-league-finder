# City League Cancel Detector

This program aims to detect the cancellation of the current Pokemon Card City League.

## How to use

### 1. Install this tool

```
go install github.com/owlinux1000/city-league-cancel-detector
```

### 2. Create a config file

```yaml
endpoint: https://players.pokemon-card.com
prefecture:
  - Kanagawa
  - Tokyo
```

### 3. Run this tool with the config

```sh
$ export CONFIG_PATH=config.yaml city-league-cancel-detector
```
