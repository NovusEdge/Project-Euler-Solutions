from functools import lru_cache
from time import time

@lru_cache
def fibonacci(n):
    if n in [0, 1]:
        return n
    return fibonacci(n - 1) + fibonacci(n - 2)



def main():
    i = answer = 0
    while (fibb_num := fibonacci(i)) < 4_000_000:
        if fibb_num % 2 == 0:
            answer += fibb_num
        i += 1

    return answer



if __name__ == '__main__':
    start = time()
    print(f"\nAnswer: { main() }")
    print(f"Time Taken: { time() - start }\n")
