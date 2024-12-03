import pytest
from AOC import *

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))", 161),
    ])
def test_part1(puzzle_input, expected):
    """Test part 1 on example input."""
    assert part1(puzzle_input) == expected

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))", 48),
    ])
def test_part2(puzzle_input, expected):
    """Test part 2 on example input."""
    assert part2(puzzle_input) == expected
