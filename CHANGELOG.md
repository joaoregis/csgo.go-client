# Changelog for Client Updates

## v0.1.0 (02-04-2022)

> Note: That is a BIG update since first release.

- All config system has been recoded from scratch.
- Added protected configs to prevent user changes (disabled for now, until custom GUI for configs be done)
- Added HWID-based configs, it means that your config is linked to your hardware id and always will be
- Discountinued support for old config formats
- Added config versioning to prevent fails due to not up-to-date configs
- Added encryption on a several code parts related to user data
- Now the `configs` folder doesn't exists anymore, and configs will be stored at default `Documents` path of your windows user
- Client has been changed to executes in x86 (32 bits) mode to prevent inconsistencies across process
- Added tons of checks and improvements across the code to prevent crashs and bugs
- Fixed an issue with AutoWeapons that make the feature force to firing when you shouldn't actually fire. This issue is mainly related to the buy menu of CSGO, but it occours always that mouse cursor were visible. Now, that was fixed.
- Randomized some strings and filenames on build to improve security.
- Added a lot of encrypting and hashing logic to several parts of the code to improve security against data leaks

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
