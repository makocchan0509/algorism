from __future__ import annotations
from typing import Any,Optional

class Node(object):
    def __init__(self,data: Any = None ,prev:Node = None ,next:Node= None):
        self.data = data
        self.next = next
        self.prev = prev

class DoubleLinkedList(object):
    def __init__(self,head=None) -> None:
        self.head = head

    def append(self,data) -> None:
        new_data = Node(data)
        if self.head is None:
            self.head = new_data
            return

        current_node = self.head
        while current_node.next:
            current_node = current_node.next
        new_data.prev = current_node
        current_node.next = new_data
    
    def insert(self,data) -> None:
        new_data = Node(data)
        if self.head is None:
            self.head = new_data
            return

        next_node = self.head
        new_data.next = next_node
        next_node.prev = new_data
        self.head = new_data
    
    def print(self) -> None:
        current_node = self.head
        while current_node:
            print(current_node.data)
            current_node = current_node.next
    
    def remove(self,data) -> None:
        current_node = self.head

        if current_node and current_node.data == data:
            if current_node.next is None:
                current_node = None
                self.head = None
                return
            else:
                next_node = current_node.next
                next_node.prev = None
                current_node = None
                self.head = next_node
                return
        
        while current_node and current_node.data != data:
            current_node = current_node.next

        if current_node is None:
            return
        
        if current_node.next is None:
            prev_node = current_node.prev
            prev_node.next = None
            current_node = None
        else:
            next_node = current_node.next
            prev_node = current_node.prev
            next_node.prev = prev_node
            prev_node.next = next_node
            current_node = None
    
    def reverse_iterative(self) -> None:
        previous_node = None
        current_node = self.head

        while current_node:
            previous_node = current_node.prev
            current_node.prev = current_node.next
            current_node.next = previous_node
            current_node = current_node.prev
        
        if previous_node: 
            self.head = previous_node.prev

    def reverse_recursive(self) -> None:
        def _reverse_recursive(current_node: Node) -> Optional[Node]:
            if not current_node:
                return None
            previous_node = current_node.prev
            current_node.prev = current_node.next
            current_node.next = previous_node

            if current_node.prev is None:
                return current_node
            return _reverse_recursive(current_node.prev)
        self.head = _reverse_recursive(self.head)
    
    def sort(self) -> None:
        if self.head is None:
            return
        current_node = self.head
        while current_node.next:
            next_node = current_node.next
            while next_node:
                if current_node.data > next_node.data:
                    current_node.data,next_node.data = next_node.data,current_node.data
                next_node = next_node.next
            current_node = current_node.next


if __name__ == '__main__':
    l = DoubleLinkedList()
    l.append(2)
    l.append(4)
    l.append(6)
    l.append(0)
    l.print()
    l.sort()
    l.print()


    