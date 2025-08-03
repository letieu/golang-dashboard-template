# Dashboard app template
Starter point for build a fullstack dashboard app use:

- golang
- gin
- templ
- html

## Run
```sh
# run
templ generate -watch -cmd "go run cmd/web/main.go"

# compile tailwindcss
tailwindcss -i web/styles/input.css -o web/static/main.css -w
```
