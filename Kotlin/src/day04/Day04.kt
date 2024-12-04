package day04
import readInput

fun main() {
    fun part1(input: List<String>): Int {
        val rows = input.size
        val columns = input[0].length
        var count = 0
        val directions = listOf(
            Pair(-1, 0), Pair(1, 1), Pair(1, 0), Pair(1, -1),
            Pair(0, -1), Pair(-1, -1), Pair(-1, 1), Pair(0, 1)
        )
        for (r in 0 until rows) {
            for (c in 0 until columns) {
                if (input[r][c] != 'X') continue
                for ((dr, dc) in directions){
                    if (r + 3 * dr in 0 until rows && c + 3 * dc in 0 until columns) {
                        val builder = StringBuilder()
                        repeat(3) { step ->
                            builder.append(input[r + (step + 1) * dr] [c + (step + 1) * dc])
                        }
                        if (builder.toString() == "MAS"){
                            count++
                        }
                    }
                }
            }
        }
        return count
    }

    fun part2(input: List<String>): Int {
        val rows = input.size
        val columns = input[0].length
        var count = 0
        val validCorners = setOf("MMSS", "MSSM", "SSMM", "SMMS")
        for (r in 1 until rows - 1) {
            for (c in 1 until columns - 1) {
                if (input[r][c] == 'A'){
                    val corners = buildString {
                        append(input[r - 1][c - 1])
                        append(input[r - 1][c + 1])
                        append(input[r + 1][c + 1])
                        append(input[r + 1][c - 1])
                    }
                    if (corners in validCorners) count++
                }
            }
        }
        return count
    }

    val testInput = readInput("day04/test")
    check(part1(testInput) == 18)
    check(part2(testInput) == 9)

    val input = readInput("day04/input")
    println("part1: ${part1(input)}")
    println("part2: ${part2(input)}")
}