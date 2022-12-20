FROM golang:1.19-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o /Junaid-The-Controller

##
## Deploy
##

FROM gcr.io/distroless/base-debian10:debug

WORKDIR /
COPY --from=build /Junaid-The-Controller /Junaid-The-Controller

ENTRYPOINT ["/Junaid-The-Controller"]