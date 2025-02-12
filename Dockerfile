FROM golang:1.23 AS build

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o st-app ./main.go

FROM golang:1.23 AS prod

WORKDIR /app
COPY --from=build /app/st-app .

ENTRYPOINT ["./st-app"]