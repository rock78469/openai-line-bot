# build stage
FROM golang:alpine AS build-env
ARG APP_NAME
ADD ./ /src

WORKDIR /app
ADD ./.env /app

WORKDIR /src
RUN go mod vendor && go build -o /app


#CMD ["tail","/dev/null","-f"]

# final stage
FROM alpine
ARG APP_NAME
WORKDIR /app
RUN apk update && apk add tzdata
COPY --from=build-env /app /app
ENTRYPOINT /app/${APP_NAME} server