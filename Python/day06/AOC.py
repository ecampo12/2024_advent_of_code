def find_start(input):
    for row in range(len(input)):
        for col in range(len(input[row])):
            if input[row][col] == "^":
                return (row, col)
    return None

def traverse(input, start, part2=False):
    directions = [(-1, 0), (0, -1), (1, 0), (0, 1)]    
    curr = start
    path = set()
    facing = 0
    while True > 0:
        path.add(curr)
        dr, dc = directions[facing]
        if dr + curr[0] < 0 or dr + curr[0] >= len(input) or dc + curr[1] < 0 or dc + curr[1] >= len(input[0]):
            break
        if input[curr[0] + dr][curr[1]+ dc] == "#":
            facing = (facing - 1) % 4
        else:
            curr = (curr[0] + directions[facing][0], curr[1] + directions[facing][1])
    return path
    
def part1(input):
    grid = input.splitlines()
    return len(traverse(grid, find_start(grid)))

def loop(grid, start):
    directions = [(-1, 0), (0, -1), (1, 0), (0, 1)]
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
    start = find_start(input)
    paths = traverse(input, start)
    count = 0
    for p in paths:
        row, col = p
        if p != start:
            grid = input.copy()
            grid[row] = grid[row][:col] + "#" + grid[row][col+1:]
            if loop(grid, start):
                count += 1
    return count

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()