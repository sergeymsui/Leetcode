class Solution:

    def shiftingLetters(self, s: str, shifts: list):

        cshift = [0 for _ in s]

        for [i, f, d] in shifts:
            j = i
            while j <= f:
                cshift[j] += -1 if d == 0 else 1
                j += 1

        ns = [""] * len(s)

        for i, v in enumerate(cshift):

            ns[i] = chr(ord(s[i]) + v)

            if ord(ns[i]) > 122:
                ns[i] = chr(ord(ns[i]) - 122 + ord("a"))
            if ord(ns[i]) < 97:
                ns[i] = chr(ord("z") - (96 - ord(ns[i])))

        return "".join(ns)


def main():
    s = "xuwdbdqik"
    shifts = [
        [4, 8, 0],
        [4, 4, 0],
        [2, 4, 0],
        [2, 4, 0],
        [6, 7, 1],
        [2, 2, 1],
        [0, 2, 1],
        [8, 8, 0],
        [1, 3, 1],
    ]

    r = Solution()
    print(r.shiftingLetters(s, shifts))


if __name__ == "__main__":
    main()
