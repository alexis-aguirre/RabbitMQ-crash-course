from json import dumps
from typing import Dict
from app.config import open_connection


class PublishManager:
    def __init__(self, exchange: str, routing_key: str):
        self.__exchange = exchange
        self.__routing_key = routing_key
        self.__connection = open_connection()
        self.__channel = self.__connection.channel()

    def publish(self, payload: Dict[str, str]):
        self.__declare_exchange('topic')
        self.__channel.basic_publish(exchange=self.__exchange,
                                     routing_key=self.__routing_key,
                                     body=dumps(payload))

    def __declare_exchange(self, exchange_type: str):
        self.__channel.exchange_declare(
            exchange=self.__exchange, exchange_type=exchange_type, durable=True)

    def close_connection(self):
        self.__connection.close()
