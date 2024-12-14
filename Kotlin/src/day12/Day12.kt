package day12
import readInput

fun main() {
    data class Point(var row: Int, var col: Int)
    data class PointD(var row: Double, var col: Double)
    fun floodFill(grid: Map<Point, Char>, current: Point, letter: Char, seen: MutableSet<Point>): Pair<List<Point>, Int> {
        if (!grid.containsKey(current) || grid[current] != letter) {
            return Pair(emptyList(), 1)
        }
        if (current in seen) {
            return Pair(emptyList(), 0)
        }
        seen.add(current)
        val points = mutableListOf(current)
        var perimeter = 0

        val directions = listOf(Point(current.row, current.col + 1),
            Point(current.row, current.col - 1),
            Point(current.row + 1, current.col),
            Point(current.row - 1, current.col)
        )
        for (next in directions) {
            val result = floodFill(grid, next, letter, seen)
            points.addAll(result.first)
            perimeter += result.second
        }
        return Pair(points, perimeter)
    }
    fun part1(input: List<String>): Int {
        val grid  = input.withIndex().flatMap { (r, row) ->
            row.withIndex().map { (c, char) ->
                Point(r, c) to char
            }
        }.toMap()
        val seen = mutableSetOf<Point>()

        return input.flatMapIndexed { r, row ->
            row.mapIndexedNotNull { c, crop ->
                if (!seen.contains(Point(r, c))) {
                    val (area, perimeter) = floodFill(grid, Point(r, c), crop, seen)
                    area.size * perimeter
                } else {
                    null
                }
            }
        }.sum()
    }

    fun findSides(region: Set<Point>): Int {
        val possibilities = mutableSetOf<PointD>()
        var corners = 0
        for ((r, c) in region) {
            for (point in listOf(PointD(r - 0.5, c - 0.5), PointD(r + 0.5, c - 0.5), PointD(r + 0.5, c + 0.5), PointD(r - 0.5, c + 0.5))) {
                possibilities.add(point)
            }
        }

        for ((r, c) in possibilities) {
            val locations = mutableListOf<Int>()
            for (point in listOf(PointD(r - 0.5, c - 0.5), PointD(r + 0.5, c - 0.5), PointD(r + 0.5, c + 0.5), PointD(r - 0.5, c + 0.5))) {
                val p = Point(point.row.toInt(), point.col.toInt())
                if (p in region) locations.add(1)
                else locations.add(0)
            }
            val num = locations.sum()
            if (num == 1 || num == 3) corners++
            if (num == 2) {
                if (locations[0] == locations[2] || locations[1] == locations[3]) {
                    corners += 2
                }
            }
        }
        return corners
    }

    fun part2(input: List<String>): Int {
        val grid = input.withIndex().flatMap { (r, row) ->
            row.withIndex().map { (c, char) ->
                Point(r, c) to char
            }
        }.toMap()
        val seen = mutableSetOf<Point>()

        return input.flatMapIndexed { r, row ->
            row.mapIndexedNotNull { c, crop ->
                if (!seen.contains(Point(r, c))) {
                    val (area, _) = floodFill(grid, Point(r, c), crop, seen)
                    area.size * findSides(area.toSet())
                } else {
                    null
                }
            }
        }.sum()
    }

    val testInput = readInput("day12/test")
    check(part1(testInput) == 1930)
    check(part2(testInput) == 1206)

    val input = readInput("day12/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}