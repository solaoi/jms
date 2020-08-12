# WHAT IS JMS

this is a WIP project.

Jms is a JSON Management System.   
This application is a tool to generate and serve JSON :)

1. we make a form-template and a form.
2. we get JSON from a form.
3. we serve JSON.

# DEVELOPMENT

## RUN

```zsh
# backend
cd backend/ && go build
cd backend/ && ./jms run

# frontend
cd frontend/ && npm run dev
```

## MIGRATION
```zsh
# Up migration
cd backend/cmd/ && goose up
# Down Migration
cd backend/cmd/ && goose down
```

## UPDATE GRAPHQL SCHEMA
```zsh
vi backend/cmd/graph/schema.graphqls

# reflect this schema.graphqls changes
cd backend/cmd/ && gqlgen

# fix resolver implementations in order to use gorm
vi backend/cmd/graph/schema.resolvers.go
```

# BUILD

```zsh
# backend
cd backend/ && go build

# frontend
cd frontend/ && npm run export
# then serve out/
# example:
npm i -g http-server
cd frontend/out/ && http-server
```

# INSTALL

```zsh
cd backend/ && go install
```