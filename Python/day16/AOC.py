import heapq
from collections import deque
import networkx as nx
import time

DIRECTIONS = [(-1, 0), (0, 1), (1, 0), (0, -1)]  

def parse_maze(maze):
    grid = maze.splitlines()
    start, end = None, None
    for r, row in enumerate(grid):
        for c, char in enumerate(row):
            if char == "S":
                start = (r, c)
            elif char == "E":
                end = (r, c)
    return grid, start, end

def dijkstra(grid, start, end):
    rows, cols = len(grid), len(grid[0])
    pq = []  
    visited = set()  
    start_state = (0, start[0], start[1], 1)
    heapq.heappush(pq, start_state)
    
    while pq:
        cost, r, c, direction = heapq.heappop(pq)
        
        if (r, c) == end:
            return cost
        
        if (r, c, direction) in visited:
            continue
        visited.add((r, c, direction))
        
        dr, dc = DIRECTIONS[direction]
        nr, nc = r + dr, c + dc
        if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] != "#":
            heapq.heappush(pq, (cost + 1, nr, nc, direction))
        
        heapq.heappush(pq, (cost + 1000, r, c, (direction + 1) % 4))
        heapq.heappush(pq, (cost + 1000, r, c, (direction - 1) % 4))
    
    return float('inf')

def part1(input):
    grid, start, end = parse_maze(input)
    return dijkstra(grid, start, end)


def dijkstra_with_paths(grid, start, end):
    pq = [(0, start[0], start[1], 1)]
    lowest_cost = {(start[0], start[1], 1): 0}
    backtrack = {}
    best_cost = float("inf")
    end_states = set()
    
    while pq:
        cost, r, c, direction = heapq.heappop(pq)
        if cost > lowest_cost.get((r, c, direction), float("inf")): continue
        if (r, c) == end:
            if cost > best_cost: break
            best_cost = cost
            end_states.add((r, c, direction))
        dr, dc = DIRECTIONS[direction]
        moves = [
            (cost + 1, r + dr, c + dc, direction), 
            (cost + 1000, r, c, (direction + 1) % 4), 
            (cost + 1000, r, c, (direction - 1) % 4)
        ]
        for new_cost, nr, nc, new_direction in moves:
            if grid[nr][nc] == "#":
                continue
            lowest = lowest_cost.get((nr, nc, new_direction), float("inf"))
            if new_cost > lowest: continue
            if new_cost < lowest:
                backtrack[(nr, nc, new_direction)] = set()
                lowest_cost[(nr, nc, new_direction)] = new_cost
            backtrack[(nr, nc, new_direction)].add((r, c, direction))
            heapq.heappush(pq, (new_cost, nr, nc, new_direction))
            
    return deque(end_states), set(end_states), backtrack

def part2(input):
    grid, start, end = parse_maze(input)
    states, seen, backtrack  = dijkstra_with_paths(grid, start, end)
    while states:
        key = states.popleft()
        for last in backtrack.get(key, []):
            if last in seen:
                continue
            states.append(last)
            seen.add(last)
            
    return len({(r, c) for r, c, _ in seen})

G = nx.DiGraph()
def build_graph(grid):
    rows, cols = len(grid), len(grid[0])
    # G = nx.DiGraph() 
    
    for r in range(rows):
        for c in range(cols):
            if grid[r][c] == "#":
                continue
            
            for direction, (dr, dc) in enumerate(DIRECTIONS):
                nr, nc = r + dr, c + dc
                if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] != "#":
                    G.add_edge((r, c, direction), (nr, nc, direction), weight=1)
                    
                G.add_edge((r, c, direction), (r, c, (direction + 1) % 4), weight=1000)
                G.add_edge((r, c, direction), (r, c, (direction - 1) % 4), weight=1000)
    return G

# Alternative solutions, using networkx. Just wanted to see if it was faster.
def p1_alt(input):
    grid, start, end = parse_maze(input)
    G = build_graph(grid)
    
    start_state = (*start, 1)
    end_state = [(end[0], end[1], d) for d in range(4)]
    min_cost = float("inf")
    
    for end in end_state:
        cost = nx.shortest_path_length(G, start_state, end, weight="weight")
        min_cost = min(min_cost, cost)
    return min_cost

def p2_alt(input):
    _, start, end = parse_maze(input)
    # G = build_graph(grid)
    start_state = (*start, 1)
    end_state = [(end[0], end[1], d) for d in range(4)]
    seen = set()
    
    for end in end_state:
        paths = nx.all_shortest_paths(G, start_state, end, weight="weight")
        [seen.add(state) for path in paths for state in path]
        
    return len({(r, c) for r, c, _ in seen})

def main():
    input = open("input.txt", "r").read()
    t1 = time.perf_counter()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")
    print(f"Time: {time.perf_counter() - t1:.3f} seconds")
    
    print("-" * 10)
    
    t1 = time.perf_counter()
    print(f"Part 1 Alt: {p1_alt(input)}")
    print(f"Part 2 Alt: {p2_alt(input)}")
    print(f"Time: {time.perf_counter() - t1:.3f} seconds")
if __name__ == "__main__":
    main()