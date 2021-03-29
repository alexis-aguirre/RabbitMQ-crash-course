from json import dumps, loads, decoder
from http.server import HTTPServer, BaseHTTPRequestHandler
from http import HTTPStatus
from app import PublishManager


def process_handler(handler: BaseHTTPRequestHandler):
    if handler.headers['Content-Type'] != 'application/json':
        bad_request(handler)
        return

    content_length = int(handler.headers['Content-Length'])
    raw_data = handler.rfile.read(content_length)
    try:
        post_data = loads(raw_data)
    except decoder.JSONDecodeError:
        bad_request(handler)
        return

    required_keys = ["id", "location", "image"]
    if not all(k in post_data for k in required_keys):
        bad_request(handler)
        return

    manager = PublishManager(exchange='my-exchange', routing_key='main-queue')
    manager.publish({
        "id": post_data['id'],
        "location": post_data['location'],
        "image": post_data['image'],
    })
    manager.close_connection()

    success_response(handler)
    handler.wfile.write(bytes(dumps({}), "utf-8"))


def success_response(handler: BaseHTTPRequestHandler):
    handler.send_response(HTTPStatus.OK)
    handler.send_header('Content-type', 'application/json')
    handler.end_headers()


def bad_request(handler: BaseHTTPRequestHandler):
    handler.send_response(HTTPStatus.BAD_REQUEST)
    handler.end_headers()


def not_found(handler: BaseHTTPRequestHandler):
    handler.send_response(HTTPStatus.NOT_FOUND)
    handler.end_headers()


class Handler(BaseHTTPRequestHandler):
    def do_POST(self):
        if self.path == "/webhook":
            process_handler(self)
        else:

            not_found(self)


def run(server_class=HTTPServer, handler_class=Handler):
    server_address = ('', 8080)
    httpd = server_class(server_address, handler_class)
    httpd.serve_forever()


if __name__ == '__main__':
    run()
