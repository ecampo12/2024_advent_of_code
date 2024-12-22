import pytest
from AOC import *

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ([
            "1",
            "10",
            "100",
            "2024"
        ], 37327623)
    ])
def test_part1(puzzle_input, expected):
    """Test part 1 on example input."""
    assert part1(puzzle_input) == expected

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ([
            "1",
            "2",
            "3",
            "2024"
        ], 23)
    ])
def test_part2(puzzle_input, expected):
    """Test part 2 on example input."""
    assert part2(puzzle_input) == expected
