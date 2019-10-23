ARG IMAGE_NAME

FROM ${IMAGE_NAME}

ARG ENVIRONMENT
ARG APP_NAME
ARG PORT

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

RUN apk --update add --no-cache bash gcc musl-dev git mercurial openssh curl && \
  rm -rf /var/lib/apt/lists/*
  # rm /var/cache/apk/*

# SSH for private repository
# RUN mkdir -p /root/.ssh && \
#     echo "$SSH_KEY" > /root/.ssh/id_rsa && \
#     chmod 0600 /root/.ssh/id_rsa && \
#     eval `ssh-agent` && \
#     ssh-add /root/.ssh/id_rsa && \
#     ssh-keyscan github.com >> /root/.ssh/known_hosts && \
#     echo -e "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig && \
#     echo -e "Host github.com\n\tStrictHostKeyChecking no\n" >> /root/.ssh/config

RUN go get -u github.com/golang/dep/cmd/dep && \
    go get -u bitbucket.org/liamstask/goose/cmd/goose

# Fresh for rebuild on code change, no need for production
RUN go get -u github.com/pilu/fresh

WORKDIR $GOPATH/src/${APP_NAME}/app
ENV WORK_DIR $GOPATH/src/${APP_NAME}/app

COPY . ./

RUN dep ensure -v -vendor-only=true
# RUN dep ensure -v 
RUN    go build 
    # rm /root/.ssh/id_rsa && \
    # unset SSH_KEY && \
    # cd $WORK_DIR

EXPOSE ${PORT}

CMD fresh
