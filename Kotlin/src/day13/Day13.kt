package day13
import kotlin.io.path.Path
import kotlin.io.path.readText

fun main() {
    fun solveMachine(machine: List<Long>, error: Long): Long {
        val ax = machine[0]
        val ay = machine[1]
        val bx = machine[2]
        val by = machine[3]
        val px = machine[4] + error
        val py = machine[5] + error

        val ca = (px * by - py * bx) / (ax * by - ay * bx)
        val cb = (px - ax * ca) / bx
        return if (ax * ca + bx * cb == px && ay * ca + by * cb == py){
            3 * ca + cb
        } else 0
    }
    fun part1(input: String): Long {
        return input.split("\n\n").sumOf { machine ->
            solveMachine(
                """(\d+)""".toRegex().findAll(machine).map { it.value.toLong() }.toList(), 0
            )
        }
    }

    fun part2(input: String): Long {
        return input.split("\n\n").sumOf { machine ->
           solveMachine(
               """(\d+)""".toRegex().findAll(machine).map { it.value.toLong() }.toList(), 10000000000000
           )
        }
    }

    val testInput = Path("src/day13/test.txt").readText().trimIndent()
    check(part1(testInput) == 480L)

    val input = Path("src/day13/input.txt").readText().trimIndent()
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}