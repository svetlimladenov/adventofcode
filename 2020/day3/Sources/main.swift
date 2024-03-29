import Foundation

let input = """
...........#..#.#.###....#.....
...#..#...........#.#...#......
#.....#..#........#...#..##....
..#...##.#.....#.............#.
#.#..#......#.....#....#.......
.....#......#..#....#.....#....
.......#.#..............#......
.....#...#..........##...#.....
...#....#.#...#.#........#...#.
..#.........###.......##...##..
.#....#...........#........#..#
..#.............##.............
..#.##.#....#................#.
.....##.#.......#....#...#.....
......#.#....##................
..#..........###..#..#.#..#....
....#..............#....##..#.#
.#.........#.#....#.#.#....#...
..#.....#......##.#....#.......
..#.#....#..#.#...##....###....
...#......##...#........#.#..#.
.##.#.......##....#............
...##..#.#............#...#.#..
.##...##.#..#..................
..#......##......#......##.....
.....##...#..#...#.........#...
.##.#.....#..#..#.##....##....#
..#.#......#.......##..........
......................#......##
##.#...#.................#.#.#.
......#.#..........#.....##.#..
#.#......#.....#...........#...
.....#...#.......#..#..#.#...#.
...........#......#.#...#......
....##...##...........#......#.
.........#.##..................
......#...#....#......##.##...#
......#...#.#########......#...
.......#.#...#.......#..#......
............#...#...#.###...##.
...........#..........#...#....
...#..#.#................#.#..#
..#....#.....#.............#.#.
....##......#........#....#....
........##...............#....#
........#..#...#..............#
...#....#.#...#..#...#....#.#.#
.........#.......#....##.......
#.##..............#.#........##
......................###......
.........#....##..##....#.#.#..
.#...##.#.#...#....##..#.....#.
....................#.#......#.
.#..#.......................#..
..#..#.............#..#....#...
...#...#..#...#...#.#..#.##....
........#....#...#....#........
.#.....#........#.........#...#
...#......#..#..#..####..#.....
#....#..............#.##.......
.#....#.............##...#.....
....#...#.##........##......#..
##....#...#.......#..#........#
....##........................#
..................#..#.........
..#....#........#..#.......#...
#...#..#....#...##...........#.
.........#..#..#.#.##..........
....#.#..#.....#..#.#.#.#..#.##
##................#.##.....#...
.#.....###..#..#..#.....#....##
...#...........#..........####.
.#.....#....#......#.##..#.#...
..#...##....#................#.
........#.......#......#.#.....
....#.#.#.#....#...#......#..#.
...........#......#..#.........
###...##......##.#..#....##....
##....##.........#..#....###...
#.#.....#....#......#...#..##..
#....##.#..............#.##....
.#........#.#.........#...#....
......................#......#.
........#..#..##.....#..#.#....
..#...###.................#..#.
...#...#............#..........
.##.......#..#.........#....#..
.#..............#....#....##...
...............##..#.#.......##
.#.....#....#...#..#.......#..#
#..#.............#....#......#.
.....#.#......#.........###..#.
.#...#.#...............#....#..
#......#.............#.........
.#.##.#.####...#..#.##..#.....#
.....#......#..#...#.......#...
#........###...#.....#..#.....#
....#.#.....#...#..........#...
...#.#.......#.#......#..#.##..
..#..........#.#..#.......#.#..
#...#.#..............#...###.#.
...#..#...#............###.....
..#..#...#.#............#..#...
.###.#.....................#..#
....#....#..#.....##.##........
#....#....#.#..#.........#.....
.#.....##....#............##..#
#....#.#...#...#..#.#......#...
#.....##.....##.....##.#...##..
...##...#..#..####.#........#..
.........#...#......##..##..#..
..#.....###.#..#..####.#.......
.......#......#...#..##....#...
.#.....#.#.#..#....#...#..##...
..........#.#...#...#.#..#.....
....#.....#........#.....##..#.
..#.#.##.........#...##.....##.
.........#.##....#............#
............##.....#.......#.#.
......#.....#...#..#..###......
##.....#.......#...##.#....#...
...........##......#..##...#.#.
..#.#.#.#...#.......#....#...#.
#.............#.....#.#...###..
##....#.......#.....#..##.#....
...#.......#....#.........##...
......#.......#......##.##.....
..#......#...#.#........#......
....#.#....#.#.##......#.....#.
#......#.........#..#....#.....
........#..#....##.....#.......
#......##....#.##....#......#..
..#.......#............##.....#
...........#...#...........#...
#.......#...#....#....#...#.#.#
..###..#........#........#.....
..#..###...........#.#......#..
.#...........#......#..........
.#.......#.....#.......#.......
..#......##.#............#....#
#..........#.....#.#..#........
.....#...##.##.......#......#..
..........#.....#........#.#.#.
....#......#.....#......#.#....
.........#.#.#..#...##....#...#
.........#.......#...##...#.#..
.##........#...............#...
.......#....#...........##.....
.........###......#.........#.#
......#.......#...#..........#.
...#.#..........##......#...#..
#.......#.#..........#.........
................#..#......#..##
.....#...#....#.#.....#........
#.....#....#...........#....#..
#....#.#..#...##....#...##.#...
...#.....#......#.#....#..#..#.
..#................#...#.#..##.
..........#..............#..#.#
.....##.....#..#.###...........
....#.#......#.#...........#...
.#....#.#.........##.#....#....
.#.#........#........###....#..
##.#................#...#..#...
.......#......##..#.....#..#.#.
...#............#......###...##
#.#...........#.........#......
.....#.#.#.................#...
....#..............#...#.#.....
...#.#.....##..#...............
.#..##...#....##.....###..#....
...............#...#...#.#.###.
.###....#.....#...#.#......#...
...#..#.....#.......#..##.#....
...........#..#....##..#...#...
...#...#..........#.......##.#.
............#.#.......#........
....#.........#.....#..........
...#.###.##..#...##..####..#..#
.#.#...#..#...................#
.....#..#.....##..#............
....#......#...##..#.##........
...#...............#..#.....##.
...#......#.........#.#..##....
.#....#.##.......#......#......
....#.......#....#..........#..
#.#.#....###.#.#.............#.
..##..###........#.#..#.#..#...
......#.#............##.#...###
.........#.#....#####.....##...
............##......#.#..#.....
...#.....#.....###....#........
##..........####.##.#.#........
....................##.....##.#
#.#............#........#......
....#...##.....#......#....#...
...###.#..##..................#
..###......#..............#.#.#
.#...#...........#....#........
....#............#..#..#.#.....
...#.........#.#.........#.####
..#...#...#...#...........#....
...............#.#......##..#..
#....#.#.......#.#..#......#..#
........#.#....#..#...#..#...#.
...#..#.......#...........#....
...........#.......#...........
.#......#................#.....
....#.#.#...#......#..#...#....
................#.#.#....#.....
.........................##..#.
.#...........#............##...
#...............#.....##.#.#..#
.........#.......###....#.....#
....##...#...#.....#..#........
........#.....#..#.#.#...#..#..
......#.......#.#.........#.#..
#......#............#...#....#.
#..##...#..#................#..
.##...#...#.....#.##.......#..#
.......#.##........##..##......
##.#..##...............#.....#.
......#....#..##...#......###.#
#........##..#....#.#......#...
.#......##.#...#.#...#.........
.#.#...#..#.............#......
.##..........#..........#......
.#.....#.....#..............#.#
..#.........#..#.#.....#.#....#
..#.##..............##...#..###
....................#..........
......###..#..#...........#....
..#..........#.......#...#.....
...#......#......#.............
....##..............#.#.....#..
........#.#......#..#........##
.............#...#.#.........##
...###...#..........##.......#.
.#..........#...##..#.#.....#..
##...#.........#...............
......#....#....#.....#.....#..
..........#....#...#...#..#...#
...##....#.#.#..#...##.........
#......#.#...##.###...#....#...
##.......##.#......##..#...#...
......#.............#.##.....##
#.......###....####.#...##....#
..#...#..#.......#..........#..
#.....#..#..#..#.##...###...#..
.....##.#..#..#..#.#....#...#..
..#...#..................##....
....#.#........##..............
#...#.......##...#...#.#.......
..#...#........##....#.#.......
..........###...###...#......#.
#.....#..###...##...##..#..#..#
..#.....##.....#.......##..#.#.
........#........#.........#...
.................#....#.......#
.......#...#.....#...#.#.......
....##...............#...##...#
.##...#................#...#...
.............#.................
.#..#....#....#.#....#.........
.#.#..#..........#.......#.....
.....##.....##...#..#..........
#...#.#.........#......#..#....
........#....#...#....#.#.##...
....#..#........#...#...#......
.#..#.....#.#...#.........#....
.#..#..#...........#..#....#...
....###.............#..#.......
#......#..#..##..........#.#...
#..#..#.##..#...#.#.#..........
....###......#.##.....#....#...
.##..#...#......##.#...........
..#..#.......#.....#.##....#.#.
.......#.#.#........#....##....
..##...#....#...............###
#..##..#...........#.#....##...
...##..#.....................#.
###......#....#....###..#...##.
.........##............#..#...#
..#..........#...#.#.#......#.#
.......#.....##..##......#.##..
#..........#.....##.#..........
#.......#.#...#...#....#.......
#...#.....##.......#.#..#.#.#..
.........#.#.#..#..#...#.###...
.................##...#....#...
###.......#..........##...#....
#.#..#.........#....##.#.......
......#.#.....#........#.......
.......#.#........#......#.#..#
..............#..#...##....#..#
#...........#...##.....#..#.#..
..#....#..#.#.#...#..#....#.#..
...##.#.....#..#...##..#.....#.
..#.#................#........#
......#...#.............#......
.##............#....#...#..#...
....#...#...........#.......#..
.###..#.......#.............#.#
.#.#....#.#...........#.#......
...#.........#.........#..#....
...#..........#.#.....#.#......
.....#........#....##......#...
..#.#.#......#..#.#......#....#
.#.#..#................#.#.....
.#.#.........##...#.......#.#.#
#..#.....#...#..#...........#..
..##......####......#..#....###
#.....###....#.#........#..#..#
..##.#...#.#..##..........#..#.
#.........#.#.............#...#
...#.#...#...#.#.#....##.......
##.##...#.....#...#...........#
....#........#.#.....#.........
.................##..#..##...##
.....##....#...#...#.....#..#..
....#...#........#............#
..#...........##....#...#...##.
.....#......#.........#..##.#..
"""

let rows = input.split(whereSeparator: \.isNewline)

func getCurrentRow(_ row: Substring, _ n: Int) -> Substring {
    var result = ""
    for _ in 0..<n{
       result += row
    }
    
    return Substring(result)
}

let slopes = [
    [1,1],
    [3,1],
    [5,1],
    [7,1],
    [1,2]
]

var result = 1
for slope in slopes {
    let right = slope[0]
    let down = slope[1]
    
    print("Right \(right), Down \(down)")
    
    var treesEncountered = 0
    
    var currentCol = 0
    var rowLenght = rows[0].count
    let treeSymbol = "#"
    var currentRow = rows[0]
    
    for i in stride(from: 1, to: rows.count, by: down){
        var row = rows[i]
        
        currentCol += right
        
        if currentCol >= rowLenght {
            rowLenght = currentRow.count * 2
        }
        
        currentRow = getCurrentRow(row, rowLenght / row.count)
        
        let currentIndex = currentRow.index(currentRow.startIndex, offsetBy: currentCol)
        let symbol = currentRow[currentIndex]
        
        if String(symbol) == treeSymbol {
            treesEncountered += 1
        }
    }
    
    print("res \(treesEncountered)")
    result *= treesEncountered
}

print("Final reuslt \(result)")
// 82, 242, 71, 67, 26
