# website

The current "Schulen Im Chaos" Website and In-Memory-Database-System using go, gin and htmx.

## Usage

- **Install Perquisites:** You have to have [go](https://go.dev/), [air](https://github.com/cosmtrek/air) and [bun](https://bun.sh/) installed.
- **Install Dependencies:** Install dependencies of go (in [root](/)) and bun (in [content](static/content/)).
- **Run Dev:** Finally, You have to use the `air` command, it's pre-configured in the [air-toml](.air.toml).
- **Build and Bundle:** To build and bundle the project you have to run the export script: `./export_project.sh`. *Don't forget to set the permissions to run this file using `chmod`.*
