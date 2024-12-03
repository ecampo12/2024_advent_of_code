package day03

import readInput

fun main() {
    fun part1(input: List<String>): Int {
        return """mul\((\d+),(\d+)\)""".toRegex()
            .findAll(input.joinToString(""))
            .sumOf { matchResult ->
                val (a, b) = matchResult.destructured
                a.toInt() * b.toInt()
            }
    }

    fun part2(input: List<String>): Int {
        return part1(
            listOf(
                input.joinToString("")
                    .replace("""don't\([^)]*\).*?(?=do\([^)]*\)|${'$'})""".toRegex(), "")
            )
        )
    }

    val testInput = readInput("day03/test")
    check(part1(testInput) == 161)
    check(part2(testInput) == 48)

    val input = readInput("day03/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}
