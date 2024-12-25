import pytest
from AOC import *

@pytest.mark.parametrize("puzzle_input,expected",
    [
        ("""#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####""", 3),
    ])
def test_part1(puzzle_input, expected):
    """Test part 1 on example input."""
    assert part1(puzzle_input) == expected
