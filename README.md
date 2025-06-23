# FX Bot
A small [telegram](https://telegram.org) written in [golang](https://golang.org) to answer the conversion rate between different currencies from MasterCard.

[![License](https://img.shields.io/github/license/locnh/fxbot)](/LICENSE)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/locnh/fxbot?sort=semver)](/Dockerfile)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/locnh/fxbot?sort=semver)](/Dockerfile)
[![Docker](https://img.shields.io/docker/pulls/locnh/fxbot)](https://hub.docker.com/r/locnh/fxbot)

## Usage
### Run a Docker container

Default production mode

```sh
docker run --name fxbot --restart unless-stopped -e BOT_API_KEY=<YOU_KEY_HERE> -d locnh/fxbot
```

## Contribute
1. Fork me
2. Make changes
3. Create pull request
4. Grab a cup of tee and enjoy
