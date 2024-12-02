def is_safe(report):
    diff = [report[i + 1] - report[i] for i in range(len(report) - 1)]
    return all(x in {1, 2, 3} for x in diff) or all(x in {-1, -2, -3} for x in diff)

def part1(input):
    reports = (list(map(int, x.split())) for x in input)
    return sum(1 for report in reports if is_safe(report))

def part2(input):
    reports = (list(map(int, x.split())) for x in input)
    count = 0
    for report in reports:
        if is_safe(report):
            count += 1
        else:
            if any(is_safe(report[:i] + report[i+1:]) for i in range(len(report))):
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