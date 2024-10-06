FROM rust:1-slim

RUN apt-get update && apt-get install -y tini

RUN cargo install spf-milter

ENTRYPOINT ["tini", "/usr/local/cargo/bin/spf-milter"]
