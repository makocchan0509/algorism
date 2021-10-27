class Node(object):
    def __init__(self, value: int) -> None:
        self.value = value
        self.left = None
        self.right = None

def insert(node: Node,value : int) -> Node:
    if node is None:
        return Node(value)

    if value < node.value:
        node.left = insert(node.left,value)
    if value > node.value:
        node.right = insert(node.right,value)
    return node

def inorder(node: Node) -> None:
    if node is not None:
        inorder(node.left)
        print(node.value)
        inorder(node.right)

def search(node:Node, value:int) -> bool:
    if node is None:
        return False
    
    if node.value == value:
        return True
    elif node.value < value:
        result = search(node.right,value)
    else:
        result = search(node.left,value)
    return result

def min_value(node: Node) -> Node:
    current = node
    while current.left:
        current = current.left
    return current

def remove(node: Node,value:int) -> Node:
    if node is None:
        return node
    
    if node.value < value:
        node.right = remove(node.right,value)
    elif node.value > value:
        node.left = remove(node.left,value)
    else:
        if node.left is None:
            return node.right
        elif node.right is None:
            return node.left
        
        temp = min_value(node.right)
        node.value = temp.value
        node.right = remove(node.right,temp.value)
    return node


if __name__ == "__main__":
    root = None
    root = insert(root,3)
    root = insert(root,1)
    root = insert(root,6)
    root = insert(root,2)
    root = insert(root,5)
    root = insert(root,7)
    root = insert(root,10)

    inorder(root)
    print(search(root,3))
    print(search(root,10))
    print(search(root,11))
    root = remove(root,6)
    inorder(root)

    