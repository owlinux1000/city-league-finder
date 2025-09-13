# :eyes: City League Finder

This program aims to find the open slot of the current Pokemon Card City League. You can receive the result via some tools.

## How to use

### 1. Install this tool

```
go install github.com/owlinux1000/city-league-finder@latest
```

### 2. Create a config file

> [!IMPORTANT]
> You need to define one of notifier and its configuration.

```yaml
prefecture:
  - Kanagawa
  - Tokyo
notifier: [slack, discord]
slack:
  channel: "#test"
  # memberID: test
discord:
  webhook: https://example.com
  # memberID: test
```

#### 2.1 Slack

If you want to notify the result to Slack, you need to setup [Slack App](https://api.slack.com/apps). The required scopes of OAuth is `chat:write`. 

Also you can mention a specific user by setting `memberID` field. The memberID is that you can get your profile in your browswer.

#### 2.2 Discord

To notify the result to Discord, this tool uses Webhook. Therefore, you need to get a Webhook URL of the channel you want to post. The ways to get the URL is explained this official page. 
https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks

Also you can mention a specific user by setting `memberID` field. Once you enable the developer mode of Discord app from settings, you can see `Copy User ID` by right-click a user.
https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID


### 3. Run this tool with the config

```sh
$ export CONFIG_PATH=/path/to/config.yaml
$ export SLACK_TOKEN=xoxb-<censord> # When you want to use Slack notifier
$ city-league-cancel-detector
```
