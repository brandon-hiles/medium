import argparse
import time

from src.extract import MediumExtractor
from src.parser import MediumParser

username = "brandon.j.hiles"
url = "https://medium.com/feed/@" + username

host = 'localhost'
port = 27017

parser = argparse.ArgumentParser()
parser.add_argument("-m", "--medium", help="Gather information from medium",
                    action="store_true")

if __name__ == '__main__':

    # Analyze arguments passed to python script
    args = parser.parse_args()

    start_time = time.time() # Start time for crawler
    if args.medium:
        print(f"Gathering information on {username} from Medium")
        extractor = MediumExtractor(url=url)
        articles = extractor.parse_articles(tag="link")
        parser = MediumParser(username=username, data=articles, host=host, port=port, database="medium")
        parser.store_information()


    print(f"Time Execution: {time.time() - start_time} s")