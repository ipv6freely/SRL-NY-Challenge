# SRL-NY-Challenge

A tiny Go CLI that embeds SR Linux front-panel JPGs directly into the binary and renders them inline in modern terminals (Kitty protocol or iTerm2/WezTerm inline images).

## Requirements

- **Go:** **1.20+** (recommended: **Go 1.24**)
- A modern terminal emulator for image rendering (examples: **Ghostty**, **WezTerm**, **Kitty**, **iTerm2**)

Check your Go version:

```bash
go version
```

## Project Layout
```text
.
├── assets
│   ├── 7220_IXR-D3L.jpg
│   └── 7220_IXR-D5.jpg
└── main.go
```

## Install dependencies

1. Initialize a Go module
```bash
go mod init srl-ny-challenge
```

2. Add required dependencies
```bash
go get github.com/BourgeoisBear/rasterm
go get golang.org/x/term
```

## Run (development)
```bash
go run . d3l
go run . d5
```

## Build a local binary (development)
```bash
go build -o srl-ny-challenge
```

## Run the local binary (development)
```bash
./srl-ny-challenge d3l
./srl-ny-challenge d5
```

## Build a binary for x86_64
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o srl-ny-challenge
```

## Copy to SRL
```bash
scp srl-ny-challenge linuxadmin@srl-router:/tmp
```

## Run on SRL
```bash
/tmp/srl-ny-challenge d3l
/tmp/srl-ny-challenge d5
```