import re
from z3 import Ints, Optimize, sat

# Basically we are solving a linear Diophantine equations:
#       xA * x + xB * y = pX
#       yA * x + yB * y = pY
# where x, y are the number of coins and pX, pY are the position of the prize.
# The problem specifies that the cost of the machine is 3 * x + y, so we need to minimize this cost.
def solve_machine(machines, error = 0):
    coins = 0

    for machine in machines:
        xA, yA, xB, yB, pX, pY = machine
        pX += error
        pY += error
        
        s = Optimize()
        x, y = Ints("x y")
        s.add(x >= 0, y >= 0)
        s.add(xA* x + xB * y == pX)
        s.add(yA * x + yB * y == pY)
        cost = 3 * x + y
        s.minimize(cost)
        if s.check() == sat:
            min_cost = s.model().evaluate(cost).as_long()
            coins += min_cost
    return coins

# z3 can be used solve both parts, so I did.
def part1(input):
    return solve_machine([list(map(int, re.findall(r"(\d+)", parts))) for parts in input.split("\n\n")])

def part2(input):
    return solve_machine([list(map(int, re.findall(r"(\d+)", parts))) for parts in input.split("\n\n")], 10000000000000)

def main():
    input = open("input.txt", "r").read()
    part1_ans = part1(input)
    part2_ans = part2(input)
    print(f"Part 1: {part1_ans}")
    print(f"Part 2: {part2_ans}")

if __name__ == "__main__":
    main()