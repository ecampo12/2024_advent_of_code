from collections import Counter

def parse_input(input):
    list1, list2 = [], []
    for i in input:
        a, b = map(int, i.split())
        list1.append(a)
        list2.append(b)
    return list1, list2

def part1(input):
    list1, list2 = parse_input(input) 
    sorted_list1 = sorted(list1)
    sorted_list2 = sorted(list2)

    return sum([abs(sorted_list1[i] - sorted_list2[i]) for i in range(len(sorted_list1))])

def part2(input):
    list1, list2 = parse_input(input)
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