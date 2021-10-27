from typing import List

def synmetory(input: List[int]) -> List[int]:
    dic ={}
    result = []
    for data in input:
        dic[data[0]]=data[1]
        if data[1] in dic and dic[data[1]] == data[0]:
            result.append((data[0],data[1]))
    return result




input = [(1,2),(3,5),(4,7),(5,3),(7,4)]
print(synmetory(input))

