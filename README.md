# Api Gateway Lamba Go Example

## Pre-requisites

`docker version 17+` See how to download and install in [Docker site.](https://docs.docker.com/install/linux/docker-ce/ubuntu/)

`docker-compose version 1.20+` See how to download and install in [Docker site.](https://docs.docker.com/compose/install/#install-compose)

`golang version 1.11+`  See how to download and install in [Golang site.](https://golang.org/doc/install)

`awscli-local` See how to download and install in [Aws cli github](https://github.com/localstack/awscli-local)

---

## Environment variables

Copy of `.env.sample` as `.env` and change to your needs.

```bash
cp .env.sample .env
```

---

## Development

Run:

```bash
make up
```

### Tests

#### Unit

Unit tests are done with native [GO language tests](https://golang.org/pkg/testing/) and they are scattered with source code with same name of file to be tested, but a `_test` suffix.

To run unit tests:
```bash
make test
```
---
