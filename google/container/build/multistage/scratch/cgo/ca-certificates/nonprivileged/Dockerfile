
### Build stage

FROM golang as build

WORKDIR /go/src
ADD main.go .

RUN CGO_ENABLED=0 go build -o google .
RUN cat /etc/passwd | grep nobody > passwd.nobody


### Run stage

FROM scratch

COPY --from=build /go/src/google .
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/passwd.nobody /etc/passwd

USER nobody
CMD ["./google"]
