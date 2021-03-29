import pika

credentials = pika.PlainCredentials("guest", "guest")
connection_params = pika.ConnectionParameters(
    "localhost",
    credentials=credentials,
)


def open_connection():
    return pika.BlockingConnection(connection_params)

