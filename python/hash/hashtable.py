import hashlib
from typing import Any

class HashTable(object):

    def __init__(self,size=10) -> None:
        self.size = size
        self.table = [[] for _ in range(self.size)]

    def hash(self,key) -> int:
        return int(hashlib.md5(key.encode()).hexdigest(),base=16) % self.size
    
    def add(self,key,value) -> None:
        index = self.hash(key)
        for data in self.table[index]:
            if data[0] == key:
                data[1] = value
                break
        else:
            self.table[index].append([key,value])
    
    def print(self) -> None:
        for index in range(self.size):
            print(index, end=' ')
            for data in self.table[index]:
                print("->",end=' ')
                print(data,end=' ')
            print()
    def get(self,key):
        index = self.hash(key)
        for data in self.table[index]:
            if data[0] == key:
                return data[1]
    
    def __setitem__(self,key,value):
        self.add(key,value)
    
    def __getitem__(self,key):
        return self.get(key)

    
if __name__ == "__main__":
    m = HashTable()
    m.add("car","toyota")
    m.add("sns","youtube")
    m.add("car","nissan")
    m.print()
    print(m.get('car'))
    m['language'] = "eng"
    print(m['language'])

