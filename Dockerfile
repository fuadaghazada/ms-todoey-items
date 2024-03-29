# syntax=docker/dockerfile:1

#########
#   Build
#########
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /ms-todoey-items

##########
#   Deploy
##########
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /ms-todoey-items /ms-todoey-items

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/ms-todoey-items"]
