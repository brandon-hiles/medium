import argparse
import time

parser = argparse.ArgumentParser()
parser.add_argument("-m", "--medium", help="Gather information from medium",
                    action="store_true")

if __name__ == '__main__':

    # Analyze arguments passed to python script
    args = parser.parse_args()

    start_time = time.time() # Start time for crawler
    if args.medium:
        print("Gathering information on Brandon Hiles from Medium")
        # Do stuff here

    print(f"Time Execution: {time.time() - start_time} s")