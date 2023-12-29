## Example of GO server using Templ and Tailwind

Created a simple todo list app.  
Besides the app itself, there is also some tools and code for better developer experience.

### Features

- server side rendering with **types** using the power of [templ](https://github.com/a-h/templ)
- super fast web building using [htmx](https://htmx.org/) and [tailwindcss](https://tailwindcss.com/)
- automatic refresh of UI using websockets
- automatic restart of server and building of templates using nodemon
- addition of tailwind.config.js for tailwind auto-complete (recommended vs extensions: [tailwind](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss), [templ](https://marketplace.visualstudio.com/items?itemName=a-h.templ))

### How to generate templates:

```bash
~/go/bin/templ generate
```

### How to run the server:

```bash
go run main.go
```

### Automate server restart and templ builds upon file changes:

1. install nodemon globally (or locally):

```bash
npm install -g nodemon
```

2. run bash script:

```bash
./run.sh
```
