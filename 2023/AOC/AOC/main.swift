//
//  main.swift
//  AOC
//
//  Created by Svetlin Mladenov on 1.12.23.
//

import Foundation

var solutions : [Int: (String) -> Int] = [:]

solutions[1] = Day1Reverse
solutions[2] = Day2PartTwo

func Solve(_ day: Int) {
    let input = ReadInput(day)
    let result = solutions[day]!(input)
    print("Result from day \(day) - \(result)")
}

func ReadInput(_ day: Int) -> String {
    let currentDirectoryPath = FileManager.default.currentDirectoryPath
    let inputUrl = URL(filePath: "\(currentDirectoryPath)/Inputs/\(day)")

    var input = ""
    do {
        input = try String(contentsOf: inputUrl)
    } catch {
        print("Input for Day \(day) not found \(error)")
    }
    
    return input
}


Solve(1)

