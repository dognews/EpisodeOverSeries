<div align="center">
    <h1>Episode Over Series</h1>
    <img width="1467" height="165" alt="trimed" src="https://github.com/user-attachments/assets/480391bf-cc91-4fca-befd-ea6eccb63f29" />
    Force Emby to use the episode thumbnail instead of the series thumbnail.
</div>

## Quick Start
> [!CAUTION]
> This program should run on the same network as your Emby server, and use its local IP address.
### Requirements
* go (>=1.26)
### Build
```
go build -o ./build/project ./cmd/project
```
### Run
```
./build/project [flags] [command]
```
## CLI Usage
```
Usage:
    me.exe [flags] [command]
Commands:
     applytoseries [seriesid]    Apply episode thumbnails to series
     removefromseries [seriesid] Remove episode thumbnails from series
     applytoall                  Apply episode thumbnails to every series
     removefromall               Remove episode thumbnails from every series
     help                        Show cli commands
Flags:
    -ip [ip-address]  The ip address of the emby server  (REQUIRED)
    -apikey [api-key] The api key of the emby server     (REQUIRED)
```
