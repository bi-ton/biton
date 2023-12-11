FROM golang:1.21.4-alpine AS build
RUN apk update
RUN apk add git
LABEL stage=build
RUN go install github.com/msw-x/vgen/cmd/vgen@latest
WORKDIR /app
COPY . ./
RUN vgen go src
WORKDIR /app/src
RUN go mod download
RUN go build -o /biton

FROM gcr.io/distroless/base-debian10
COPY --from=build /biton ./
COPY biton.conf ./
ENTRYPOINT ["/biton"]
