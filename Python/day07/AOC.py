import re
import time
from functools import lru_cache

def parse_input(input):
    lines = input.split("\n")
    equations = []
    for line in lines:
        nums = list(map(int, re.findall(r"\d+", line)))
        equations.append((nums[0], tuple(nums[1:])))
    return equations

# cache and exiting early if the current sum is greater than the target, reduces the runtime from 7 seconds to 4 seconds
@lru_cache
def evaluate(nums: list[int], target: int, curr=0, index=0, part2=False):
    if curr > target:
        return False
    if index == len(nums):
       return curr == target
   
    num = nums[index]
    
    add = evaluate(nums, target, curr + num, index + 1, part2)
    if add:
        return True
    
    mult = evaluate(nums, target, curr * num if curr != 0 else num, index + 1, part2)
    if mult:
        return True
    
    if part2:
        concat = evaluate(nums, target, int(f'{curr}{num}'), index + 1, part2)
        if concat:
            return True
    
    return False
def part1(input):
    return sum(eq[0] for eq in parse_input(input) if evaluate(eq[1], eq[0]))

def part2(input):
    return sum(eq[0] for eq in parse_input(input) if evaluate(eq[1], eq[0], part2=True))

def main():
    input = open("input.txt", "r").read()
    t1 = time.perf_counter()
    print(f"Part 1: {part1(input)}")
    print(f"Took: {time.perf_counter() - t1:.3f} seconds")
    t1 = time.perf_counter()
    print(f"Part 2: {part2(input)}")
    print(f"Took: {time.perf_counter() - t1:.3f} seconds")

if __name__ == "__main__":
    main()