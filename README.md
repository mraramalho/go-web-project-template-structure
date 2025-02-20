# What's it?

This is a template structure for a web project using Go.

# How to use it?

Clone this repository and use it as a initial template for your project.

# What's inside?

```
|-- cmd
| `-- web
|       |-- main.go
|       |-- middlewares.go
|       `-- routes.go
|-- config
| `-- config.go
|-- pkg
|   |-- handlers
|   |   `-- handlers.go
| |-- models
| | `-- templatedata.go
|   `-- render
| `-- render.go
|-- readme.md
`-- templates
|-- about.page.tmpl
|-- base.layout.tmpl
|-- contact.page.tmpl
`-- home.page.tmpl
```

# What does every file do?

main.go

- main.go is the entry point of the application.
- It will be used to start the application.

middlewares.go

- middlewares for the application.

routes.go

- routes for the application.

config.go

- config for the application.

handler.go

- all handlers for the application go here.

templatedata.go

- all models for the application go here.

render.go

- all render for the application go here.

templates .tmpl

- all templates for the application go here.
- for base layout, create a file called base.layout.tmpl
- partial layouts, create files called file_name.layout.tmpl
- for pages, create files called file_name.page.tmpl

  base.layout.tmpl

  - base layout for the application.
    home.page.tmpl
  - home template for the application.
    about.page.tmpl
  - about template
