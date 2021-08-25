from time import time


main = lambda: sum([i for i in range(1000) if (i % 3 == 0 or i % 5 == 0)])


if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
