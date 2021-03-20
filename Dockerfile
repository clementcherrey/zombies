# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o bin .

# final image
FROM alpine
WORKDIR /app
COPY --from=build-env /src/bin /app/
ENTRYPOINT ["/app/bin"]
