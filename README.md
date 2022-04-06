# CSGO.GO Client

## About

This is a light-weight CSGO hack client made in GoLang that adds some features like wallhack, aimbot and triggerbot for study purposes only.

As I said, this is for study purposes only, and I don't recommend to use this on your main account or any account that you cares. If you get banned, this is not my responsability. Counter-Strike Global Offensive have cleary terms that rejects any type of cheating in online matches and it should be respected.

This client don't have any relations to Valve or similar, and use this on online matches isn't allowed officially by Valve.

## Features

|Option | Description  |
|---|---|
|Glow| Draw and outline around your enemies|
|ESP Visuals| OpenGL overlay that render in-game info about your enemies|
|Auto Weapons| Auto shot for weapons that isn't full auto (pistols)|
|Radar Hack| Spot enemies on your in-game radar|
|Bunnyhop| Auto-jump based bhop to allow you better movements|
|Triggerbot| Auto shoot against your enemies when someone is in your crosshair|
|Aimbot| Auto aim at your enemies to improve your accuracy|

> Now, configs are available on your default documents folder

## Currently known issues

- Name ESP is drawn out of center in some circunstances (visual impact only/ low priority)

## CGO Issues relatable commands

```shell
$ go env -w CGO_ENABLED=1
$ go env -w CGO_CFLAGS="-g -O2 -m32 -w"
$ go env -w CGO_CPPFLAGS="-g -O2 -m32 -w"
```
