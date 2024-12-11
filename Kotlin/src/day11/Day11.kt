package day11
import readInput

fun main() {
    val cache = mutableMapOf<Pair<Long,Int>, Long>()
    fun countStones(stone: Long, steps: Int): Long {
        if (steps == 0) return 1
        if (stone == 0L) return countStones(1, steps - 1)
        if (Pair(stone, steps) in cache) return cache[Pair(stone, steps)]!!

        val stoneStr = stone.toString()
        val length = stoneStr.length

        if (length % 2 == 0) {
            val left = stoneStr.substring(0, length/ 2).toLong()
            val right = stoneStr.substring(length/ 2).toLong()
            cache[Pair(stone, steps)] = countStones(left, steps - 1) + countStones(right, steps - 1)
            return cache[Pair(stone, steps)]!!
        }
        return countStones(stone * 2024, steps - 1)
    }

    fun part1(input: List<String>, blinks: Int=25): Long {
        return input[0].split(" ")
            .map(String::toLong)
            .sumOf { countStones(it, blinks) }
    }

    fun part2(input: List<String>): Long {
        return part1(input, 75)
    }

    val testInput = readInput("day11/test")
    check(part1(testInput, 6) == 22L)
    check(part1(testInput, 25) == 55312L)

    val input = readInput("day11/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}