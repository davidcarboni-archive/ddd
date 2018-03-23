
### Build stage

FROM golang as build

WORKDIR /go/src
ADD main.go .

RUN go build -o google .


### Run stage

FROM golang

WORKDIR /app
COPY --from=build /go/src/google .

CMD ["./google"]
