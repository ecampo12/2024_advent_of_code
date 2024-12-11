from functools import cache

def part1(input, blinks=25):
    nums = list(map(int, input.split(" ")))
    for _ in range(blinks):
        new_nums = []
        for n in nums:
            stone_str = str(n)
            length = len(stone_str)
            if len(str(n)) % 2 == 0:
                new_nums.append(int(stone_str[:length//2]))
                new_nums.append(int(stone_str[length//2:]))
            elif n == 0:
                new_nums.append(1)
            else:
                new_nums.append(n*2024)
        nums = new_nums
    return len(nums)


@cache
# Counts the number of stones after steps steps instead of keeping track of a list of stones.
def count_stones(stone, steps):
    if steps == 0:
        return 1
    if stone == 0:
        return count_stones(1, steps-1)
    stone_str = str(stone)
    length = len(stone_str)
    if length % 2 == 0:
        left = int(stone_str[:length//2])
        right = int(stone_str[length//2:])
        return  count_stones(left, steps-1) + count_stones(right, steps-1)
    
    return count_stones(stone*2024, steps-1)

def part2(input):
    return sum(count_stones(stone, 75) for stone in map(int, input.split(" ")))

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()