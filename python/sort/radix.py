
from typing import List

def counting_sort(numbers: List[int], place: int) -> List[int]:
    counts = [0] * 10
    results = [0] * len(numbers)

    for num in numbers:
        index = int(num / place) % 10
        counts[index] += 1
    
    for i in range(1, len(counts)):
        counts[i] += counts[i-1]

    i = len(numbers) -1
    while i >= 0:
        index = int(numbers[i]/place) % 10
        results[counts[index]-1] = numbers[i]
        counts[index] -= 1
        i -= 1

    return results

def radix_sort(numbers: List[int]) -> List[int]:
    max_num = max(numbers)
    place = 1
    while max_num > place:
        numbers = counting_sort(numbers,place)
        print(numbers)
        place *=10
    return numbers


if __name__ == '__main__':
    nums = [4,3,6,2,3,4,7]
    import random
    nums = [random.randint(0,1000) for i in range(10)]
    print(nums)
    print(radix_sort(nums))
