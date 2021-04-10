# Domane

[![Api](https://img.shields.io/github/workflow/status/kostaspt/domane/Api?label=Api)](https://github.com/kostaspt/domane/actions/workflows/api.yml)
[![Client](https://img.shields.io/github/workflow/status/kostaspt/domane/Client?label=Client)](https://github.com/kostaspt/domane/actions/workflows/client.yml)
[![Server](https://img.shields.io/github/workflow/status/kostaspt/domane/Server?label=Server)](https://github.com/kostaspt/domane/actions/workflows/server.yml)

### How to run locally

#### Requirements:
* [Docker](https://docs.docker.com/engine/install/)
* [Docker Compose](https://docs.docker.com/compose/install/)

You'll need to create an `.env` file. This file will contain any environment variables needed for it to run locally. There is already a `.env.example` file with all the needed values you'll need.

Additionally, you'll need to set up a domain name locally for the app to resolve. To do so, add these values to your [hosts file](https://en.wikipedia.org/wiki/Hosts_(file)).

```apache
127.0.0.1   domane.localhost
127.0.0.1   grafana.domane.localhost     # Optional
127.0.0.1   prometheus.domane.localhost  # Optional
```

*Be careful:* If in the you've changed the `DOMAIN` variable in `.env` file, you should use that instead of `domane.localhost`.

Once everything is set, run:
```shell
$ docker-compose up -d --build --force-recreate --remove-orphans

# Alternative, for unix with `make` installed:
$ make
```
