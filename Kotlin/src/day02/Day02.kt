package day02
import readInput

fun main() {
    fun isSafe(report: List<Int>): Boolean {
        val diff = (0..<report.lastIndex).map { report[it] - report[it +1] }
        return diff.all { intArrayOf(1, 2, 3).contains(it) } or diff.all { intArrayOf(-1, -2,-3).contains(it) }
    }

    fun part1(input: List<String>): Int {
        val reports = input.map{ s -> s.split(" ").map{it.toInt()} }
        return reports.count { isSafe(it) }
    }

    fun part2(input: List<String>): Int {
        val reports = input.map{ s -> s.split(" ").map{it.toInt()} }
        return reports.count { report ->
            isSafe(report) || report.indices.any { i ->
                isSafe(report.filterIndexed { index, _ -> index != i })
            }
        }
    }

    val testInput = readInput("day02/test")
    check(part1(testInput) == 2)
    check(part2(testInput) == 4)

    val input = readInput("day02/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}
