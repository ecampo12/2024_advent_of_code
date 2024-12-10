from collections import defaultdict

def find_paths(input):
    starts = [(r, c) for r in range(len(input)) for c in range(len(input[0])) if input[r][c] == "0"]
    all_paths = []
    for start in starts:
       all_paths.extend(dfs(input, start, 0, []))
    return all_paths

def dfs(input, pos, num, path):
    i, j = pos
    if input[i][j] != str(num): return []
    path = path + [pos]
    if num == 9:
        return [path]
    paths = []
    for new_pos in [(i-1, j), (i+1, j), (i, j-1), (i, j+1)]:
        if 0 <= new_pos[0] < len(input) and 0 <= new_pos[1] < len(input[0]):
            paths.extend(dfs(input, new_pos, num+1, path))
    return paths

def trailheadRating(paths):
    trailheads = defaultdict(set)
    for path in paths:
        trailheads[path[0]].add(path[-1])
    return trailheads
    
def part1(input):
    return sum(len(t) for t in trailheadRating(find_paths(input)).values())

def part2(input):    
    return len(find_paths(input))

def main():
    input = open("input.txt", "r").read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()