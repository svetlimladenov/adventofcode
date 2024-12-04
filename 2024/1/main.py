def main():
    input = ""
    with open("./in.txt", "r") as file:
        input = file.read().rstrip()

    data = input.split()
    firstColumn, secondColumn = sorted(data[0::2]), sorted(data[1::2])

    print(f"First: {partOne(firstColumn, secondColumn)}")
    print(f"Second: {partTwo(firstColumn, secondColumn)}")


def partOne(firstColumn: list, secondColumn: list) -> int:
    res = 0

    for i, num in enumerate(firstColumn):
        res += abs(int(num) - int(secondColumn[i]))

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
