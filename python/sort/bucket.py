from typing import List

def bucket_sort(numbers :List[int]) -> List[int]:
    len_number = len(numbers)
    max_num = max(numbers)
    size = max_num // len_number
    buckets = [[] for _ in range(size)]

    for num in numbers:
        i = num // size
        if i != size:
            buckets[i].append(num)
        else:
            buckets[size-1].append(num)
    
    for i in range(size):
        insertion_sort_t(buckets[i])
    
    result = []
    for i in range(size):
        result += buckets[i]
    
    return result

def insertion_sort_t(numbers :List[int]) -> List[int]:
    len_number = len(numbers)
    for i in range(1,len_number):
        temp = numbers[i]
        j = i - 1
        while j > 0 and numbers[j] > temp:
            numbers[j+1] = numbers[j]
            j -= 1
        numbers[j] = temp
    return numbers

if __name__ == '__main__':
    nums = [5,3,4,2,6,8,7,1]
    import random
    nums =[random.randint(0,1000) for _ in range(10)]
    print(bucket_sort(nums))