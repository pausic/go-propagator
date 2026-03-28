FROM golang:1.26 AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /go-propagator ./cmd/go-propagator

FROM gcr.io/distroless/static-debian12
COPY --from=build /go-propagator /go-propagator
ENTRYPOINT ["/go-propagator"]
