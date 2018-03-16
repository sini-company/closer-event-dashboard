## CLOSER Event Dashboard

This project is an easy-to-install web dashboard which collects and visualizes events from CLOSER webhook.

### Components

* [Nginx](https://nginx.org/): HTTP and reverse proxy server
* [PostgreSQL](https://www.postgresql.org/): Open source database
* [Grafana](https://grafana.com/): Data visualization tool

### Prerequisites

* docker (https://docs.docker.com/engine/installation/)
* docker-compose (https://docs.docker.com/compose/install/)

### Usage

1.  Clone this repository:

```sh
$ git clone https://github.com/sinicompany/closer-event-dashboard.git
```

2.  Run docker-compose to run http server (`-d`: run in daemon mode)

```sh
$ docker-compose up [-d]
```

3.  Use via HTTP

    * GET `/`: Grafana dashboard
    * POST `/webhook-endpoint`: Webhook endpoint
