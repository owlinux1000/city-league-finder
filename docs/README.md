# 👀 City League Finder
[![.github/workflows/test-and-lint.yaml](https://github.com/owlinux1000/city-league-finder/actions/workflows/test-and-lint.yaml/badge.svg)](https://github.com/owlinux1000/city-league-finder/actions/workflows/test-and-lint.yaml) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Find open slots for the current **Pokémon Card City League**—and get notified instantly.  
Stay ahead of the crowd with results delivered straight to your favorite tools.

---

## 🚀 Quick Start

### 1. Install

```sh
go install github.com/owlinux1000/city-league-finder@latest
```

### 2. Create a Config file

> [!IMPORTANT]
> notifier doesn't work unless you set `true` to the `enabled` field.

```yaml
prefecture:
  - Tokyo
  - Osaka
league:
  - open
  - junior
notifier:
  slack:
    channel: "#test"
    memberID: test
  discord:  
    webhook: https://example.com
    memberID: test
    enabled: true
```

### 🔔 Notifier Setup
#### Slack

To send results to Slack, you’ll need a Slack App.
Make sure it has the `chat:write` OAuth scope.

To mention a specific user, set the memberID field.
You can find your member ID in your Slack profile (via browser).

#### Discord

For Discord, this tool uses webhooks.
Grab a webhook URL for the channel you want to post to:  
👉 [Intro to Webhooks](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks)

To mention a specific user, enable Developer Mode in Discord → right-click the user → Copy User ID.  
👉 [Where to find User/Server/Message IDs](https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID)


### 3. Run

```sh
$ export CONFIG_PATH=/path/to/config.yaml
$ export SLACK_TOKEN=xoxb-<censored>   # Required if using Slack
$ city-league-finder
```
