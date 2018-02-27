FROM golang:1.7.3 as build
WORKDIR /go/src/github.com/gfk-ba/aws-ecs-experiment/quote-service
COPY *.go ./
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest as production
#RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/gfk-ba/aws-ecs-experiment/quote-service/app .
EXPOSE 8080
ENTRYPOINT ["./app"]
