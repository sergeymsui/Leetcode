class Solution:
    def multiply(self, num1: str, num2: str) -> str:

        num1 = [e for e in num1]
        num2 = [e for e in num2]

        len1 = len(num1)
        len2 = len(num2)

        max_len = max(len1, len2)
        num1 = ["0"] * (max_len - len1) + num1
        num2 = ["0"] * (max_len - len2) + num2

        res = [0] * (max_len**2)

        num1.reverse()
        num2.reverse()

        for i, v in enumerate(num2):
            for j, b in enumerate(num1):
                m = int(v) * int(b)
                res[i + j] += m

        i = 0
        while i < len(res) - 1:
            v = res[i]
            res[i] = v % 10
            res[i + 1] += int(v / 10)
            i += 1

        res.reverse()

        i = 0
        while i < len(res) - 1 and res[i] == 0:
            i += 1

        return "".join([str(e) for e in res[i:]])


if __name__ == "__main__":
    num1 = "1999"
    num2 = "0"

    s = Solution()
    print(s.multiply(num1, num2))
