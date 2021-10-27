from typing import List

def remove_zero(numbers :List[int]):
    numbers.pop(0)
    remove_zero(numbers)

def list_to_int(numbers :List[int]) -> int:
    sum_numbers = 0
    for i,num in enumerate(reversed(numbers)):
        sum_numbers += num * (10**i)
    return sum_numbers

def list_to_int_plus(numbers :List[int]) -> int:
    i = len(numbers) -1
    numbers[i] += 1
    while 0 < i:
        if numbers[i] != 10:
            remove_zero(numbers)
            break
        numbers[i] = 0
        numbers[i-1] += 1
        i -= 1
    else:
        if numbers[0] == 10:
            numbers[0] = 1
            numbers.append(0)
    return list_to_int(numbers)

def list_sum(nums: List[int])-> None:
    result = []
    len_nums = len(nums)
    incre = True

    for i in range(len_nums-1,-1,-1):
        if incre:
            nums[i] += 1
            incre = False
            if nums[i] - 10 == 0:
                incre = True
                nums[i] = 0
            result.insert(0,nums[i])
        elif nums[i] != 0:
            result.insert(0,nums[i])
    if incre:
        result.insert(0,1)
    print(result)

    digit = 1
    sum = 0
    for n in range(len(result)-1,-1,-1):
        sum = sum + (result[n] * digit)
        digit *= 10
    print(sum)
    
#nums = [1]
#nums = [2,3]
#nums = [9,9]
#nums = [7,8,9]
nums = [9,9,9]
#nums = [0,0,0,9,9,9,9]

print(list_to_int_plus(nums))