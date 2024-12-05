def parse_input(input):
    rules_section, updates_section = input.split("\n\n")
    rules = [tuple(map(int, i.split("|"))) for i in rules_section.split("\n")]
    updates = [list(map(int, i.split(","))) for i in updates_section.split("\n")]
    return rules, updates

def is_valid(rules, update):
    num_to_index = {update: i for i, update in enumerate(update)}
    for a, b in rules:
        if a in num_to_index and b  in num_to_index:
            if num_to_index[a] > num_to_index[b]:
                return False, (a, b)
    return True, None

def part1(input):
    rules, updates = parse_input(input)
    return sum(update[len(update) // 2] for update in updates if is_valid(rules, update)[0])

def correct_update(relation, update):
    a, b = relation
    index_a = update.index(a)
    index_b = update.index(b)
    update[index_a], update[index_b] = update[index_b], update[index_a]
    return update

def part2(input):
    rules, updates = parse_input(input)
    sum = 0
    
    for update in updates:
        valid, rule = is_valid(rules, update)
        if not valid:
            while True:
                update = correct_update(rule, update)
                valid, rule = is_valid(rules, update)
                if valid:
                    sum += update[len(update) // 2]
                    break
    return sum

def main():
    input = open("input.txt", "r").read()
    print(f"Part 1: {part1(input)}")
    print(f"Part 2: {part2(input)}")

if __name__ == "__main__":
    main()