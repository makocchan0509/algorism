from typing import Generator

def palindrome(chars:str) -> bool:
    new_str = ''.join(list(reversed(chars)))
    if new_str == chars:
        return True
    return False

def palindrome_v2(chars:str) -> bool:
    return chars == chars[::-1]

def find_palindrome(strings: str, left:int,right:int) -> Generator[str,None,None]:
    while 0 <= left and right <= len(strings) -1:
        if strings[left] != strings[right]:
            break
        yield strings[left: right+1]
        left -= 1
        right += 1

def find_all_palindrome(strings:str) -> Generator[str,None,None]:
    len_str = len(strings)
    if not len_str:
        yield
    if len_str == 1:
        yield strings
    for i in range(1,len_str-1):
        yield from find_palindrome(strings,i-1,i+1)
        yield from find_palindrome(strings,i-1,i)

if __name__ == '__main__':
    for s in find_all_palindrome('fdafiewaafcabacdfafdaf'):
        print(s)

