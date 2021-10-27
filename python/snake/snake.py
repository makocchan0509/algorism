from typing import List
import operator

def snake(numbers :List[any],count :int,rownum: int) -> None:
    snake = [[] for i in range(rownum)]
    index = -(-rownum // 2)-1
    print(index)
    slide = -1
    while count > 0:
        for num in numbers:
            for i in range(len(snake)):
                if i == index:
                    snake[i].append(num)
                else:
                    snake[i].append(" ")
            if index == 0 or index == rownum -1:
                slide *= -1
            index = index + slide
        count -= 1
    for data in snake:
      print("".join(str(i) for i in data))

def snake_string(chars: str,depth:int) -> List[List[str]]:
    result = [[] for _ in range(depth)]
    result_indexes = {i for i in range(depth)}
    insert_index = int(depth /2)

    op = operator.neg

    for s in chars:
        result[insert_index].append(s)
        for rest_index in result_indexes - {insert_index}:
            result[rest_index].append(" ")
        if insert_index <= 0:
            op = operator.pos
        if insert_index >= depth -1:
            op = operator.neg
        insert_index += op(1)
    
    return result

nums = [0,1,2,3,4,5,6,7,8,9]
#snake(nums,5,20)

import string
alphabet = [s for _ in range(2) for s in string.ascii_lowercase]
strings = "".join(alphabet)
for line in snake_string(strings,10):
    print("".join(line))