from collections import defaultdict

Prices = {}
def get_secret_number(secret: int) -> int:
    key = secret
    Prices[key] = []
    for _ in range(2000):
        secret ^= (secret * 64) % 16777216
        secret ^= (secret // 32) % 16777216
        secret ^= (secret * 2048) % 16777216
        if key in Prices: # we need this for part 2, so we'll calculate these in part 1
            Prices[key].append(secret%10)
    return secret

def part1(input: list[str]) -> int:
    return sum([get_secret_number(initial_secret) for initial_secret in [int(line) for line in input]])

def get_sequences(price: list[int], changes: list[int]) -> dict:
    sequence = defaultdict(int)
    for i in range(len(changes)-3):
        pattern = tuple(changes[i:i+4])
        if pattern not in sequence:
            sequence[pattern] = price[i+4]
            
    return sequence

def part2(input: list[str]) -> int:
    scores = defaultdict(int)
    for line in input:
        if int(line) not in Prices: # should literally only happen during testing, since the example input changes.
            get_secret_number(int(line))
        price = Prices[int(line)]
        changes = [price[i+1] - price[i] for i in range(len(price)-1)]
        seq = get_sequences(price, changes)
        for k,v in seq.items():
            scores[k] += v

    return max(scores.values())

def main():
    input = open("input.txt", "r").read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()