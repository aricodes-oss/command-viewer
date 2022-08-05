FROM golang:latest as build

ENV CGO_ENABLED 0
RUN mkdir /comviewer
WORKDIR /comviewer

COPY . .
RUN go build

FROM scratch

ENV GIN_MODE release
COPY --from=build comviewer/templates ./templates
COPY --from=build comviewer/comviewer .

ENTRYPOINT ["./comviewer"]
