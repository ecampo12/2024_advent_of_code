import pytest
from AOC import *

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ("2333133121414131402", 1928),
    ])
def test_part1(puzzle_input, expected):
    """Test part 1 on example input."""
    assert part1(puzzle_input) == expected

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ("2333133121414131402", 2858),
    ])
def test_part2(puzzle_input, expected):
    """Test part 2 on example input."""
    assert part2(puzzle_input) == expected
