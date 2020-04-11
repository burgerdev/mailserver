FROM golang:1.13 as build

ADD . /src

RUN cd /src && make static

FROM scratch

COPY --from=build /src/static /x

ENTRYPOINT ["/x"]

