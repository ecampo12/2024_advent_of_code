from collections import defaultdict

def flood_fill(grid, r, c, letter, seen):
    if r < 0 or r >= len(grid) or c < 0 or c >= len(grid[0]) or grid[r][c] != letter:
        return [], 1
    if (r, c) in seen:
        return [], 0
    
    seen.add((r, c))
    points = [(r, c)]
    perimeter = 0
    
    for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
        a, p = flood_fill(grid, r + dr, c + dc, letter, seen)
        points.extend(a)
        perimeter += p
        
    return points, perimeter
    
def part1(input):
    region = []
    seen = set()
    for r in range(len(input)):
        for c in range(len(input[0])):
            if (r, c) not in seen:
                letter = input[r][c]
                area, perimeter = flood_fill(input, r, c, letter, seen)
                region.append((letter, area, perimeter))
                
    return sum(len(area) * perimeter for _, area, perimeter in region)

# Fun Fact: The number of corners of a shape is equal to the number of sides.
# We use a half cartesian coordinate system to find the corners of a shape by checking each direction from a point.
# The number of points within the region tells us if the point is a corner or not.
# example: where 'A" is a crop region
#  A |                                              A |   two points in region means it is not a corner, A |          | A
# ___| onw point in region means it is a corner     A |   less the points are across from each other     __|__ or   __|__
#                                                                                                          | A or   A | 
#                                                                                                             
#  A | A                                                    A | A 
# ___|___ three points in region means it is a corner       __|__ four means we're in the middle of the region
#  A |                                                      A | A                                           
def find_sides(region):
    possible_corners = set()
    for r, c in region:
        for nr, nc in [(r - 0.5, c - 0.5), (r + 0.5, c - 0.5), (r + 0.5, c + 0.5), (r - 0.5, c + 0.5)]:
            possible_corners.add((nr, nc))
            
    corners = 0
    for r, c in possible_corners:
        locations = [(nr, nc) in region for nr, nc in [(r - 0.5, c - 0.5), (r + 0.5, c - 0.5), (r + 0.5, c + 0.5), (r - 0.5, c + 0.5)]]
        num = sum(locations)
        if num == 1 or num == 3:
            corners += 1
        elif num == 2:
            if locations[0] and locations[2] or locations[1] and locations[3]:
                corners += 2
    return corners
    

def part2(input):
    region = []
    seen = set()
    for r in range(len(input)):
        for c in range(len(input[0])):
            if (r, c) not in seen:
                letter = input[r][c]
                area, perimeter = flood_fill(input, r, c, letter, seen)
                region.append((letter, area, perimeter))
                
    return sum(len(area) * find_sides(area) for _, area, _ in region)

def main():
    input = open("input.txt", "r").read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()