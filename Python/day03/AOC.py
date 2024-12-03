import re
def part1(input):
    return sum([int(match[0]) * int(match[1]) for match in re.findall(r"mul\((\d+),(\d+)\)", input)])

# The input had new lines which threw off the regex, so I had to use re.DOTALL
# used re.sub to remove the don't() part of the string
def part2(input):
    return part1(re.sub(r"don't\(\).*?(?=do\(\)|$)", "", input, flags=re.DOTALL))
    

def main():
    input = open("input.txt", "r").read()
    part1_ans = part1(input)
    part2_ans = part2(input)
    print(f"Part 1: {part1_ans}")
    print(f"Part 2: {part2_ans}")

if __name__ == "__main__":
    main()