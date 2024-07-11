FROM golang:1.22.5-alpine as build

WORKDIR /usr/local/go/src/GoWB/
ADD . .
RUN CGO_ENABLED=0 go build -trimpath -v -a -o GoWB -ldflags="-w -s"

FROM scratch

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/local/go/src/GoWB/GoWB /

ENTRYPOINT ["./GoWB"]