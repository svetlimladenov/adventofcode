//
//  Day1.swift
//  AOC
//
//  Created by Svetlin Mladenov on 1.12.23.
//

import Foundation

var numberAsStrings = [
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9
]

func Day1WithDictionary(_ input: String) {
    let lines = input.split(whereSeparator: \.isNewline)
    
    var sum = 0
    for line in lines {
        let numbers = getStringNumbers(String(line))
        let sorted = Array(numbers.keys).sorted(by: <)
        let firstNumber = numbers[sorted[0]]!
        let lastNumber = numbers[sorted[sorted.count - 1]]!
        sum += Int("\(firstNumber)\(lastNumber)")!
    }
    
    print(sum)
}

func getStringNumbers(_ line: String) -> [Int: Int] {
    var ranges: [Int: Int] = [:]
    for num in numberAsStrings.keys {
        var startIndex = line.startIndex
        while startIndex < line.endIndex {
            if let range = line[startIndex...].range(of: num) {
                let index = line.distance(from: line.startIndex, to: range.lowerBound)
                ranges[index] = numberAsStrings[num]
                startIndex = range.upperBound
            } else {
                break
            }
        }
    }
    
    for (index, l) in line.enumerated() {
        if l.isNumber {
            ranges[index] = Int(String(l))!
        }
    }
    
    return ranges
}

func findNumber(_ line: String) -> Int? {
    for num in numberAsStrings {
        if line.starts(with: num.key) {
            return num.value
        }
    }
    
    return nil
}

func getStringNumbersReversed(_ line: String) -> Int {
    var firstNumber = 0
    // firstNumber
    for (index, l) in line.enumerated() {
        if l.isNumber {
            firstNumber = Int(String(l))!
            break
        }
        
        let currentLine = String(line.dropFirst(index))
        if let num = findNumber(currentLine) {
            firstNumber = num
            break
        }
    }
    
    // lastNumber
    var lastNumber = firstNumber
    var index = line.index(before: line.endIndex)
    while index != line.startIndex {
        let l = line[index]
        if l.isNumber {
            lastNumber = Int(String(l))!
            break
        }
        
        let currentLine = String(line[index...])
        if let num = findNumber(currentLine) {
            lastNumber = num
            break
        }
        
        index = line.index(before: index)
    }
    
    return Int("\(firstNumber)\(lastNumber)")!
}

func Day1Reverse(_ input: String) -> Int {
    let lines = input.split(whereSeparator: \.isNewline)
    
    var sum = 0
    for line in lines {
        let n = getStringNumbersReversed(String(line))
        sum += n
    }
    
    return sum
}
