from typing import List,Iterator

def listExcludeIndices(data, indices=[]):
    return [x for i,x in enumerate(data) if i not in indices]

def permutation(numbers: List[int]) -> None:
    for i in range(len(numbers)):
        for j in range(len(numbers)-1):
            for k in range(len(numbers)-2):
                jData = listExcludeIndices(numbers,[i])
                kData = listExcludeIndices(jData,[j])
                print([numbers[i],jData[j],kData[k]])

def all_perms(elements: List[int]) -> Iterator[List[int]]:
    print("start",elements)
    if len(elements) <= 1:
        yield elements
    else:
        for perm in all_perms(elements[1:]):
            print("m",elements[1:])
            print("perm",perm)
            print("len",len(elements))
            for i in range(len(elements)):
                print("i",i)
                print("1",perm[:i])
                print("2",elements[0:1])
                print("3",perm[i:])
                yield perm[:i] + elements[0:1] + perm[i:]
            print("-")

l = [1,2,3]
#permutation(l)
for p in all_perms(l):
    print(p)