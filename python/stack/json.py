

def valid_json(chars: str) -> bool:
    lookup = {'{':'}', '[':']','(':')'}
    stack = []
    for char in chars:
        if char in lookup.keys():
            stack.append(lookup[char])
        if char in lookup.values():
            if not stack:
                return False
            if char != stack.pop():
                return False
    if stack:
        return False
    return True