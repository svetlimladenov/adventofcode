//
//  Day3.swift
//  AOC
//
//  Created by Svetlin Mladenov on 3.12.23.
//

import Foundation

func Day3(_ input: String) -> Int {
    let lines = input.split(whereSeparator: \.isNewline)
    
    for (col, line) in lines.enumerated() {
        let symbols = line.split(separator: "")
        
        var foundANumber = false
        var shouldCheckIfValid = false
        var number = ""
        for (row, char) in symbols.enumerated() {`
            if let num = Int(String(char)) {
                foundANumber = true
                number += String(char)
            } else if foundANumber == true {
                foundANumber = false
                number = ""
            }
        }
        
    }
    
    return 0
}

