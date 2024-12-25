from itertools import product

def parse_input(input):
    locks, keys = [], []
    for s in input.split("\n\n"):
        if s[0] == '.':
            transposed = [list(x) for x in zip(*s.split("\n"))] # learned this last year, tanspose rows and columns
            locks.append([6 - ''.join(line).index('#') for line in transposed]) # 6 is the height of the lock/key
        else:
            transposed = [list(x) for x in zip(*s.split("\n"))]
            keys.append([''.join(line).rindex('#') for line in transposed])

    return locks, keys

def part1(input):
    return sum(1 for key, lock in product(*parse_input(input)) if all(k + l < 6 for k, l in zip(key, lock)))

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")

if __name__ == "__main__":
    main()