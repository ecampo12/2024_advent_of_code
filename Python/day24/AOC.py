import re
from itertools import combinations
from collections import namedtuple

Gate = namedtuple("Gate", ["op", "left", "right", "output"])

OPERATORS = {
    "OR": lambda x, y: x | y,
    "AND": lambda x, y: x & y,
    "XOR": lambda x, y: x ^ y,
}

def parse_input(input_data):
    parts = input_data.split("\n\n")
    pattern = r"(.+?)\s+(AND|OR|XOR)\s+(.+?)\s+->\s+(.+)"
    values, gates = {}, []
    for part in parts[0].split("\n"):
        key, value = part.split(":")
        values[key] = int(value)   
    for part in parts[1].split("\n"):
        left, op, right, output = re.match(pattern, part).groups()
        gates.append(Gate(op, left, right, output))
    return values, gates


def simulate_system(values, gates):
    wires = dict(values)
    unresolved_gates = gates[:]

    while unresolved_gates:
        for g in gates:
            if g.output in wires:
                continue
            if g.left in wires and g.right in wires:
                wires[g.output] = OPERATORS[g.op](wires[g.left], wires[g.right])
                unresolved_gates.remove(g)
    return wires

def calculate_output(wires):
    z_wires = sorted(((k, v) for k, v in wires.items() if k.startswith("z")), reverse=True)
    return int("".join(str(v[1]) for v in z_wires), 2)

def part1(input):
    return calculate_output(simulate_system(*parse_input(input))) 

# We are dealing with a ripple-carry adder here (https://www.circuitstoday.com/ripple-carry-adder).
# It adds two 45-bit numbers, x and y, and outputs a 46-bit number z. The first 45 bits of z are the sum of the corresponding bits in x and y, and the last bit is the carry out.
# Illustrations of what can be found in this folder and here: https://www.reddit.com/r/adventofcode/comments/1hl8tl4/2024_day_24_part_2_using_a_graph_visualization/
# Bascially, we "traverse" the graph of operations and check if we can get to the end. If we can't, we try to swap two outputs and see if we can get further.
def check_ripple(gates):
    ops = {}
    for op, x1, x2, output in gates:
        ops[(frozenset([x1, x2]), op)] = output
    def get_output(x1, x2, op):
        return ops.get((frozenset([x1, x2]), op), None)
    carries = {}
    correct = set()
    prev_intermediates = set()
    for i in range(45):
        pos = f"0{i}" if i < 10 else str(i)
        predigit = get_output(f"x{pos}", f"y{pos}", "XOR")
        precarry1 = get_output(f"x{pos}", f"y{pos}", "AND")
        if i == 0:
            carries[i] = precarry1
            continue
        digit = get_output(carries[i - 1], predigit, "XOR")
        if digit != f"z{pos}":
            return i - 1, correct
        correct.add(carries[i - 1])
        correct.add(predigit)
        for wire in prev_intermediates:
            correct.add(wire)
        precarry2 = get_output(carries[i - 1], predigit, "AND")
        carry_out = get_output(precarry1, precarry2, "OR")
        carries[i] = carry_out
        prev_intermediates = set([precarry1, precarry2])
    return i, correct

def part2(input):
    _, gates = parse_input(input)
    swaps = []

    base, base_used = check_ripple(gates)
    for _ in range(4):
        for i, j in combinations(range(len(gates)), 2):
            op_i, left_i, right_i, ouput_i = gates[i]
            op_j, left_j, right_j, output_j = gates[j]
            if "z00" in (ouput_i, output_j):
                continue
            if ouput_i in base_used or output_j in base_used:
                continue
            gates[i] = op_i, left_i, right_i, output_j
            gates[j] = op_j, left_j, right_j, ouput_i
            attempt, attempt_used = check_ripple(gates)
            if attempt > base:
                swaps.extend([ouput_i, output_j])
                base, base_used = attempt, attempt_used
                break
            gates[i] = op_i, left_i, right_i, ouput_i
            gates[j] = op_j, left_j, right_j, output_j

    return ",".join(sorted(swaps))

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()