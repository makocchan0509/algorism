from typing import List

def gnome_sort(numbers :List[int]) -> List[int]:
    len_numbers = len(numbers)
    for i in range(len_numbers-1):
#        print("i",i)
        if numbers[i] > numbers[i +1]:
            numbers[i],numbers[i+1] = numbers[i+1],numbers[i]
            for j in range(i-1 ,-1,-1):
#                print("j:",j)
                if numbers[j] > numbers[j+1]:
                    numbers[j],numbers[j +1] = numbers[j+1],numbers[j]
                else:
                    break
#        print(numbers)
    return numbers

def gnome_sort_t(numbers):
    len_numbers = len(numbers)
    index = 0
    while index < len_numbers:
        if index == 0:
            index += 1
        if numbers[index] >= numbers[index-1]:
            index += 1
        else:
            numbers[index],numbers[index-1] = numbers[index-1],numbers[index]
            index -= 1
    return numbers

if __name__ == "__main__":
    nums = [4,6,7,3,1,2,4,8]
    import random
    nums = [random.randint(0,1000) for _ in range(10)]
    print(gnome_sort_t(nums))