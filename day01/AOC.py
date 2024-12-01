from collections import Counter

def part1(input):
    list1, list2 = map(list, zip(*[map(int, i.split()) for i in input]))
    return sum([abs(a - b) for a, b in zip(sorted(list1), sorted(list2))])

def part2(input):
    list1, list2 = map(list, zip(*[map(int, i.split()) for i in input]))
    count = Counter(list2)
    return sum([i * c for i in list1 for c in [count[i] if i in count else 0]])

def main():
    input = open("input.txt", "r").read().splitlines()
    part1_ans = part1(input)
    part2_ans = part2(input)
    print(f"Part 1: {part1_ans}")
    print(f"Part 2: {part2_ans}")

if __name__ == "__main__":
    main()