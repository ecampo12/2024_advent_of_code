import re
def run_program(registers, program):
    A, B, C = registers
    
    def combo_value(operand):
        if operand < 4 : return operand
        elif operand == 4: return A
        elif operand == 5: return B
        elif operand == 6: return C
        else: raise ValueError("Invalid combo operand")
    
    output = []
    pointer = 0  

    while pointer < len(program):
        opcode = program[pointer]
        operand = program[pointer + 1]
        match opcode:
            case 0:
                A = A // (2 ** combo_value(operand))
            case 1:
                B = B ^ operand
            case 2:
                B = combo_value(operand) % 8
            case 3:
                if A != 0:
                    pointer = operand
                    continue
            case 4:
                B = B ^ C
            case 5:
                output.append(combo_value(operand) % 8)
            case 6:
                B = A // (2 ** combo_value(operand))
            case 7:
                C = A // (2 ** combo_value(operand))
            case _:
                raise ValueError(f"Invalid opcode {opcode}")
        
        pointer += 2
        
    return output

def part1(input):
    parts = input.split("\n\n")
    registers = map(int, re.findall(r"\d+", parts[0]))
    program = list(map(int, re.findall(r"\d+", parts[1])))
    return ','.join(map(str, run_program(registers, program)))

# Recursive function to find the value of register A, checking from the end of the program to the start.
# The way the input program seems to work is that the output is dependent on the value of register A
# the value of register A divided by 8 after output. The input also only has single digits values, no more than 8.
# The input seesm to be octals, A seems to be an octal number, theres also 8's all over the problem. 
def find_Ahh(program, index, A):
    for candidate in range(8):
        if run_program((A * 8 + candidate, 0, 0), program) == program[index:]:
            if index == 0:
                return A * 8 + candidate
            ret = find_Ahh(program, index - 1, A * 8 + candidate)
            if ret is not None:
                return ret
    return None

def part2(input):
    program = list(map(int, re.findall(r"\d+",input.split("\n\n")[1])))
    return find_Ahh(program, len(program) - 1, 0)

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()