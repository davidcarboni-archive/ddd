
### Build stage

FROM golang as build

WORKDIR /go/src
ADD main.go .

RUN go build -o google .

### Run stage

FROM scratch

COPY --from=build /go/src/google .

CMD ["./google"]
