FROM golang:1.13.5 AS build

WORKDIR /go/src

RUN useradd go -M

COPY ./src /go/src/vault-generator

RUN ls -l /go/src/vault-generator

RUN go install vault-generator

# Development ############################

FROM python:3.8.0 AS development

RUN useradd -m vgen

WORKDIR /home/vgen

COPY --chown=vgen:vgen --from=build /go/bin/vault-generator .

RUN pip install ansible==2.8.3

USER vgen

CMD [ "/home/vgen/vault-generator" ]

# Production #############################
# Try to reduce image size
