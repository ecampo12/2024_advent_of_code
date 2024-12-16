def parse_input(input_data):
    grid, moves = input_data.split("\n\n")
    grid = [list(line) for line in grid.splitlines()]
    moves = moves.replace("\n", "")
    return grid, moves

def find_robot_position(grid):
    for r, row in enumerate(grid):
        for c, cell in enumerate(row):
            if cell == "@":
                return (r, c)
            

def calculate_gps(grid):
    gps_sum = 0
    for r, row in enumerate(grid):
        for c, cell in enumerate(row):
            if cell in "O[":
                gps_sum += 100 * r + c
    return gps_sum

def part1(input):
    grid, moves = parse_input(input)
    r, c = find_robot_position(grid)
    directions = {"^": (-1, 0), "v": (1, 0), "<": (0, -1), ">": (0, 1)}
    for move in moves:
        dr, dc = directions[move]
        cr = r
        cc = c
        boxes = []
        moveable = True
        while True:
            cr += dr
            cc += dc
            char = grid[cr][cc]
            if char == "#":
                moveable = False
                break
            if char == "O":
                boxes.append((cr, cc))
            if char == ".":
                break
        if not moveable:
            continue
        grid[r][c] = "."
        grid[r + dr][c + dc] = "@"
        for br, bc in boxes:
            grid[br + dr][bc + dc] = "O"
        r += dr
        c += dc
    return calculate_gps(grid)

def scale_up(input):
    expand = {".": "..", "#": "##", "O": "[]", "@": "@."}
    grid = []
    for row in input:
        new_row = ""
        for c in row:
            new_row += expand[c]
        grid.append(list(new_row))
    return grid

def part2(input):
    grid, moves = parse_input(input)
    grid = scale_up(grid)
    r, c = find_robot_position(grid)
    directions = {"^": (-1, 0), "v": (1, 0), "<": (0, -1), ">": (0, 1)}
    for move in moves:
        dr, dc = directions[move]
        boxes = [(r, c)]
        moveable = True
        for cr, cc in boxes:
            nr = cr + dr
            nc = cc + dc
            if (nr, nc) in boxes:
                continue
            char = grid[nr][nc]
            if char == "#":
                moveable = False
                break
            if char == "[":
                boxes.append((nr, nc))
                boxes.append((nr, nc + 1))
            if char == "]":
                boxes.append((nr, nc - 1))
                boxes.append((nr, nc))
        if not moveable:
            continue
        copy = [row[:] for row in grid]
        grid[r][c] = "."
        grid[r + dr][c + dc] = "@"
        for br, bc in boxes[1:]:
            grid[br][bc] = "."
        for br, bc in boxes[1:]:
            grid[br + dr][bc + dc] = copy[br][bc]
        r += dr
        c += dc
    return calculate_gps(grid)

def main():
    input = open("input.txt", "r").read()
    part1_ans = part1(input)
    part2_ans = part2(input)
    print(f"Part 1: {part1_ans}")
    print(f"Part 2: {part2_ans}")

if __name__ == "__main__":
    main()