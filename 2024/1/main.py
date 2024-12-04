def main():
    input = ""
    with open("./in.txt", "r") as file:
        input = file.read().rstrip()

    firstColumn = []
    secondColumn = []

    lines = input.split("\n")
    for i in lines:
        a, b = i.split("   ")
        firstColumn.append(a)
        secondColumn.append(b)

    print(f"Part : {partOne(firstColumn, secondColumn)}")
    print(f"second: {partTwo(firstColumn, secondColumn)}")


def partOne(firstColumn: list, secondColumn: list) -> int:
    res = 0
    for _ in range(0, len(firstColumn)):
        maxIndexFirst = firstColumn.index(min(firstColumn))
        maxIndexSecond = secondColumn.index(min(secondColumn))

        maxFirst = firstColumn.pop(maxIndexFirst)
        maxSecond = secondColumn.pop(maxIndexSecond)

        diff = abs(int(maxFirst) - int(maxSecond))
        res += diff

    return res


def partTwo(first: list, second: list) -> int:
    res = 0
    for num in first:
        count = 0
        for secondNum in second:
            if num == secondNum:
                count += 1
        res += count * int(num)

    return res


if __name__ == "__main__":
    main()
