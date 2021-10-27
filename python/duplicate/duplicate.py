
from typing import List

def remove_dep(numbers: List[int]) -> List[int]:
    result = []

    for num in numbers:
        if num not in result:
            result.append(num)
    return result

def remove_dep_v2(numbers: List[int]) -> List[int]:
    i = len(numbers)-1
    while i > 0:
        if numbers[i] == numbers[i-1]:
            numbers.pop(i)
        i -= 1
    return numbers


l = [1,3,3,5,5,7,7,7,10,12,12,15]
print(remove_dep(l))
print(remove_dep_v2(l))