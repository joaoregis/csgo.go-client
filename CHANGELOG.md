# Changelog for Client Updates

## v0.0.5 (01-04-2022)

- Now, the client will check if CSGO is on focus (Foreground window) to prevent hack from perform invalid actions when the game is minimized
- Massively config changes to make it cleaner and understandable
- Improve mechanisms to detect when the client go exit and make a "graceful" exiting
- Improve aimbot system to make sure that active weapon is a valid weapon and prevent from aimboting using Zeus, Knifes and any other "invalid" weapon
- Improve trigger to make sure that active weapon is a valid weapon and prevent from triggering using Zeus, Knifes and any other "invalid" weapon (In the near future, the client will receive new features like "Auto Zeus" and "Auto Knifebot" to replace old mechanism with triggerbot)
- Added new feature "Automatic Weapons" that allows you to constantly fire weapons that are semi-auto (pistols) with configurable delay between shots

## v0.0.4 (03-31-2022)

- Compiled binaries for latest GoLang build (1.18)

## v0.0.3 (03-31-2022)

- Added pattern scanning for main features to keep game offsets up-to-date

## v0.0.2 (03-30-2022)

- Added new remote source for game offsets
- Improve enginechams to clear residues during client exiting
- Added mechanisms of graceful exiting when user aborts the program
- Added the same graceful behavior when csgo game closes

## v0.0.1 (03-30-2022)

- The very first release
- This is a migration project between repos, so from the previous verson we had adding some new things:
  - Improve hotkey to reload configs, preventing key press for too long
  - Added support for all relevant virtual keys from keyboard, available on `available_keys.txt` file
  - Moved offsets to a config file compatible with `Hazedumper` output in json format
  - Several of code cleanups and improves about QoL of code maintenance
