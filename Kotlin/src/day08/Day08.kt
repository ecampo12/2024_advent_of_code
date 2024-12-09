package day08
import readInput
import kotlin.math.max

fun main() {
    data class Point(val row: Int, val col: Int)

    fun parseInput(input: List<String>): HashMap<Char, List<Point>> {
        val points = HashMap<Char, List<Point>>()
        input.indices.forEach { row ->
            input[row].indices.forEach { col ->
                val character = input[row][col]
                if (character.isLetterOrDigit()) {
                    if (!points.containsKey(character)) {
                        points[character] = mutableListOf()
                    }
                    points[character] = points[character]!!.plus(Point(row, col))
                }
            }
        }
        return points
    }

    fun combinations(elements: List<Point>): Sequence<Pair<Point, Point>> {
        return sequence {
            for (i in elements.indices) {
                for (j in i + 1 until elements.size) {
                    yield(Pair(elements[i], elements[j]))
                }
            }
        }
    }

    fun getPoint(a: Point, b: Point): List<Point> {
        val changeRow = a.row - b.row
        val changeCol = a.col - b.col
        val p1 = Point(a.row + changeRow, a.col + changeCol)
        val p2 = Point(b.row - changeRow, b.col - changeCol)
        return mutableListOf(p1, p2)
    }

    fun part1(input: List<String>): Int {
        val antennas = parseInput(input)
        val antinodes = mutableSetOf<Point>()
        for (p in antennas.values) {
            for (points in combinations(p)) {
                antinodes += getPoint(points.first, points.second)
            }
        }
        return antinodes.count { it.row in input.indices && it.col in 0..<input[0].length }
    }

    fun pointsOnLines(a: Point, b: Point, height: Int, width: Int): Set<Point> {
        val changeRow = a.row - b.row
        val changeCol = a.col - b.col
        val points = mutableSetOf<Point>()
        for (directions in listOf(-1, 1)) {
            repeat(max(height, width)) { i ->
                points.add(Point(a.row + directions * i * changeRow, a.col + directions * i * changeCol))
            }
        }
        return points
    }

    fun part2(input: List<String>): Int {
        val antennas = parseInput(input)
        val antinodes = mutableSetOf<Point>()
        for (p in antennas.values) {
            for (points in combinations(p)) {
                antinodes += pointsOnLines(points.first, points.second, input.size, input[0].length)
            }
        }
        return antinodes.count { it.row in input.indices && it.col in 0..<input[0].length }
    }

    val testInput = readInput("day08/test")
    check(part1(testInput) == 14)
    check(part2(testInput) == 34)

    val input = readInput("day08/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}