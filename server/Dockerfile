FROM rust:1.40.0 AS development

WORKDIR /home/rust

RUN cargo install --version 7.3.0 cargo-watch

COPY src ./src
COPY Cargo.toml Cargo.lock ./

RUN cargo build

CMD cargo-watch -x check -x test -x run

# Build Release #############################

FROM rust:1.40.0 AS build-release

WORKDIR /home/rust

COPY --from=development /home/rust/src ./src

COPY Cargo.toml Cargo.lock ./

RUN cargo build --release

# Production ################################

FROM python:3.8.0 AS production

RUN useradd -m vgen

WORKDIR /home/vgen

# COPY --chown=vgen:vgen --from=build /home/rust/target/

RUN pip install ansible==2.8.3

USER vgen

# CMD [ "/home/vgen/vault-generator" ]
