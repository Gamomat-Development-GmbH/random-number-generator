FROM debian:buster-slim

RUN apt-get update && apt-get install -y \
  dieharder \
  curl \
  xxd

COPY entrypoint.sh .

ENTRYPOINT ["/entrypoint.sh"]
