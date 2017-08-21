# micro-gateway

This project generates API gateway for [micro](https://github.com/micro/micro).
A lot of code is borrowed from grpc-gateway and micro/go-api.

Please note this project is just POC & very alpha quality now.

## Features

- You can(must) choose what service to be exposed by gateway.
- Generate swagger.json from `.proto` to generate clients.
- Run authentication middleware in both gateway and service are supported.

## How to use

See [example](./example).