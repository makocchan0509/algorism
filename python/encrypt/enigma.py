import string
import random


ALPHABET = string.ascii_uppercase

class PlugBoard(object):
    def __init__(self,map_alphabet):
        self.alphabet = ALPHABET
        self.forward_map = {}
        self.backward_map = {}
        self.mapping(map_alphabet)
    
    def mapping(self,map_alphabet):
        self.forward_map = dict(zip(self.alphabet,map_alphabet))
        self.backward_map = {v:k for k,v in self.forward_map.items()}
    
    def forward(self,index):
        char = self.alphabet[index]
        char = self.forward_map[char]
        return self.alphabet.index(char)
    
    def backward(self,index):
        char = self.alphabet[index]
        char = self.backward_map[char]
        return self.alphabet.index(char)

class Rotor(PlugBoard):

    def __init__(self,map_alphabet,offset=0):
        super().__init__(map_alphabet)
        self.offset = offset
        self.rotations = 0
    
    def rotate(self,offset=None):
        if not offset:
            offset = self.offset
        self.alphabet = self.alphabet[offset:] + self.alphabet[:offset]
        self.rotations += offset
        return self.rotations
    
    def reset(self):
        self.rotations = 0
        self.alphabet = ALPHABET

class Reflector(object):

    def __init__(self,map_alphabet):
        self.map = dict(zip(ALPHABET,map_alphabet))
        for x,y in self.map.items():
            if x != self.map[y]:
                raise ValueError(x,y)
    
    def reflect(self,index_num):
        char = self.map[ALPHABET[index_num]]
        return ALPHABET.index(char)

class EnigmaMachine(object):

    def __init__(self,plug_board,rotors,reflector):
        self.plug_board = plug_board
        self.rotors = rotors
        self.reflector = reflector
    
    def encrypt(self,text):
        return ''.join([self.go_through(c) for c in list(text)])
    
    def decrypt(self,text):
        for rotor in self.rotors:
            rotor.reset()
        return ''.join([self.go_through(c) for c in list(text)])
    
    def go_through(self,char):
        char = char.upper()
        if char not in ALPHABET:
            return char

        index_num = ALPHABET.index(char)
        index_num=self.plug_board.forward(index_num)

        for rotor in self.rotors:
            index_num = rotor.forward(index_num)
        
        index_num = self.reflector.reflect(index_num)

        for rotor in reversed(self.rotors):
            index_num = rotor.backward(index_num)
        
        index_num = self.plug_board.backward(index_num)

        for rotor in reversed(self.rotors):
            if rotor.rotate() % len(ALPHABET) != 0:
                break
        
        return ALPHABET[index_num]
        

if __name__ == "__main__":
    # plugboard = PlugBoard("BADC")
    # encrypted = plugboard.forward(ALPHABET.index("A"))
    # print(ALPHABET[encrypted])
    # decrypted = ALPHABET[plugboard.backward(encrypted)]
    # print(decrypted)

    # rotor = Rotor("BADC",1)
    # rotor.rotate()

    # encrypted = rotor.forward(ALPHABET.index("A"))
    # print(ALPHABET[encrypted])
    # decrypted = ALPHABET[rotor.backward(encrypted)]
    # print(decrypted)

    # r = Reflector("BADC")
    # print(ALPHABET[r.reflect(ALPHABET.index("A"))])


    get_random_alphabet = lambda : ''.join(
        random.sample(ALPHABET,len(ALPHABET))
    )
    print(get_random_alphabet())

    p = PlugBoard(get_random_alphabet())
    r1 = Rotor(get_random_alphabet(),3)
    r2 = Rotor(get_random_alphabet(),2)
    r3 = Rotor(get_random_alphabet(),1)

    r = list(ALPHABET)
    indexes = [i for i in range(len(ALPHABET))]
    for _ in range(int(len(indexes)/2)):
        x = indexes.pop(random.randint(0,len(indexes)-1))
        y = indexes.pop(random.randint(0,len(indexes)-1))
        r[x],r[y] = r[y],r[x]
    print(''.join(r))
    reflector = Reflector(''.join(r))
    enigma = EnigmaMachine(p,[r1,r2,r3],reflector)

    e = enigma.encrypt("ATTACK SILICON VALLEY")
    print(e)
    print(enigma.decrypt(e))
