FROM golang:1.18 as build
ARG NAME
ARG VERSION
ARG COMMIT
ARG TAG
ARG BUILD_TIME
ARG MAIN_PATH
ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /release
ADD . .
RUN go mod download && go mod verify
RUN go build -v --ldflags "-X cmd.Version=${VERSION}" -o ${NAME} ${MAIN_PATH}
FROM alpine:3.16 as prod
EXPOSE 80/tcp
WORKDIR /
COPY --from=build /release/${NAME} /
ENTRYPOINT ["/${NAME}"]