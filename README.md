## CLOSER Event Dashboard

This project is an easy-to-install web dashboard which collects and visualizes events from CLOSER webhook.

![Screenshot](/screenshot.png?raw=true 'Screenshot')

---

### Components

- [Nginx](https://nginx.org/): HTTP and reverse proxy server
- [PostgreSQL](https://www.postgresql.org/): Relational database that supports arbitary JSON column
- [Grafana](https://grafana.com/): Data visualization tool

### Prerequisites

- docker (https://docs.docker.com/installation/)
- docker-compose (https://docs.docker.com/compose/install/)

### Usage

1.  Clone this repository:

    ```sh
    $ git clone https://github.com/sinicompany/closer-event-dashboard.git
    ```

2.  Run docker-compose to run http server (`-d`: run in daemon mode)

    ```sh
    $ docker-compose up [-d]
    ```

3.  Access dashboard via HTTP endpoints

    - GET `/`: Grafana dashboard (※ Initial credential: `admin` / `admin` )
    - POST `/webhook-endpoint`: Webhook endpoint
    
4.  Register your webhook endpoint to the CLOSER Bot Settings > Webhook

    ![Instruction](/instruction.png?raw=true 'Instruction')
