from typing import List

NUM_ALPHABET_MAPPING = {
    0: '+',
    1: '@',
    2: 'ABC',
    3: 'DEF',
    4: 'GHI',
    5: 'JKL',
    6: 'MNO',
    7: 'PQRS',
    8: 'TUV',
    9: 'WXYZ',
}

def phone_mnemonic_v1(numbers:str) -> List[str]:
    phone_numbers = [int(s) for s in numbers.replace("-","")]
    candicate = []
    tmp = [""] * len(phone_numbers)

    def find_candicate_alphabet(digit:int=0):
        if digit == len(phone_numbers):
            candicate.append("".join(tmp))
        else:
            for char in NUM_ALPHABET_MAPPING[phone_numbers[digit]]:
                tmp[digit] = char
                find_candicate_alphabet(digit+1)
    find_candicate_alphabet()
    return candicate

def phone_mnemonic_v2(numbers:str) -> List[str]:
    phone_number = [int(s) for s in numbers.replace("-","")]
    candicate = []
    stack = ['']

    while len(stack) != 0:
        alphabet = stack.pop()
        if len(alphabet) == len(phone_number):
            candicate.append(alphabet)
        else:
            for char in NUM_ALPHABET_MAPPING[phone_number[len(alphabet)]]:
                stack.append(alphabet+char)
    return candicate

print(phone_mnemonic_v1("23"))
print(phone_mnemonic_v2("23"))
