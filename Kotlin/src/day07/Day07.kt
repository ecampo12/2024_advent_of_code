package day07
import readInput
import kotlin.time.measureTimedValue

fun main() {
    fun canEvaluate(target: Long, nums:List<Long>, part2:Boolean = false): Boolean {
        if (nums.size ==1 ) return target == nums[0]
        if (target % nums.last() == 0L && canEvaluate(target / nums.last(), nums.subList(0, nums.size - 1), part2)) return true
        if (target > nums.last() && canEvaluate(target - nums.last(), nums.subList(0, nums.size - 1), part2)) return true

        if (part2) {
            val strTarget = target.toString()
            val strLast = nums.last().toString()
            if (strTarget.endsWith(strLast) && strTarget.length > strLast.length) {
                val subStr = strTarget.substring(0, strTarget.length-strLast.length).toLong()
                return (canEvaluate(subStr, nums.subList(0, nums.size - 1), true))
            }
        }

        return false
    }

    fun part1(input: List<String>): Long {
        return input.sumOf { line ->
            val x = """(\d+)""".toRegex().findAll(line).toList().map { it.value.toLong() }
            if (canEvaluate(x[0], x.subList(1, x.size))) x[0] else 0
        }
    }

    fun part2(input: List<String>): Long {
        return input.sumOf { line ->
            val x = """(\d+)""".toRegex().findAll(line).toList().map { it.value.toLong() }
            if (canEvaluate(x[0], x.subList(1, x.size), true)) x[0] else 0
        }
    }

    val testInput = readInput("day07/test")
    check(part1(testInput) == 3749L)
    check(part2(testInput) == 11387L)

    val input = readInput("day07/input")
    val (ans1, time1) = measureTimedValue { part1(input) }
    println("part1: $ans1")
    println("Took Part1: $time1\n")
    val (ans2, time2) = measureTimedValue { part2(input) }
    println("Part2: $ans2")
    println("Took Part2: $time2")
}