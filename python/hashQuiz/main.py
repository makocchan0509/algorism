import time

cache={}

def long_func(num:int) -> int:
    r = 0
    if num in cache:
        return cache[num]
    for i in range(10000000):
        r += num * i
    cache[num] = r
    return r


def memoize(f):
    cache = {}
    def _wrapper(n):
        if n not in cache:
            cache[n] = f(n)
        return cache[n]
    return _wrapper

@memoize
def long_func_t(num:int) -> int:
    r = 0
    for i in range(10000000):
        r += num * i
    return r


for i in range(10):
    print(long_func_t(i))
start = time.time()
for i in range(10):
    print(long_func_t(i))
print(time.time() - start)