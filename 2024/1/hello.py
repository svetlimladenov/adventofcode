
def main():
    input = """3   4
4   3
2   5
1   3
3   9
3   3"""

    firstColumn = []
    secondColumn = []

    lines = input.split("\n")
    for _, i in enumerate(lines):
        a, b = i.split("   ")
        firstColumn.append(a)
        secondColumn.append(b)

    res = 0
    for x in range(0, len(firstColumn)):
        maxIndexFirst = firstColumn.index(min(firstColumn))
        maxIndexSecond = secondColumn.index(min(secondColumn))

        maxFirst = firstColumn.pop(maxIndexFirst)
        maxSecond = secondColumn.pop(maxIndexSecond)

        diff = abs(int(maxFirst) - int(maxSecond))
        print(diff)
        res += diff

    print(res)

if __name__ == "__main__":
    main()
