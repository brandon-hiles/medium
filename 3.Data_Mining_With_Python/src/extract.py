import requests
import xml.etree.ElementTree as ET

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

        # This can be improved through regex
        condition_1 = self.url.split("://")[0] == 'https' # Check url contains https://
        condition_2 = len(self.url.split("medium.com")) > 1 # Check url contains medium.com
        condition_3 = len(self.url.split("/feed/")) > 1 # Check url contains /feed/
        condition_4 = len(self.url.split("/@")) > 1 # Check url contains /@

        if (condition_1 == False):
            return False
        elif (condition_2 == False):
            return False
        elif (condition_3 == False):
            return False
        elif (condition_4 == False):
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
        Main Function: Queries the medium RSS feed and returns a list of articles published.
        """

        articles = []
        
        if (self.validate_url_structure() == True and self.validate_account() == True):
            data = requests.get(self.url).content.decode("utf-8")
            root = ET.fromstring(data)
            for idx in range(0, len(root[0])):
                if (root[0][idx].tag == 'item'):
                    articles.append(root[0][idx])
        else:
            print("Please enter a valid medium account")

        return articles

    def parse_articles(self, tag):
        articles = self.grab_articles()
        info = []
        for article in articles:
            for idx in range(0, len(article)):
                if (article[idx].tag == tag):
                    info.append(article[idx].text)
        return info