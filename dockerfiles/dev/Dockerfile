FROM golang:1.8.3-alpine

ARG app
ARG github_token
ENV APP_DIR $GOPATH/src/github.com/eyecuelab/$app

RUN apk add --update --no-cache git
RUN git config --global url."https://$github_token:x-oauth-basic@github.com/".insteadOf "https://github.com/"

RUN go get -u github.com/golang/dep && go install github.com/golang/dep/cmd/dep
RUN go get -u github.com/pilu/fresh

WORKDIR $APP_DIR
RUN go get -u github.com/eyecuelab/$app

RUN dep ensure -update -v

CMD fresh -c runner.conf main.go

EXPOSE 3000
