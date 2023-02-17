import json
from urllib.parse import urlparse, parse_qsl
from http.server import HTTPServer, BaseHTTPRequestHandler

host = ('', 8003)

class AddHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        parsed_url = urlparse(self.path)
        qs = dict(parse_qsl(parsed_url.query))
        a = int(qs.get("a", 0))
        b = int(qs.get("b", 0))
        self.send_response(200)
        self.send_header("content-type", "application/json")
        self.end_headers()
        self.wfile.send(json.load({
            'result': a + b
        }).encode('utf-8'))


if __name__ == '__main__':
    server = HTTPServer(host, AddHandler)
    server.serve_forever()
