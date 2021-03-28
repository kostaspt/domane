# Domane

### How to run locally
It's pretty straightforward. First of all, make sure you have the right software installed.

#### Requirements:
* [Docker](https://docs.docker.com/engine/install/)
* [Docker Compose](https://docs.docker.com/compose/install/)

You will need to create an `.env` file. This file will contain any environment variables needed for it to run locally. To make it easier for you, there is already a `.env.example` file with all the needed values you'll need.

You will need to set up a domain name locally for the app to resolve. To do so, open your [hosts file](https://en.wikipedia.org/wiki/Hosts_(file)) and add this value.

```apache
127.0.0.1   domane.localhost
127.0.0.1   grafana.domane.localhost     # Optional
127.0.0.1   prometheus.domane.localhost  # Optional
```

*Be careful:* If in the last step you've changed the `DOMAIN` variable in `.env` file, you should use that instead of `domane.localhost`.

Once everything is set, run:
```shell
$ docker-compose up -d --build --force-recreate --remove-orphans

// Alternative, for unix with `make` installed:
$ make serve-local
```