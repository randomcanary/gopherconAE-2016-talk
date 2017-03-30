from tornado import gen
from tornado.ioloop import IOLoop
import tornado.web
import tornado_mysql
import json
import random

class MainHandler(tornado.web.RequestHandler):
    
    @gen.coroutine
    def get(self):
        conn = yield tornado_mysql.connect(host='127.0.0.1', port=3306, user='root', passwd='', db='world')
        cur = conn.cursor()
        yield cur.execute("SELECT Name FROM City")
        list_result , result = [], []
        for row in cur:
            result.append( {"CityName" : str(row[0])} )
            list_result.append(str(row[0]).decode('Latin1').encode('utf-8')) 
            
        list_result.sort()
        self.write({"Names" : list_result})
        cur.close()
        conn.close()

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