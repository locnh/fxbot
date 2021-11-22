# MasterFX Bot
A small [telegram](https://telegram.org) written in [golang](https://golang.org) to answer the conversion rate between different currencies from MasterCard.

Live here: [@mastercardfxbot](https://t.me/mastercardfxbot)

These are the Docker Hub autobuild images located [here](https://hub.docker.com/r/locnh/mastercardfxbot/).

[![License](https://img.shields.io/github/license/locnh/mastercardfxbot)](/LICENSE)
[![Build Status](https://travis-ci.com/locnh/mastercardfxbot.svg?branch=master)](https://travis-ci.com/locnh/mastercardfxbot)
[![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/locnh/mastercardfxbot?sort=semver)](/Dockerfile)
[![Docker Image Version (latest semver)](https://img.shields.io/docker/v/locnh/mastercardfxbot?sort=semver)](/Dockerfile)
[![Docker](https://img.shields.io/docker/pulls/locnh/mastercardfxbot)](https://hub.docker.com/r/locnh/mastercardfxbot)

## Usage
### Run a Docker container

Default production mode

```sh
docker run --name masterfxbot --restart unless-stopped -e BOT_API_KEY=<YOU_KEY_HERE> -d locnh/mastercardfxbot
```

## Contribute
1. Fork me
2. Make changes
3. Create pull request
4. Grab a cup of tee and enjoy
