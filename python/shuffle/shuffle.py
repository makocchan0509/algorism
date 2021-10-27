from typing import List

def shuffle(numbers: List[int]) -> List[int]:
    i,j = 0,len(numbers)-1

    while i < j:
        if numbers[i] % 2 ==0:
            i += 1
        else:
            numbers[i],numbers[j] = numbers[j],numbers[i]
            j -= 1
    return numbers

def sort_index(numbers:List[int],chars:List[str]) -> str:
    tmp = [None] * len(chars)
    for i,index in enumerate(numbers):
        tmp[index] = chars[i]
    return "".join(tmp)

def sort_index_v2(numbers:List[int],chars:List[str]) -> str:
    i, len_indexes = 0,len(numbers)-1
    while i < len_indexes:
        while i != numbers[i]:
            index = numbers[i]
            chars[index],chars[i] = chars[i],chars[index]
            numbers[index],numbers[i] = numbers[i],numbers[index]
        i += 1
    return "".join(chars)


l = [0, 1, 3, 4, 2, 4, 5, 1, 6, 9, 8]
print(shuffle(l))
nums=[3,1,5,0,2,4]
chars=["h","y","n","p","t","o"]
print(sort_index_v2(nums,chars))
