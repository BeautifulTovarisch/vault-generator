FROM golang:1.13.5 AS development

WORKDIR /go/src/vault-generator

# Install entr
RUN curl -L http://eradman.com/entrproject/code/entr-4.3.tar.gz -o /tmp/entr-4.3 \
    && tar xvf /tmp/entr-4.3 -C /tmp \
    && cd /tmp/entr-4.3 \
    && ./configure; make \
    && ln -s /tmp/entr-4.3/entr /usr/local/bin

COPY ./src /go/src/vault-generator

RUN go install -i

CMD /bin/ls * | entr -r sh -c "go test && go run vault-generator"

# Production #############################
# TODO :: Try to reduce image size
FROM python:3.8.0 AS production

RUN useradd -m vgen

WORKDIR /home/vgen

COPY --chown=vgen:vgen --from=development /go/bin/vault-generator .

RUN pip install ansible==2.8.3

USER vgen

CMD [ "/home/vgen/vault-generator" ]
