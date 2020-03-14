FROM golang:1.13 AS build

COPY smtptest.go /tmp/smtptest.go

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' \
    -o /tmp/smtptest /tmp/smtptest.go

FROM scratch

COPY --from=build /tmp/smtptest /smtptest

ENTRYPOINT ["/smtptest"]
