FROM rust:1.31.0 AS development

RUN useradd rust -m

WORKDIR /home/rust

RUN curl -L http://eradman.com/entrproject/code/entr-4.3.tar.gz -o entr-4.3 \
    && tar xvf entr-4.3 -C /tmp \
    && cd /tmp/entr-4.3 \
    && ./configure; make \
    && ln -s /tmp/entr-4.3/entr /usr/local/bin

COPY src ./src
COPY Cargo.toml Cargo.lock ./

RUN cargo build

CMD /bin/ls -d src/* | entr -s "cargo check && cargo test"

# Production #############################

FROM python:3.8.0 AS production

RUN useradd -m vgen

WORKDIR /home/vgen

# COPY --chown=vgen:vgen --from=build /home/rust/target/

RUN pip install ansible==2.8.3

USER vgen

# CMD [ "/home/vgen/vault-generator" ]
