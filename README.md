# Hexing

A simple I Ching hexagram generator, that uses Go backend and template
to serve stuff onto an HTML page.

## Usage

Run *hexing* by `go build && ./hexing` or just by `go run *.go` in the project
root directory. Then, in a web browser, visit http://localhost:8080 to
see your newly generated hexagram(s).

### Couple of curiosities

- The app does not use any JavaScript/Ajax etc. so just refresh the page to generate
  new shapes.
- The coins method of randomization never gets used because nothing can change
  the boolean value to true. This may change in the future if I'm willing to
  bring some JS into the picture.
- I'm still very new to Golang, the code quality probably reflects that.

