import pika

credentials = pika.PlainCredentials("guest", "guest")
connection_params = pika.ConnectionParameters(
    "localhost",
    credentials=credentials,
)

broker_connection = pika.BlockingConnection(connection_params)

channel = broker_connection.channel()
print("Hello RabbitMQ")
