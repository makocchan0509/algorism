from typing import List

def counting_sort(numbers: List[int]) -> List[int]:
    max_number = max(numbers)
    counts = [0] * (max_number +1)
    results = [0] * len(numbers)

    for num in numbers:
        counts[num] += 1
    
    for i in range(1, len(counts)):
        counts[i] += counts[i-1]

    i = len(numbers) -1
    while i >= 0:
        index = numbers[i]
        results[counts[index]-1] = numbers[i]
        counts[index] -= 1
        i -= 1

    return results

if __name__ == '__main__':
    nums = [4,3,6,2,3,4,7]
    import random
    nums = [random.randint(0,1000) for i in range(10)]
    print(counting_sort(nums))