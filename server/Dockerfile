FROM golang:1.13.5 AS development

ENV environment=dev

WORKDIR /go/src/vault-generator

# Install entr
RUN curl -L http://eradman.com/entrproject/code/entr-4.3.tar.gz -o /tmp/entr-4.3 \
    && tar xvf /tmp/entr-4.3 -C /tmp \
    && cd /tmp/entr-4.3 \
    && ./configure; make \
    && ln -s /tmp/entr-4.3/entr /usr/local/bin

COPY ./src /go/src/vault-generator

RUN go test vault-generator/...
RUN go install -i vault-generator/main

CMD find -name "*.go" | entr -r sh -c "go test vault-generator/... && go run vault-generator/main"

# Production #############################
# TODO :: Try to reduce image size
FROM python:3.8.0 AS production

ENV environment=prod

RUN useradd -m vgen

WORKDIR /home/vgen

COPY --chown=vgen:vgen --from=development /go/bin/main .

RUN pip install ansible==2.8.3

USER vgen

CMD [ "/home/vgen/vault-generator" ]
