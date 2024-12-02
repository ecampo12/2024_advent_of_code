package day01
import readInput
import kotlin.math.abs

fun main() {
    fun part1(input: List<String>): Int {
        val (list1, list2) = input
            .map { it.split("   ").let { parts -> parts[0].toInt() to parts[1].toInt() } }
            .unzip()
            .let { (l1, l2) -> l1.sorted() to l2.sorted() }

        return list1.indices.sumOf { i -> abs(list1[i] - list2[i]) }
    }

    fun part2(input: List<String>): Int {
        val (list1, list2) = input
            .map { it.split("   ").let { parts -> parts[0].toInt() to parts[1].toInt() } }
            .unzip()
        return list1.sumOf { i -> i * list2.count { it == i } }
    }

    val testInput = readInput("day01/test")
    check(part1(testInput) == 11)
    check(part2(testInput) == 31)

    val input = readInput("day01/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}
