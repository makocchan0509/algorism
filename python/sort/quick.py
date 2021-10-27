from typing import List

def partition(numbers: List[int],low: int,high: int) -> int:
    # i=-1 j=0からスタートする
    i = low -1
    pivot = numbers[high]
    for j in range(low,high):
        if numbers[j] <= pivot:
            i +=1
            numbers[j],numbers[i] = numbers[i],numbers[j]
    numbers[i+1],numbers[high] = numbers[high],numbers[i+1]
    return i +1

def quick_sort(numbers: List[int]) -> List[int]:
    def _quick_sort(numbers:List[int], low: int,high: int) -> int:
        if low < high:
            partition_index = partition(numbers,low,high)
            _quick_sort(numbers,partition_index+1,high)
            _quick_sort(numbers,low,partition_index-1)
    
    _quick_sort(numbers,0,len(numbers)-1)
    return numbers

if __name__ == "__main__":
    nums = [1,8,3,9,4,5,7]
    print(quick_sort(nums))