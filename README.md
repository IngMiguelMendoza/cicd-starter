# learn-cicd-starter (Notely)

This repo contains the starter code for the "Notely" application for the "Learn CICD"

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. I

![CI status](https://github.com/IngMiguelMendoza/cicd-starter/actions/workflows/ci.yml/badge.svg)
