from time import time
from math import lcm
from functools import reduce


main = lambda: reduce(lcm, range(1, 21))


if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
