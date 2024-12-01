import pytest
from AOC import *

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ([
            "3   4",
            "4   3",
            "2   5",
            "1   3",
            "3   9",
            "3   3"
            ], 11),
    ])
def test_part1(puzzle_input, expected):
    """Test part 1 on example input."""
    assert part1(puzzle_input) == expected

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ([
            "3   4",
            "4   3",
            "2   5",
            "1   3",
            "3   9",
            "3   3"
            ], 31),
    ])
def test_part2(puzzle_input, expected):
    """Test part 2 on example input."""
    assert part2(puzzle_input) == expected
