# build stage
FROM golang:1.23-bookworm AS build

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o ./app ./cmd/main.go

# deploy stage
# FROM golang:1.23-alipine AS deploy
FROM gcr.io/distroless/base-debian12 AS deploy

WORKDIR /app

COPY --from=build /app/app ./

EXPOSE 8081

# needed for distroless base
USER nonroot:nonroot

ENTRYPOINT [ "./app" ]
