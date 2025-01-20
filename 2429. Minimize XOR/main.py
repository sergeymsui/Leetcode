class Solution:

    def bits_num(self, num):
        b = [v for v in bin(num)][2:]
        zeros = 0
        for v in b:
            if v == "1":
                zeros += 1
        return zeros

    def minimizeXor(self, num1: int, num2: int) -> int:
        b1 = [v for v in bin(num1)][2:]
        b2 = [v for v in bin(num2)][2:]
        nonzero_counts = self.bits_num(num2)
        results_counts = self.bits_num(num1)

        b1 = ["0"] * (max(len(b1), len(b2)) - len(b1)) + b1

        s = 1
        while nonzero_counts > results_counts:
            if b1[-s] == "0":
                results_counts += 1
            b1[-s] = "1"
            s += 1

        s = 1
        while nonzero_counts < results_counts:
            if b1[-s] == "1":
                results_counts -= 1
            b1[-s] = "0"
            s += 1

        return int("".join(b1), 2)


def main():
    num1 = 3
    num2 = 5

    s = Solution()
    print(s.minimizeXor(num1, num2))


if __name__ == "__main__":
    main()
