from collections import defaultdict
from itertools import combinations

def parse_input(input):
    antennas = defaultdict(list)
    for i, line in enumerate(input):
        for j, char in enumerate(line):
            if char.isdigit() or char.isalpha():                
                antennas[char].append((i, j))
    return antennas

def get_point(a, b):
    dr, dc = (a[0] - b[0], a[1] - b[1])
    return (a[0] + dr, a[1] + dc)

def part1(input):
    width = len(input[0])
    height = len(input)
    antennas = parse_input(input)
    antinode = set()
    for _ , points in antennas.items():
        for a, b in combinations(points, 2):
            for p in(get_point(a, b), get_point(b, a)):
                if 0 <= p[0] < height and 0 <= p[1] < width:
                    antinode.add(p)
    return len(antinode)

def points_on_line(a, b, width, height):
    dr, dc = (a[0] - b[0], a[1] - b[1])
    points = {a, b}
    for direction in (-1, 1):
        for i in range(1, max(width, height)):
            p = (a[0] + dr*i*direction, a[1] + dc*i*direction)
            if 0 <= p[0] < height and 0 <= p[1] < width:
                points.add(p)
            else:
                break
    return points
    
def part2(input):
    width = len(input[0])
    height = len(input)
    antennas = parse_input(input)
    antinode = set()
    for freq in antennas:
        for a, b in combinations(antennas[freq], 2):
            antinode.update(points_on_line(a, b, width, height))
    return len(antinode)

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_ans = part1(input)
    part2_ans = part2(input)
    print(f"Part 1: {part1_ans}")
    print(f"Part 2: {part2_ans}")

if __name__ == "__main__":
    main()