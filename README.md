# docker 4 dev - Container Day 2021 


https://2021.containerday.it/

Source code for the Container Day 2021 

## node js sample

In order to run the project:

```sh
cd js-fastify
docker compose up -d
```
Open the browser to: http://localhost:4050/

To check the logs:

```sh
docker compose logs -f
```

## go sample

This sample in order to run, requires to build the dev go image with the required tools:

```sh
docker build  -t golang-compile-daemon:1.17-alpine ./go-server/_setup
```

Once built the go dev image

```sh
cd go-server
docker compose up -d
```

Open the browser to: http://localhost:4060/

To check the logs:

```sh
docker compose logs -f
```


## clean up

In order to clean up the docker containers:

```sh
docker compose down
```

