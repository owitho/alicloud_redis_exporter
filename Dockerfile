FROM golang:1.10-alpine as builder
ENV APP_DIR $GOPATH/src/github.com/owitho/alicloud_redis_exporter
RUN mkdir -p $APP_DIR
ADD . $APP_DIR
WORKDIR $APP_DIR
RUN go build . 
RUN cp $APP_DIR/alicloud_redis_exporter /bin/alicloud_redis_exporter

FROM alpine:latest
COPY --from=builder /bin/alicloud_redis_exporter /bin/
EXPOSE 9456
CMD /bin/alicloud_redis_exporter
