## Example of GO server using Templ and Tailwind

Created a simple todo list app.  
Besides the app itself, there is also some tools and code for better developer experience.

### Features

- server side rendering with **types** using the power of [templ](https://github.com/a-h/templ)
- super fast web building using [htmx](https://htmx.org/) and [tailwindcss](https://tailwindcss.com/)
- automatic refresh of UI using websockets
- automatic restart of server and building of templates using nodemon
- addition of tailwind.config.js for tailwind auto-complete (recommended vs extensions: [tailwind](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss), [templ](https://marketplace.visualstudio.com/items?itemName=a-h.templ))
- both local rest + AWS lambda integration

### How to generate templates:

```bash
templ generate
```

### How to run the server:

Locally:

```bash
go run main.go
```

or

Aws sam (aws lambda simulation):

```bash
sam local start-api
```

### Automate server restart and templ builds upon file changes:

Install nodemon globally (or locally):

```bash
npm install -g nodemon
```

Auto-build templates:

```bash
 ./scripts/build-templates.sh
```

Auto-start server on file changes:

```bash
./scripts/start.sh
```

Building zip for AWS:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go &&  zip -j function.zip main public/*
```
