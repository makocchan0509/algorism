from typing import Tuple
from collections import Counter
import operator

def counter(string : str) -> Tuple[str,int]:
    counter = {}
    for c in string:
        c = c.lower()
        if c == ' ':
            continue
        elif not c in counter:
            counter[c] = 1
        else:
            counter[c] += 1
    max_k = max(counter, key=counter.get)
            
    return (max_k,counter[max_k])

def count_charsv1(string :str) -> Tuple[str,int]:
    str = string.lower()
    l = []
    for char in str:
        if not char.isspace():
            l.append((char,str.count(char)))
    return max(l,key=operator.itemgetter(1))

def count_charsv2(string :str) -> Tuple[str,int]:
    str = string.lower()
    d = {}
    for char in str:
        if not char.isspace():
            d[char] = d.get(char,0) +1
    max_key = max(d,key=d.get)
    return max_key,d[max_key]

def count_charsv3(string :str) -> Tuple[str,int]:
    str = string.lower()
    d = Counter()
    for char in str:
        if not char.isspace():
            d[char] += 1
    print(d)
    max_key = max(d,key=d.get)
    return max_key,d[max_key]


string = "This is a pen. This is an apple. Applepen"
print(counter(string))
print(count_charsv1(string))
print(count_charsv2(string))
print(count_charsv3(string))