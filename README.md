# KSST - Knuck Sandwich Save Tool

A GUI tool to edit Knuckle Sandwich save files!

<p align="center">
  <img width="776" height="736" src="/assets/screenshot.png">
</p>

Note: Use this tool at your own risk. I'm not responsible if you lose any data. Always make a backup before editing a save.

## How to use

1. Make a backup save
2. Make sure the game is closed
3. Load a save file
4. Edit
5. Overwrite old save file with new one
6. Remove savefile.ini
7. Enjoy!

Save files can be located in:

Windows (I think): %APPDATA%/local/Knuckle_Sandwich

Linux: Steam/steamapps/compatdata/pfc/drive_c/users/steamuser/AppData/local/Knuckle_Sandwich

This tool currently only work with `knucklesandwich.sav` files. Supporting `savefile.ini` will happen once that format is mature and `.sav` a file are no longer generated.

## Roadmap

- [x] Name
- [x] Money
- [ ] Stats
    - [x] Protagonist
    - [x] Thea
    - [x] Echo
    - [x] Dolus
    - [ ] Bside
- [x] Bonds
- [ ] Pocket Items
    - [x] Protagonist
    - [x] Thea
    - [x] Echo
    - [x] Dolus
    - [ ] Bside
- [ ] PC Items
- [ ] Location
- [ ] Story Flags
- [ ] Learnt Skill Flags
- [ ] Event Flags
- [ ] Misc Flags
- [ ] Fishing Info
- [ ] Evs and Ivs (no idea what they do)
- [ ] .ini file support (if required)

## Development

Make sure you have everything required for [fyne](https://github.com/fyne-io/fyne) to work.

Run `make run`.

## Linux Install Instructions

Download the latest release, uncompress and run sudo make install.

