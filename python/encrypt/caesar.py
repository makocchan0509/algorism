import string
from typing import Generator, Tuple

def caesar_cipher(text: str,shift:int) -> str:
    result = ""
    print(len(string.ascii_uppercase))
    len_alphabet = ord('Z')-ord('A')+1

    for char in text:
        if char.isupper():
            alphabet + string.ascii_uppercase
        elif char.islower():
            alphabet = string.ascii_lowercase
        else:
            result += char
            continue

        index = (alphabet.index(char)+shift) % len_alphabet
        result += alphabet[index]
    return result


if __name__ == "__main__":
    alpha = string.ascii_uppercase
    print(alpha)
    print((alpha.index("Z")+3) % len(alpha))
    print(ord('Z'))
    print(ord('A'))