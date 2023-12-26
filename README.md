# htmx

A rework of the current website and Database-System using go, gorm and htmx.

## Usage

- **Install Perquisites:** You have to have [go](https://go.dev/), [air](https://github.com/cosmtrek/air) and [bun](https://bun.sh/) installed.
- **Install Dependencies:** Install dependencies of go (in [root](/)) and bun (in [content](static/content/)).
- **Run Dev:** Finally, You have to use the `air` command, it's pre-configured in the [air-toml](.air.toml).
- **Build:** To build the project you have to run the following command, **make sure to include in your export the static files**:
```sh
cd ./static/content && bunx tailwindcss -i ./../../web/input.css -o ./dist/output.css --minify && cd ./../../ && templ generate && go build .
```
*A very similar command is using `air` for running the project in dev-mode*
