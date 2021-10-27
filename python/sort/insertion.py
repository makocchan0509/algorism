from typing import List

def insertion_sort(numbers: List[int]) -> List[int]:
    len_number = len(numbers)

    for i in range(len_number-1):
#        print("i>--",numbers)
        if numbers[i] > numbers[i+1]:
            numbers[i],numbers[i+1] = numbers[i+1],numbers[i]
            if i != 0:
                for j in range(i,0,-1):
                    if numbers[j] < numbers[j-1]:
                        numbers[j],numbers[j-1] = numbers[j-1],numbers[j]
                    else:
                        break
#                print("j>--",numbers)
    return numbers

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
    print(insertion_sort_t(nums))
