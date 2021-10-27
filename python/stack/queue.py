from collections import deque
from typing import Any

class Queue(object):

    def __init__(self) -> None:
        self.queue = []
    
    def enqueue(self,data:Any) -> None:
        self.queue.append(data)
    
    def dequeue(self) -> Any:
        if self.queue:
            return self.queue.pop(0)


def reverse(queue):
    new_queue = deque()
    while queue:
        new_queue.append(queue.pop())
    return new_queue

d = deque()
d.append(1)
d.append(2)
d.append(3)
d.append(4)
d.reverse()
print(d)
