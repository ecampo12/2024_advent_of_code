# Original solution invloved finding subgrids of 4x4 and checking if they contained the word "XMAS".
# That solution resulted in a lot of double counting and got really messy.
# I then realized that I could just check for the word "XMAS" in all 8 directions from each "X" and that would be enough.
def part1(input):
    count = 0
    dirs = [(-1, 0), (1, 1), (1, 0), (1, -1), (0, -1), (-1, -1), (-1, 1), (0, 1)]
    rows, cols = len(input), len(input[0])
    for row in range(rows):
        for col in range(cols):
            if input[row][col] != "X":
                continue
            for dr, dc in dirs:
                if 0 <= row + 3 * dr < rows and 0 <= col + 3 * dc < cols:
                    if "".join([input[row + i*dr][col + i*dc] for i in range(4)]) == "XMAS":
                        count += 1
    return count

def part2(input):
    count = 0
    rows, cols = len(input), len(input[0])
    valid_corners = {"MMSS", "MSSM", "SSMM", "SMMS"} # all possible corner combinations if read clockwise
    for row in range(1, rows - 1):
        for col in range(1, cols - 1):
            if input[row][col] != "A":
                continue
            corners = (
                input[row - 1][col - 1] +
                input[row - 1][col + 1] +
                input[row + 1][col + 1] +
                input[row + 1][col - 1]
            )
            if corners in valid_corners: 
                count += 1
            
    return count


def main():
    input = open("input.txt", "r").read().splitlines()
    part1_ans = part1(input)
    part2_ans = part2(input)
    print(f"Part 1: {part1_ans}")
    print(f"Part 2: {part2_ans}")

if __name__ == "__main__":
    main()