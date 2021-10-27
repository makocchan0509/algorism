import sys
from typing import Optional

class MiniHeap(object):
    def __init__(self) -> None:
        self.heap = [-1 * sys.maxsize]
        self.current_size = 0
    
    def parent_index(self,index: int) -> int:
        # Parent 2//2 =1
        return index // 2
    
    def left_child_index(self,index: int) -> int:
         # Left child 2 * 2 =1
        return 2 * index

    def right_child_index(self,index: int) -> int:
         # Right child (2 * 2)+1 =5
        return (2 * index) +1
    
    def swap(self,index1:int,index2: int) -> None:
        self.heap[index1],self.heap[index2] = self.heap[index2],self.heap[index1]
    
    def heapify_up(self, index: int) -> None:
        while self.parent_index(index) > 0:
            if self.heap[index] < self.heap[self.parent_index(index)]:
                self.swap(index,self.parent_index(index))
            index = self.parent_index(index)
    
    def push(self,value) -> None:
        self.heap.append(value)
        self.current_size += 1
        self.heapify_up(self.current_size)
    
    def min_child_index(self,index:int) -> int:
        if self.right_child_index(index) > self.current_size:
            return self.left_child_index(index)
        else:
            if self.heap[self.left_child_index(index)] < self.heap[self.right_child_index(index)]:
                return self.left_child_index(index)
            else:
                return self.right_child_index(index)
    
    def heapify_down(self, index:int) -> None:
        while self.left_child_index(index) <= self.current_size:
            min_index = self.min_child_index(index)
            if self.heap[index] > self.heap[min_index]:
                self.swap(index,min_index)
            index = min_index
    
    def pop(self) -> Optional[int]:
        if len(self.heap) == 1:
            return
        
        root = self.heap[1]
        data = self.heap.pop()
        if len(self.heap) == 1:
            return root

        self.heap[1] = data
        self.current_size -= 1
        self.heapify_down(1)
        return root


if __name__ == '__main__':
    min_heap = MiniHeap()
    min_heap.push(5)
    min_heap.push(6)
    min_heap.push(2)
    min_heap.push(9)
    min_heap.push(13)
    min_heap.push(11)
    min_heap.push(1)
    print(min_heap.heap)
    print(min_heap.pop())
    print(min_heap.heap)