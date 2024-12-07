package day06
import readInput

fun main() {
    fun getStart(grid: List<String>): Pair<Int, Int>? {
        for (row in grid.indices) {
            val col = grid[row].indexOf('^')
            if (col != -1) return row to col
        }
        return null
    }

    fun traverse(grid: List<String>, start: Pair<Int, Int>): Set<Pair<Int, Int>> {
        var dr = -1
        var dc = 0
        val seen = mutableSetOf<Pair<Int, Int>>()
        var r = start.first
        var c = start.second
        while (true) {
            seen.add(Pair(r, c))
            if ( r+dr < 0 ||r +dr >= grid.size || c+dc <0 || c+dc >= grid[0].length) break
            if (grid[r+dr][c+dc] == '#') {
                dc = -dr.also { dr = dc }
            } else {
                r += dr
                c += dc
            }
        }
        return seen
    }

    fun part1(input: List<String>): Int {
        val start = getStart(input)
        return traverse(input, start!!).size
    }

    data class Position(val row: Int, val col: Int, val dr: Int, val dc: Int)
    fun findLoop(grid: List<String>, start: Pair<Int, Int>, obstruction: Pair<Int,Int>): Boolean {
        var dr = -1
        var dc = 0
        val seen = mutableSetOf<Position>()
        var r = start.first
        var c = start.second
        while (true) {
            seen.add(Position(r, c, dr, dc))
            if ( r+dr < 0 ||r +dr >= grid.size || c+dc <0 || c+dc >= grid[0].length) return false
            if (grid[r+dr][c+dc] == '#' || Pair(r+dr, c+dc) == obstruction) {
                dc = -dr.also { dr = dc }
            } else {
                r += dr
                c += dc
            }
            if (Position(r, c, dr, dc) in seen){
                return true
            }
        }
    }

    fun part2(input: List<String>): Int {
        val start = getStart(input)
        return traverse(input, start!!).count{ path -> findLoop(input, start, path) }
    }

    val testInput = readInput("day06/test")
    check(part1(testInput) == 41)
    check(part2(testInput) == 6)

    val input = readInput("day06/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}