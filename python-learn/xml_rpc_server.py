from xmlrpc.server import SimpleXMLRPCServer

class calculate:
    def add(self, x, y):
        return x + y
    def multiple(self, x,y):
        return x * y
    def substract(self, x, y):
       return abs(x - y)
    def divide(self, x, y):
       return x / y


obj = calculate()
server = SimpleXMLRPCServer(('localhost',8088))

server.register_instance(obj)

print('Listening to 8088')

server.serve_forever()