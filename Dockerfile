# build stage
FROM golang:alpine AS build-env
ARG APP_NAME
ADD ./golangserver /src
ADD ./configs/golangserver /app


WORKDIR /src/server/${APP_NAME}
RUN go build -o /app



# final stage
FROM alpine
ARG APP_NAME
WORKDIR /app
RUN apk update && apk add tzdata
COPY --from=build-env /app/${APP_NAME} /app/${APP_NAME}


ENTRYPOINT /app/${APP_NAME}