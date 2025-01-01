n = int(input())


def bracket_combination(br=0, left=0, right=0, prefix=""):

    if left == 0 and right == 0:
        print(prefix)

    if left > 0:
        bracket_combination(br + 1, left - 1, right, prefix + "(")

    if br > 0 and right > 0:
        bracket_combination(br - 1, left, right - 1, prefix + ")")

    pass


bracket_combination(0, n, n)
