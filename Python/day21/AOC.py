from itertools import permutations
from functools import cache

NUMERIC_KEYPAD = {
    '7': (0, 0), '8': (0, 1), '9': (0, 2),
    '4': (1, 0), '5': (1, 1), '6': (1, 2),
    '1': (2, 0), '2': (2, 1), '3': (2, 2),
    '0': (3, 1), 'A': (3, 2),
}

DIRECTIONAL_KEYPAD = {
    '^': (0, 1), 'A': (0, 2),
    '<': (1, 0), 'v': (1, 1), '>': (1, 2)
}

MOVES = {
    '^': (-1, 0), 'v': (1, 0), '<': (0, -1), '>': (0, 1)
}

def sequence_to_moveset(start, end, keypad):
    moves = []
    dx, dy = tuple(map(lambda x, y: x - y, end, start))
    if dx < 0:
        moves.extend(['^'] * abs(dx))
    else:
        moves.extend(['v'] * dx)
    if dy < 0:
        moves.extend(['<'] * abs(dy))
    else:
        moves.extend(['>'] * dy)
    
    valid_sequences = []
    for p in set(permutations(moves)):
        positions = [start]
        valid = True
        for move in p:
            next_pos = tuple(map(sum, zip( positions.pop(), MOVES[move])))
            if next_pos not in keypad.values():
                valid = False
                break
            positions.append(next_pos)
        if valid:
            valid_sequences.append(''.join(p) + 'A')
    
    return valid_sequences

@cache
def min_length(sequence: str, limit: int = 2, depth: int = 0) -> int:
    keypad = NUMERIC_KEYPAD if depth == 0 else DIRECTIONAL_KEYPAD
    current = keypad['A']
    
    lengths = []
    for char in sequence:
        next_pos = keypad[char]
        movesets = sequence_to_moveset(current, next_pos, keypad)
        if depth >= limit:
            lengths.append(len(min(movesets, key=len)))
        else:
            lengths.append(min(min_length(moveset, limit, depth + 1) for moveset in movesets))
        current = next_pos
    
    return sum(lengths)

def part1(input: list[str]) -> int:
    return sum(min_length(code) * int(code[:-1]) for code in input)

def part2(input: list[str]) -> int:
    return sum(min_length(code, limit=25) * int(code[:-1]) for code in input)

def main():
    input = open("input.txt", "r").read().splitlines()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()