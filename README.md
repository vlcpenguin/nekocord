# Introduction
Minimal discord library written in GO (Nothing else to it)

It works with user accounts but i wouldn't recommend it since it goes against Discords TOS and can get your account terminated even tho i've never heard of a case where some had been self-botting and got banned for it without spamming their api.

Avoid the risk of getting banned and try to only using bot-tokens and not user-tokens when using this library.

> Lacks alot of the Discord API so not good for super huge projects.

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/c8a5f8ae69694958a75f83da04f0c5ca)](https://app.codacy.com/gh/vlcpenguin/nekocord/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
[![Maintainability](https://qlty.sh/gh/vlcpenguin/projects/nekocord/maintainability.svg)](https://qlty.sh/gh/vlcpenguin/projects/nekocord)
# Installation
```sh
go get github.com/vlcpenguin/nekocord@latest
```

# Example
```sh
git clone https://github.com/vlcpenguin/nekocord.git
cd nekocord/examples
go mod init main.go
go get github.com/vlcpenguin/nekocord@latest
go mod tidy
go build -o main
export discord_token="<discord_token_here>"
./main
```
## Output
> Ignore the error it's just because i haven't set the discord_token environment variable.
```sh
Cloning into 'nekocord'...
remote: Enumerating objects: 46, done.
remote: Counting objects: 100% (46/46), done.
remote: Compressing objects: 100% (38/38), done.
remote: Total 46 (delta 11), reused 4 (delta 1), pack-reused 0 (from 0)
Receiving objects: 100% (46/46), 26.24 KiB | 248.00 KiB/s, done.
Resolving deltas: 100% (11/11), done.
go: creating new go.mod: module main.go
go: to add module requirements and sums:
	go mod tidy
go: added github.com/gorilla/websocket v1.5.3
go: added github.com/vlcpenguin/nekocord v0.1.0
2026/05/16 18:55:44 [*] Connecting to discords gateway...
2026/05/16 18:55:44 [+] Connected
2026/05/16 18:55:44 [*] Sending identify payload...
2026/05/16 18:55:44 [+] Sent identify payload
2026/05/16 18:55:44 [*] Sending heartbeat... <nil>
2026/05/16 18:55:44 [+] Sent heartbeat
2026/05/16 18:55:44 [*] Waiting for server Acknowledgement...
2026/05/16 18:55:44 Error:websocket: close 4004: Authentication failed.
```

# PRS WELCOME
