*Migrated to [Codeberg.org](https://codeberg.org/tyybbi/guiching)*

# Guiching

A simple I Ching hexagram generator, that uses Go backend and template
to serve stuff onto an HTML page.

## Usage

Run **guiching** by `go build && ./guiching` or just by `go run *.go` in the project
root directory. Then, in a web browser, visit http://localhost:3000 to
see your newly generated hexagram(s).

## Docker

1. Build docker image: `docker build -t guiching .`
2. Run docker container: `docker run -it --rm -p 80:3000 guiching:latest`
3. Visit http://localhost in your chosen web browser

