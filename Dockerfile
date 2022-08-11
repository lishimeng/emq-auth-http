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
RUN go build -v --ldflags "-X cmd.AppName=${NAME} -X cmd.Version=${VERSION} -X cmd.Commit=${COMMIT} -X cmd.Build=${BUILD_TIME}" -o ${NAME} ${MAIN_PATH}

FROM ubuntu:22.04 as prod
ARG NAME
EXPOSE 80/tcp
WORKDIR /
COPY --from=build /release/${NAME} /
RUN ln -s /${NAME} /app
CMD [ "/app"]
