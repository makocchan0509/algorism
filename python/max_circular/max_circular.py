from typing import List

def max_min_sequence_num(numbers:List[int],operator = max) -> int:
    sum_seq,result_seq = 0,0
    for num in numbers:
        sum_seq = operator(num,sum_seq + num)
        result_seq = operator(sum_seq,result_seq)
    return result_seq

def find_max_circular_sequence_sum(numbers:List[int]) -> int:
    max_sequence_sum = max_min_sequence_num(numbers)
    max_wrap_seq = sum(numbers) - max_min_sequence_num(numbers,operator=min)
    return max(max_sequence_sum,max_wrap_seq)

if __name__ == '__main__':
    l = [1, -2, 3, 6, -1, 2, 4, -5, 2]
    #print(max_min_sequence_num(l))
    print(find_max_circular_sequence_sum(l))