//
//  Day2.swift
//  AOC
//
//  Created by Svetlin Mladenov on 2.12.23.
//

import Foundation

func Day2PartOne(_ input: String) {
    let games = input.split(whereSeparator: \.isNewline)
    
    var result = 0
    for game in games {
        let gameParts = game.split(separator: ":")
        let gameId = Int(gameParts[0].split(separator: " ")[1])!
        
        let sets = gameParts[1].split(separator: ";")
        
        var isGamePossible = true
        
        for set in sets {
            let cubes = set.split(separator: ",")
            var store = [
                "red": 0,
                "blue": 0,
                "green": 0
            ]
            
            for cube in cubes {
                let cubeParts = cube.split(separator: " ")
                let count = Int(String(cubeParts[0]))!
                let color = String(cubeParts[1])
                
                store[color]! += count
            }
            
            if store["red"]! > 12 || store["green"]! > 13 || store["blue"]! > 14 {
                isGamePossible = false
            }
        }
        
        if isGamePossible {
            result += gameId
        }
    }
    
    print("Result \(result)")
}

func Day2PartTwo(_ input: String) {
    let games = input.split(whereSeparator: \.isNewline)
    
    var result = 0
    for game in games {
        let gameParts = game.split(separator: ":")
        let sets = gameParts[1].split(separator: ";")
        
        var store = [
            "red": 0,
            "blue": 0,
            "green": 0
        ]
        
        for set in sets {
            let cubes = set.split(separator: ",")
            
            for cube in cubes {
                let cubeParts = cube.split(separator: " ")
                let count = Int(String(cubeParts[0]))!
                let color = String(cubeParts[1])
                
                if store[color]! < count {
                    store[color]! = count
                }
            }
            
        }
        
        // if 0 = 0, replace with 1
        result += store["red"]! * store["blue"]! * store["green"]!
    }
    
    print("Result \(result)")
}
