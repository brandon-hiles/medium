import requests
import xml.etree.ElementTree as ET
import re

class MediumExtractor(object):

    def __init__(self, url):
        self.url = url

    def __repr__(self):
        return f"Extractor Object for {self.url}"

    def validate_url_structure(self):
        """
        Helper Function: Verifies that URL given to crawler is of
        the correct format:

        https://medium.com/feed/@username
        """

        url_structure = "https://medium.com/feed/"
        result = re.search(url_structure, self.url)
        if result == None:
            return False
        else:
            return True        

    def validate_account(self):
        """
        Helper Function: Verfiies that account exists within the medium database
        """
        if (requests.get(self.url).status_code == 200):
            return True
        elif (requests.get(self.url).status_code == 404):
            return False

    def grab_articles(self):
        """
        Query Function: Queries the medium RSS feed and returns a list of articles published.
        """

        articles = []
        
        if (self.validate_url_structure() == True and self.validate_account() == True):
            data = requests.get(self.url).content.decode("utf-8")
            root = ET.fromstring(data)
            for idx in range(0, len(root[0])):
                if (root[0][idx].tag == 'item'):
                    articles.append(root[0][idx])
        else:
            print("Please enter a valid medium url")

        return articles

    def parse_articles(self, tag):
        """
        Parser Function: Used to parse the data from the grab_articles method
        """
        articles = self.grab_articles()
        info = []
        for article in articles:
            for idx in range(0, len(article)):
                if (article[idx].tag == tag):
                    info.append(article[idx].text)
        return info
