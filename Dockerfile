FROM golang:1.21.0-alpine3.18 AS build

WORKDIR /app

# COPY go.mod go.sum ./
COPY go.mod .

RUN go mod download

# RUN go mod verify

COPY . .

RUN go build -ldflags "-w -s" -o /app-bin

FROM scratch

WORKDIR /app

COPY --from=build /app-bin /app/minaris 

COPY --from=build /app/raws /app/raws

ENTRYPOINT ["/app/minaris"]