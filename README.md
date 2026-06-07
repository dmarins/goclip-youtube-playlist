# goclip-youtube-playlist

Retrieves all video URLs from a YouTube playlist and copies them to the clipboard, one URL per line.

## Requirements

- Go 1.25+
- A [YouTube Data API v3](https://console.developers.google.com/) key

## Setup

1. Clone the repository and install dependencies:

```sh
git clone https://github.com/dmarins/goclip-youtube-playlist.git
cd goclip-youtube-playlist
go mod download
```

2. Copy `.env.example` to `.env` and fill in your API key:

```sh
cp .env.example .env
# edit .env and set YOUTUBE_API_KEY=<your key>
```

## Usage

```sh
go run . -playlist <PLAYLIST_ID>
```

| Flag         | Description                              | Default |
|--------------|------------------------------------------|---------|
| `-playlist`  | YouTube playlist ID **(required)**       | —       |
| `-env`       | Path to the `.env` file                  | `.env`  |

### Example

```sh
go run . -playlist PLxxxxxxxxxxxxxxxx
```

The tool prints every video URL to stdout (one per line) and copies the full list to the clipboard.

## Build

```sh
go build -o goclip-youtube-playlist .
./goclip-youtube-playlist -playlist PLxxxxxxxxxxxxxxxx
```
