import re
from collections import deque

def simulate_memory_fall(grid: dict, size: int) -> int:
    goal = (size - 1, size - 1)
    queue = deque([(0, 0, 0)])
    visited = set()
    visited.add((0, 0))
    
    directions = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    
    while queue:
        x, y, steps = queue.popleft()
        
        if (x, y) == goal:
            return steps
        
        for dx, dy in directions:
            nx, ny = x + dx, y + dy
            
            if (nx, ny) in grid and (nx, ny) not in visited:
                if grid[(nx, ny)] == ".":
                    visited.add((nx, ny))
                    queue.append((nx, ny, steps + 1))
                    
    return -1

# New favorite trick: using a dict to represent a grid. I makes bounds checking a lot easier.
def part1(input: list[str], bytes: int =1024, size: int =71) -> int:
    pos = [(x, y) for x, y in [map(int, re.findall(r"\d+", line)) for line in input]]
    grid = {(x, y): "." for x in range(size) for y in range(size)}
    pos = pos[:bytes]
    
    for x, y in pos:
        grid[(x, y)] = "#"
        
    return simulate_memory_fall(grid, size)

# Add all positions to the grid, then remove them one by one until I 
# find the position that makes the makes end goal unreachable.
def part2(input: list[str], bytes: int =1024, size: int=71):
    pos = [(x, y) for x, y in [map(int, re.findall(r"\d+", line)) for line in input]]
    grid = {(x, y): "." for x in range(size) for y in range(size)}
    
    for x, y in pos:
            grid[(x, y)] = "#"
            
    for i in range(len(pos) - 1, bytes, -1):
        grid[pos[i]] = "."
        steps = simulate_memory_fall(grid, size)
        if steps != -1:
            return ','.join(map(str, pos[i]))
        
    return 0

def main():
    input = open("input.txt", "r").read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()