from collections import defaultdict

def can_construct_design(design: list[str], patterns: list[str]) -> bool:
    n = len(design)
    dp = [False] * (n + 1)
    dp[0] = True
    
    patterns_by_length = defaultdict(set)
    for pattern in patterns:
        patterns_by_length[len(pattern)].add(pattern)
    
    for i in range(1, n + 1):
        for length, patterns_of_length in patterns_by_length.items():
            if i >= length:
                substring = design[i - length:i]
                if substring in patterns_of_length:
                    dp[i] = dp[i] or dp[i - length]
                    if dp[i]:  
                        break
    
    return dp[n]

# For both part 1 and two, sorting the patterns by length and storing them in a dictionary
# allows for a much faster solution, over x10 faster than the naive approach.
def part1(input: str) -> int:
    towel_patterns = input.split("\n\n")[0].split(", ")
    desired_designs = input.split("\n\n")[1].split("\n")

    return sum([1 for design in desired_designs if can_construct_design(design, towel_patterns)])

def count_arrangements(design: list[str], patterns:list[str]) -> int:
    n = len(design)
    dp = [0] * (n + 1)
    dp[0] = 1

    patterns_by_length = defaultdict(set)
    for pattern in patterns:
        patterns_by_length[len(pattern)].add(pattern)
        
    for i in range(1, n + 1):
        for length, patterns_of_length in patterns_by_length.items():
            if i >= length:
                substring = design[i - length:i]
                if substring in patterns_of_length:
                    dp[i] += dp[i - length]
                    
    return dp[n]

def part2(input: str) -> int:
    towel_patterns = input.split("\n\n")[0].split(", ")
    desired_designs = input.split("\n\n")[1].split("\n")

    return sum([count_arrangements(design, towel_patterns) for design in desired_designs])

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()