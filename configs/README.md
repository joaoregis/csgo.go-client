# Explanation about config files

## config.json
This is the main file for YOUR configs like keybinds and features enabled/disabled. Here you can change freely and change options as you like. Please, make sure to follow the patterns of JSON files, and use `#FFF000` for color patterns. Use `true/false` statements to choose between yes/no (enabling features for example). 

## csgo.json
This is the file that you SHOULD NOT TO EDIT, or maybe just prevent to edit if you don't know how to do it. The offsets file is the core of this client. The majority of game updates changes the offsets that allow us to read/write memory of the game. So it's need to be updated after game updates that change some core feature. To update this, we recommend to use HazeDumper, just start CSGO with `-insecure` launch option and runs HazeDumper after CSGO starts. It will generate a new `csgo.json` inside the folder of HazeDumper, and you can override the current csgo.json with the new generated csgo.json. If the cheat still broken, we need to update some patters scan, which is more complicated, so wait for the update.

## Other notes
We don't have a config updater yet. This means that if something changes in the config (like adding a new config option) it will break and you need to get a new config.json from repository. We working on it, and it still unfinished. Look at todo file to know what we working on.