import requests
import bs4

from src.db import Database

class MediumParser(object):

    def __init__(self, data):
        self.data = data

    def __repr__(self):
        return f"Parser Object for data: {self.data}"

    def grab_elements(self, website):
        # This extracts the Tag class from html page
        # Note: This needs an html page, NOT a url.

        soup = bs4.BeautifulSoup(website, 'html.parser')
        children = list(soup.children)
        element = [children[num] for num, value in enumerate(children)
        if type(children[num]) is bs4.element.Tag]
        return element

    def grab_text(self):

        texts = []
        for website in self.data:
            data = requests.get(website).content.decode('utf-8')
            elms = self.grab_elements(website=data)
            text = elms[0].find_all("p")
            text = [''.join(text[num].text) for num, val in enumerate(text)]
            texts.append("".join(text))
        return texts

    def store_information(self):
        pass