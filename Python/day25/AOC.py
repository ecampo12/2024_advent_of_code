from itertools import product

def parse_input(input):
    lock_schematics = []
    key_schematics = []
    
    for s in input.split("\n\n"):
        if s[0] == '.':
            key_schematics.append(s)
        else:
            lock_schematics.append(s)
    locks = []
    keys = []
    
    for s in lock_schematics:
        lock = []
        transposed = [list(x) for x in zip(*s.split("\n"))] # learned this last year, tanspose rows and columns
        for line in transposed:
            lock.append(''.join(line).rindex('#'))
        locks.append(lock)
    
    for s in key_schematics:
        key = []
        transposed = [list(x)[::-1] for x in zip(*s.split("\n"))] # we reverse the list rotate counter clockwise
        for line in transposed:
            key.append(''.join(line).rindex('#'))
        keys.append(key)
    return locks, keys

def part1(input):
    return sum(1 for key, lock in product(*parse_input(input)) if all(k + l < 6 for k, l in zip(key, lock)))

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")

if __name__ == "__main__":
    main()