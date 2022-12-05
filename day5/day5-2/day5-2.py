import re

stacks = [
    ['T', 'P', 'Z', 'C', 'S', 'L', 'Q', 'N'],
    ['L', 'P', 'T', 'V', 'H', 'C', 'G'],
    ['D', 'C', 'Z', 'F'],
    ['G', 'W', 'T', 'D', 'L', 'M', 'V', 'c'],
    ['P', 'W', 'C'],
    ['P', 'F', 'J', 'D', 'C', 'T', 'S', 'Z'],
    ['V', 'W', 'G', 'B', 'D'],
    ['N', 'J', 'S', 'Q', 'H', 'W'],
    ['R', 'C', 'Q', 'F', 'S', 'L', 'V']
]


def parse_input():
    choice = []
    with open("../day5.txt") as file:
        for line in file:
            total, src, dst = list(map(int, re.findall(r"\d+", line.rstrip())))
            choice.append([total, src, dst])
    return choice

def part2():
    choices = parse_input()
    for choice in choices:
        total, src, dst = choice
        tmpArr = stacks[src-1]
        lastItems = tmpArr[-total:]
        beginingItems = tmpArr[:-total]
        stacks[src - 1] = beginingItems
        stacks[dst - 1].extend(lastItems)
    result = ""
    for stack in stacks:
        result += stack[-1]
    return result

print(''.join(part2()))
