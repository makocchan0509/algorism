from typing import List


def select_sort(numbers:List[int]) -> List[int]:
    len_numbers = len(numbers)
    start = 0
    index = 0

    while start < len_numbers:
        min = numbers[start]
        swapped = False
        for i in range(start,len_numbers):
            if numbers[i] < min:
                min = numbers[i]
                index = i
                swapped = True
        if swapped:
            numbers[start],numbers[index] = numbers[index],numbers[start]
        start = start + 1
    return numbers

def select_sort_t(numbers :List[int]) -> List[int]:
    len_numbers = len(numbers)
    for i in range(len_numbers):
        min_idx = i
        for j in range(i+1,len_numbers):
            if numbers[min_idx] > numbers[j]:
                min_idx = j
        numbers[i],numbers[min_idx] = numbers[min_idx],numbers[i]
    return numbers

if __name__ == '__main__':
    import random

    nums = [random.randint(0,1000) for i in range(10)]
#    nums = [5,2,1,6,7,8,3,4]
    print(select_sort_t(nums))



