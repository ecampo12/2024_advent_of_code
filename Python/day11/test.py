import pytest
from AOC import *

@pytest.mark.parametrize("puzzle_input,expected,blinks",
    [
        ("125 17", 22, 6),
        ("125 17", 55312, 25),
    ])
def test_part1(puzzle_input, expected, blinks):
    """Test part 1 on example input."""
    assert part1(puzzle_input, blinks) == expected