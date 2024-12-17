import pytest
from AOC import *

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ("""Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0""", "4,6,3,5,6,3,5,2,1,0"),
    ])
def test_part1(puzzle_input, expected):
    """Test part 1 on example input."""
    assert part1(puzzle_input) == expected

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ("""Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0""", 117440),
    ])
def test_part2(puzzle_input, expected):
    """Test part 2 on example input."""
    assert part2(puzzle_input) == expected
