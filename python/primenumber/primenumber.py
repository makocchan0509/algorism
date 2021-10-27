import math
from typing import List,Generator


def prime_number(target:int) -> List[int]:
    output =[]
    for x in range(2,target+1):
        for y in range(2,x):
            if x % y == 0:
                break
        else:
            output.append(x)
    return output

def prime_number_v2(number:int) -> List[int]:
    output = []
    cache = {}
    for x in range(2,number+1):
        is_prime = cache.get(x)
        if is_prime is False:
            continue
        output.append(x)
        cache[x] = True
        for y in range(x*2,number+1,x):
            cache[y] = False
    return output

def prime_number_v3(number:int) -> Generator[int,None,None]:
    cache = {}
    for x in range(2,number+1):
        is_prime = cache.get(x)
        if is_prime is False:
            continue
        yield x
        cache[x] = True
        for y in range(x*2,number+1,x):
            cache[y] = False


def is_prime_v1(number:int) -> bool:
    if number <= 1:
        return False
    for i in range(2,number):
        if number % i == 0:
            return False
    return True

def is_prime_v2(number:int) -> bool:
    if number <= 1:
        return False
    for i in range(2,math.floor(math.sqrt(number)+1)):
        if number % i == 0:
            return False
    return True

def is_prime_v3(number:int) -> bool:
    if number <= 1:
        return False
    
    if number == 2:
        return True
    
    if number % 2 ==0:
        return False

    i = 3
    while i * i <= number:
        if number % i == 0:
            return False
        i += 2
    return True

def is_prime_v4(number:int) -> bool:
    if number <= 1:
        return False
    
    if number <= 3:
        return True
    
    if number % 2 ==0 or number % 3 == 0:
        return False

    i = 5
    while i * i <= number:
        if number % i == 0 or number % (i+2) ==0:
            return False
        i += 6
    return True

if __name__ == "__main__":

    print(prime_number(50))
    print(prime_number_v2(50))
    print([i for i in prime_number_v3(50)])
