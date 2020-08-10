# WHAT IS JMS

this is a WIP project.

jms is the JsonManagementSystem.

1. we make a form-template and a form.

2. we get JSON from a form.

# DEVELOPMENT

## RUN

```zsh
# backend
cd backend/ && go run ./main.go

# frontend
cd frontend/ && npm run dev
```

## MIGRATION
```zsh
# Up migration
cd backend/ && goose up
# Down Migration
cd backend/ && goose down
```

## UPDATE GRAPHQL SCHEMA
```zsh
vi backend/graph/schema.graphqls

# reflect this schema.graphqls changes
cd backend/ && gqlgen

# fix resolver implementations in order to use gorm
vi backend/graph/schema.resolvers.go
```

# BUILD

```zsh
# backend
cd backend/ && go build -o jms

# frontend
cd frontend/ && npm run export
# then serve out/
# example:
npm i -g http-server
cd frontend/out/ && http-server
```
