from collections import deque

# since there is only one path from start to end, we can just grab the ponits on the path.
# Populate the grid with the distance from the start to each point.
def parse_input(input):
    start, end = None, None
    grid = {(r, c): -1 for r, row in enumerate(input.split('\n')) for c, char in enumerate(row) if char in '.SE'}
    for r, row in enumerate(input.split("\n")):
        for c, char in enumerate(row):
            if char == "S":
                start = (r, c)
            if char == "E":
                end = (r, c)
    grid[start] = 0
    
    r, c = start
    while (r, c) != end:
        for nr, nc in [(r, c + 1), (r, c - 1), (r + 1, c), (r - 1, c)]:
            if (nr, nc) not in grid: continue
            if grid[(nr, nc)] != -1: continue
            grid[(nr, nc)] = grid[(r, c)] + 1
            r = nr
            c = nc
    
    return grid

# Basically walk the path and check if we can jump to a point on the path.
# For part 1 we only need to check the points that are 2 steps away from each other.
def part1(input, save=100):
    grid = parse_input(input)
    cheats = []
    for pos in grid:
        r, c = pos
        for nr, nc in [(r, c + 2), (r, c - 2), (r + 2, c), (r - 2, c)]:
            if (nr, nc) not in grid: continue
            if grid[(r, c)] - grid[(nr, nc)] >= save + 2: # +2 to account for the 2 steps between the points ??
                cheats.append(grid[(r, c)] - grid[(nr, nc)])
    return len(cheats)

# Same as part 1, but we need to check all the points that are 2 to 20 steps away from each other.
def part2(input, save=100):
    grid = parse_input(input)
    cheats = []
    for pos in grid:
        r, c = pos
        for radius in range(2, 21):
            for dr in range(radius + 1):
                dc = radius - dr
                for nr, nc in {(r + dr, c + dc), (r + dr, c - dc), (r - dr, c + dc), (r - dr, c - dc)}: # a set to avoid duplicates
                    if (nr, nc) not in grid: continue
                    if grid[(r, c)] - grid[(nr, nc)] >= save + radius:
                        cheats.append(grid[(r, c)] - grid[(nr, nc)])
    return len(cheats)

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()