FROM technosophos/goglide:1.5-0

# This follows the pattern from golang:*-onbuild
RUN mkdir -p /go/src/app
WORKDIR /go/src/app

CMD ["go-wrapper", "run"]

ONBUILD COPY . /go/src/app
ONBUILD RUN go-wrapper download
ONBUILD RUN glide up
ONBUILD RUN go-wrapper install
