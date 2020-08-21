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
# *** BACKEND ***
# we should frontend binarize firstly.
# then run below commands
cd backend/ && go build && ./jms init && ./jms start

# *** FRONTEND ONLY ***
cd frontend/ && npm run dev
```

## FRONTEND BINARIZE

```zsh
cd [root]
npm run export --prefix frontend && 
mv frontend/out backend/ &&
go-bindata -pkg=static -o=backend/static/out.go -prefix=backend/ ./backend/out/...
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
# *** BACKEND ***
# we should frontend binarize firstly.
# then run below commands
cd backend/ && go build

# *** FRONTEND ONLY ***
cd frontend/ && npm run export
# then serve out/
# example:
npm i -g http-server
cd frontend/out/ && http-server
```

# INSTALL

```zsh
# we should frontend binarize firstly.
# then run below commands
cd backend/ && go install
```
