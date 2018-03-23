
### Build stage

FROM golang as build

WORKDIR /go/src
ADD main.go .

RUN CGO_ENABLED=0 go build -o google .


### Run stage

FROM scratch

COPY --from=build /go/src/google .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["./google"]
