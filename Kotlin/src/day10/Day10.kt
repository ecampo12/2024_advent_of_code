package day10
import readInput
import java.util.*

fun main() {
    data class Point(val row: Int, val col: Int)
    fun scoreTrailhead(grid: List<String>, point: Point): Int {
        val queue = ArrayDeque<Point>()
        queue.add(point)
        val seen = mutableSetOf<Point>()
        var summits = 0
        val directions = mutableListOf(
            Point(1,0),
            Point(-1,0),
            Point(0,1),
            Point(0,-1),
        )
        while (queue.isNotEmpty()) {
            val curr = queue.removeFirst()
            for (dir in directions) {
                val value = grid[curr.row][curr.col].digitToInt()
                val newCurr = Point(curr.row + dir.row, curr.col + dir.col)
                if (newCurr.row in grid.indices && newCurr.col in grid[0].indices) {
                    if (grid[newCurr.row][newCurr.col].digitToInt() != value + 1) continue
                    if (newCurr in seen) continue
                    if (grid[newCurr.row][newCurr.col] == '9') summits++
                    else queue.add(newCurr)
                    seen.add(newCurr)
                }
            }
        }
        return summits
    }

    fun part1(input: List<String>): Int {
        return input.withIndex().flatMap { (r, row) ->
            row.withIndex().mapNotNull { (c, col) ->
                if (col == '0') Point(r, c)
                else null
            }
        }.sumOf {
            scoreTrailhead(input, it)
        }
    }
    fun rateTrailhead(grid: List<String>, point: Point): Int {
        val queue = ArrayDeque<Point>()
        queue.add(point)
        val seen = mutableMapOf(point to 1)
        var trails = 0
        val directions = mutableListOf(
            Point(1,0),
            Point(-1,0),
            Point(0,1),
            Point(0,-1),
        )
        while (queue.isNotEmpty()) {
            val curr = queue.removeFirst()
            if (grid[curr.row][curr.col] == '9') trails += seen[curr]!!
            for (dir in directions) {
                val value = grid[curr.row][curr.col].digitToInt()
                val newCurr = Point(curr.row + dir.row, curr.col + dir.col)
                if (newCurr.row in grid.indices && newCurr.col in grid[0].indices) {
                    if (grid[newCurr.row][newCurr.col].digitToInt() != value + 1) continue
                    if (newCurr in seen) {
                        seen[newCurr] = seen[newCurr]!! + seen[curr]!!
                        continue
                    }
                    seen[newCurr] = seen[curr]!!
                    queue.add(newCurr)
                }
            }
        }
        return trails
    }

    fun part2(input: List<String>): Int {
        return input.withIndex().flatMap { (r, row) ->
            row.withIndex().mapNotNull { (c, col) ->
                if (col == '0') Point(r, c)
                else null
            }
        }.sumOf {
            rateTrailhead(input, it)
        }
    }

    val testInput = readInput("day10/test")
    check(part1(testInput) == 36)
    check(part2(testInput) == 81)

    val input = readInput("day10/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}