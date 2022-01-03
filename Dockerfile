FROM golang:latest


RUN go get -u github.com/gin-gonic/gin && \
    go get -u gorm.io/gorm && \
    go get -u gorm.io/driver/postgres && \
    go get -u github.com/cosmtrek/air && \
    go get golang.org/x/crypto/bcrypt

# RUN go mod download



WORKDIR /go/src/sotsusei/src


# CMD ["go", "run", "main.go"]
CMD ["air", "-c", ".air.toml"]
