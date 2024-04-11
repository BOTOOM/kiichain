FROM golang

RUN apt update && apt upgrade -y && curl https://get.ignite.com/cli@v0.27.1! | bash

WORKDIR /app

COPY . .

EXPOSE 1317 26656 26657 26658 6060 8080 9090 9091

CMD ["bash"]
