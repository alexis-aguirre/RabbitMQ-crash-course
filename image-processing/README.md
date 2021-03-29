# Processing Service

## Introduction

Welcome to the processing service 🥶

## Before the workshop

Make sure to:
- Run with python version >= 3.9
- Install the dependencies, some alternatives
    - With poetry
        - Run `poetry install`
    - Or manually
        - `pip install pika`

That's all, see you in the workshop


## Start the server

To start the service at port 8082 

    python server.py

## Endpoint

The available endpoint is `http://localhost:8082/process`

```sh
curl --request POST \
  --url http://localhost:8082/process \
  --header 'Content-Type: application/json' \
  --data '{
    "id": 20,
    "location": "dd",
    "image": "https://github.com/alexis-aguirre/RabbitMQ-crash-course/tree/skeleton/data/license-plates/MX-10-696.jpg"
}'
```
