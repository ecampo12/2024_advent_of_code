def part1(input):
    disk = []
    id = 0
    for i, char in enumerate(input):
        if i % 2 == 0:
            disk += [id] * int(char)
            id += 1
        else:
            disk += [-1] * int(char)

    for i in range(len(disk)):
        while disk[-1] == -1:
            disk.pop()
        if len(disk) <= i:
            break
        if disk[i] == -1:
            disk[i] = disk.pop()

    return sum(disk[i] * i for i in range(len(disk)))

# the problem become alot easier fi we don't represent the disk as a list.
# Keep track of the position and size of each file and the free space makes it easier to solve.
def part2(input):
    files = {}
    free_space = []
    
    id = 0
    pos = 0
    for i, char in enumerate(input):
        if i % 2 == 0:
            files[id] = (pos, int(char))
            id += 1
        else:
            free_space.append((pos, int(char)))
        pos += int(char)
        
    while id > 0:
        id -= 1
        pos, size = files[id]
        for i, (start, length) in enumerate(free_space):
            if start >= pos:
                free_space = free_space[:i]
                break
            if size <= length:
                files[id] = (start, size)
                if size == length:
                    free_space.pop(i)
                else:
                    free_space[i] = (start + size, length - size)
                break
            
    return sum(id * x for id, (pos, size) in files.items() for x in range(pos, pos + size))

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()