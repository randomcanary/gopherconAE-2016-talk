from tornado import httpserver
from tornado import gen
from tornado.ioloop import IOLoop
import tornado.web
import random
import json

class MainHandler(tornado.web.RequestHandler):
    def get(self):
        self.write({"RandomInt" : int(random.random()*100)})

class Application(tornado.web.Application):
    def __init__(self):
        handlers = [
            (r"/?", MainHandler)
        ]
        tornado.web.Application.__init__(self, handlers)

def main():
    app = Application()
    app.listen(8080)
    IOLoop.instance().start()

if __name__ == '__main__':
    main()