def traverse(input, part2=False):
    directions = [(-1, 0), (0, -1), (1, 0), (0, 1)]
    start = (0, 0)
    for row in range(len(input)):
        for col in range(len(input[row])):
            if input[row][col] == "^":
                start = (row, col)
                break
        else:
            continue
        break
    
    curr = start
    path = set()
    facing = 0
    while True > 0:
        if part2:
            path.add((curr, facing))
        else:
            path.add(curr)
        dr, dc = directions[facing]
        if dr + curr[0] < 0 or dr + curr[0] >= len(input) or dc + curr[1] < 0 or dc + curr[1] >= len(input[0]):
            if part2:
                return False
            break
        if input[curr[0] + dr][curr[1]+ dc] == "#":
            facing = (facing - 1) % 4
        else:
            curr = (curr[0] + directions[facing][0], curr[1] + directions[facing][1])
        
        if part2 and (curr, facing) in path:
            return True
    return path
    
def part1(input):
    return len(traverse(input.splitlines()))

def loop(grid):
    directions = [(-1, 0), (0, -1), (1, 0), (0, 1)]
    start = (0, 0)
    for row in range(len(grid)):
        for col in range(len(grid[row])):
            if grid[row][col] == "^":
                start = (row, col)
                break
        else:
            continue
        break
    
    curr = start
    path = set()
    facing = 0
    while True > 0:
        path.add((curr, facing))
        dr, dc = directions[facing]
        if dr + curr[0] < 0 or dr + curr[0] >= len(grid) or dc + curr[1] < 0 or dc + curr[1] >= len(grid[0]):
            return False
        if grid[curr[0] + dr][curr[1]+ dc] == "#":
            facing = (facing - 1) % 4
        else:
            curr = (curr[0] + directions[facing][0], curr[1] + directions[facing][1])
        if (curr, facing) in path:
            return True
    
def part2(input):
    input = input.splitlines()
    paths = traverse(input)
    count = 0
    for p in paths:
        row, col = p
        if input[row][col] == ".":
            grid = input.copy()
            grid[row] = grid[row][:col] + "#" + grid[row][col+1:]
            if traverse(grid, True):
                count += 1
    return count

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()