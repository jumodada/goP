import asyncio

from grpclib.client import Channel

# generated by protoc
from grpclib_test.helloworld_pb2 import HelloRequest, HelloReply
from grpclib_test.helloworld_grpc import GreeterStub

#客户端可以使用asyncio发起连接
#服务端使用了asyncio 可以维护大量的链接
#1. 这个会影响go和python之间的互相调用吗

async def main():
    async with Channel('127.0.0.1', 50051) as channel:
        greeter = GreeterStub(channel)

        reply = await greeter.SayHello(HelloRequest(name='Dr. Strange'))
        print(reply.message)


if __name__ == '__main__':
    asyncio.run(main())