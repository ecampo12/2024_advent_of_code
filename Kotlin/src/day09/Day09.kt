package day09
import readInput

fun main() {
    fun part1(input: String): Long {
        val disk = mutableListOf<Long>()
        var id = 0L
        for ((i, char) in input.withIndex()) {
            val size = char.digitToInt()
            if (i % 2 == 0) {
                disk += List(size) {id}
                id++
            } else {
                disk += List(size) {-1L}
            }
        }

        val freeSpace = disk.withIndex().filter {it.value == -1L}.map {it.index}
        for (i in freeSpace) {
            while (disk.last() == -1L) disk.removeLast()
            if (disk.size <= i) break
            disk[i] = disk.removeLast()
        }
        return disk.indices.sumOf { disk[it] * it }
    }

    data class Block(val pos: Int, val size: Int)
    fun part2(input: String): Long {
        val files = mutableMapOf<Int, Block>()
        var freeSpace = mutableListOf<Block>()
        var id = 0
        var pos = 0

        for ((i, char) in input.withIndex()) {
            val size = char.digitToInt()
            if (i % 2 == 0) {
                files[id++] = Block(pos, size)
            } else {
                freeSpace.add(Block(pos, size))
            }
            pos += size
        }

        for (fid in (id -1) downTo 0) {
            val file = files[fid]!!
            val updateFreeSpace = mutableListOf<Block>()
            var allocated = false

            for (space in freeSpace) {
                if (allocated || space.pos >= file.pos) {
                    updateFreeSpace.add(space)
                    continue
                }

                if (file.size <= space.size) {
                    files[fid] = Block(space.pos, file.size)
                    allocated = true

                    if (file.size < space.size) {
                        updateFreeSpace.add(Block(space.pos + file.size, space.size - file.size))
                    }
                } else {
                    updateFreeSpace.add(space)
                }
            }

            freeSpace = updateFreeSpace
        }

        return files.entries.sumOf { (x, block) ->
            (block.pos until block.pos + block.size).sumOf { j -> x * j.toLong() }
        }
    }

    val testInput = readInput("day09/test")
    check(part1(testInput[0]) == 1928L)
    check(part2(testInput[0]) == 2858L)

    val input = readInput("day09/input")
    println("part1: ${part1(input[0])}")
    println("part2: ${part2(input[0])}")
}