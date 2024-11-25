FROM golang:1.20 AS base

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


#final stage - Distroless image
FROM gcr.io/distroless/static:nonroot

COPY --from=base /app/main .

COPY --from=base /app/static ./static 

EXPOSE 8080

CMD [ "./main" ]
