import re
import time

def part1(input, width=101, height=103):
    grid = [[0 for _ in range(width)] for _ in range(height)]
    robots = [[x, y, vx, vy] for x, y, vx, vy in [map(int, re.findall(r"(-?\d+)", line)) for line in input]]
    qs = [0,0,0,0]
    w_half = width // 2
    h_half = height // 2
    for r in robots:
        r[0] += 100*r[2]
        r[0] %= width
        r[1] += 100*r[3]
        r[1] %= height
        grid[r[1]][r[0]] += 1

        if r[0] < w_half and r[1] < h_half:
            qs[0] += 1
        elif r[0] > w_half and r[1] < h_half:
            qs[1] += 1
        elif r[0] < w_half and r[1] > h_half:
            qs[2] += 1
        elif r[0] > w_half and r[1] > h_half:
            qs[3] += 1
    return qs[0] * qs[1] * qs[2] * qs[3]

# Pervious brute force solution showed me that the easter egg is within a box of robots, 
# so just find part of the box to find when the second the easter egg is revealed.
def part2(input):
    width = 101
    height = 103
    robots = [[x, y, vx, vy] for x, y, vx, vy in [map(int, re.findall(r"(-?\d+)", line)) for line in input]]

    for i in range(100000):
        grid = [['.' for _ in range(width)] for _ in range(height)]
        for r in robots:
            grid[r[1]][r[0]] = '#'
            r[0] += r[2]
            r[0] %= width
            r[1] += r[3]
            r[1] %= height

        for h in range(height):
            if re.search(r'\#{10,}', ''.join(grid[h])): # 10 consecutive '#', because that sus
                for row in grid:
                    print(''.join(row))
                return i

    return 0

def main():
    input = open("input.txt", "r").read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()