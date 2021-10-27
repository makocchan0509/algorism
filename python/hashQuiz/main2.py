from typing import List
from collections import Counter

def quiz(x:List[int],y:List[int]):
    countX = {}
    countY = {}
    resultX = []
    resultY = []

    for n in x:
        countX[n] = x.count(n)
        resultX.append(n)
    for m in y:
        countY[m] = y.count(m)
        resultY.append(m)

    for data in x:
        if data in x and data in y and countX[data] < countY[data]:
            while data in resultX:
                resultX.remove(data)
        elif data in x and data in y and countX[data] > countY[data]:
            while data in resultY:
                resultY.remove(data)
    return resultX,resultY

def quiz_t(x:List[int],y:List[int]) -> None:

    counter_x = Counter(x)
    counter_y = Counter(y)

    for key_x,value_x in counter_x.items():
        value_y = counter_y.get(key_x)
        if value_y:
            if value_x < value_y:
                x[:] = [i for i in x if i != key_x]
            elif value_x > value_y:
                y[:] = [i for i in y if i != key_x]

inputX = [1,2,3,4,4,5,5,8,10]
inputY = [4,5,5,5,6,7,8,8,10]
#print(quiz(inputX,inputY))
quiz_t(inputX,inputY) 
print(inputX)
print(inputY)
