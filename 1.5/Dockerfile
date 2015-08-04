FROM golang:1.5

# This is required for new Glide
ENV GO15VENDOREXPERIMENT 1

# Fetch glide and build it. Since golang:1.5 puts $GOPATH/bin in the $PATH,
# we don't have to do that here.
RUN go get -u -d github.com/Masterminds/glide
WORKDIR /go/src/github.com/Masterminds/glide
RUN git checkout feature/go15
RUN make bootstrap build install
WORKDIR /go
