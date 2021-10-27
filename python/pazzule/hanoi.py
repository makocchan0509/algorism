

def hanoi(disk:int,src:str,dest:str,support:str):
    if disk < 1:
        return
    
    hanoi(disk-1,src,support,dest)
    print(f'move {disk} from {src} to {dest}')
    hanoi(disk-1,support,dest,src)

if __name__ == '__main__':
    hanoi(3,'A','C','B')
