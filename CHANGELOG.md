# Changelog for Client Updates

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
