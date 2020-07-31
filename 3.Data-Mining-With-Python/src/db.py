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
        """ 
        Helper Function: Selects the database instance
        """
        return self.client[self.database]

    def select_collection(self, collection):
        """
        Helper Function: Selects the collection instance form mongo
        """
        return self.select_database(self.database)[collection]

    def check_collection(self, collection, query):
        """
        Boolean Function: Test whether the given query already exists within the
        database to avoid duplication of data
        """

        collection = self.select_collection(collection=collection)
        result = collection.find(query)
        if result.count() > 0:
            return True
        else:
            return False