from pymongo import MongoClient

class Database(object):

    def __init__(self, host, port, database):
        self.host = host
        self.port = port
        self.database = database
        self.client = MongoClient(self.host, self.port)

    def __repr__(self):
        return f"Database Object for {self.host}:{self.port}"