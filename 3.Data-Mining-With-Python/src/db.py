from pymongo import MongoClient

class Database(object):

    def __init__(self, host, port, database):
        self.host = host
        self.port = port
        self.database = database
        self.client = MongoClient(self.host, self.port)

    def __repr__(self):
        return f"Database Object for {self.host}:{self.port}"
        
    def select_database(self, database):
        return self.client[self.database]

    def select_collection(self, collection):
        return self.select_database(self.database)[collection]

    def check_collection(self, collection, query):
        # Check if collection exists in db

        collection = self.select_collection(collection=collection)
        result = collection.find(query)
        if result.count() > 0:
            return True
        else:
            return False