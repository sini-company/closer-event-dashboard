FROM golang:1.10-alpine AS build

# Install tools required to build the project
# We will need to run `docker build --no-cache .` to update those dependencies
RUN apk add --no-cache git
RUN go get github.com/golang/dep/cmd/dep

# Gopkg.toml and Gopkg.lock lists project dependencies
# These layers will only be re-built when Gopkg files are updated
COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /go/src/project/
RUN dep ensure -vendor-only

# Copy all project and build it
# This layer will be rebuilt when ever a file has changed in the project directory
COPY . /go/src/project/

# Install library dependencies
# RUN go build -o /bin/app (only work in same linux environment)
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /bin/healthcheck ./healthcheck
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o /bin/app

# This results in a single layer image
FROM scratch
ENV PORT=3000
COPY --from=build /bin /bin
HEALTHCHECK CMD ["/bin/healthcheck"]
ENTRYPOINT ["/bin/app"]