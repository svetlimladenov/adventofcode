def main():
    input = ""
    with open("./in.txt", "r") as file:
        input = file.read().rstrip()

    data = input.split()
    first_column, second_column = sorted(data[0::2]), sorted(data[1::2])

    print(f"First: {partOne(first_column, second_column)}")
    print(f"Second: {partTwo(first_column, second_column)}")


def partOne(first_col: list, second_col: list) -> int:
    res = 0

    for i, num in enumerate(first_col):
        res += abs(int(num) - int(second_col[i]))

    return res


def partTwo(first: list, second: list) -> int:
    res = 0
    for num in first:
        count = 0
        for second_num in second:
            if num == second_num:
                count += 1
        res += count * int(num)

    return res


if __name__ == "__main__":
    main()
