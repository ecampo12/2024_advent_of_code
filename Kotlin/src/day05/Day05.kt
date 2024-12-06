package day05
import readInput
import java.util.*

data class Rule(val a: Int, val b: Int)
data class Validation(val isValid: Boolean, val violatedRule: Rule?)
fun main() {
    fun parseInput(input: List<String>): Pair<List<Rule>, List<List<Int>>>  {
        val (ruleSelection, updateSection) = input.partition { it.contains("|") }
        val rules = ruleSelection.map {
            val (a, b) = it.split("|").map(String::toInt)
            Rule(a, b)
        }
        val updates = updateSection.drop(1).map { line -> line.split(",").map(String::toInt) }
        return rules to updates
    }

    fun isValid(rules: List<Rule>, updates: List<Int>): Validation {
        val indexMap = updates.withIndex().associate { (index, value) -> value to index }
        for (rule in rules) {
            val (a, b) = indexMap[rule.a] to indexMap[rule.b]
            if (a != null && b != null && a > b) {
                return Validation(false, rule)
            }
        }
        return Validation(true, null)
    }

    fun part1(input: List<String>): Int {
        val (rules, updates) = parseInput(input)

        return updates.sumOf { update ->
            val validation = isValid(rules, update)
            if (validation.isValid) update[update.size / 2] else 0
        }
    }

    fun correctUpdate(rule: Rule, update: List<Int>): List<Int> {
        val (a, b) = rule
        val indexA = update.indexOf(a)
        val indexB = update.indexOf(b)
        return update.apply { Collections.swap(this, indexA, indexB) }
    }

    fun part2(input: List<String>): Int {
        val (rules, updates) = parseInput(input)
        return updates.filterNot { isValid(rules, it).isValid }
            .sumOf { update ->
            var current = update
            while (true) {
                val v = isValid(rules, current)
                if (v.isValid) break
                current = correctUpdate(v.violatedRule!!, current)
            }
            current[current.size / 2]
        }
    }

    val testInput = readInput("day05/test")
    check(part1(testInput) == 143)
    check(part2(testInput) == 123)

    val input = readInput("day05/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}