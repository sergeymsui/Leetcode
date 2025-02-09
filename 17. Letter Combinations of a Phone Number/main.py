class Solution:
    def letterCombinations(self, digits: str):

        if len(digits) == 0:
            return []

        m_dict = {}

        m_dict[2] = "abc"
        m_dict[3] = "def"
        m_dict[4] = "ghi"
        m_dict[5] = "jkl"
        m_dict[6] = "mno"
        m_dict[7] = "pqrs"
        m_dict[8] = "tuv"
        m_dict[9] = "wxyz"

        res = []

        for n in digits:
            d = int(n)

            if len(res) == 0:
                res = [s for s in m_dict[d]]
            else:
                proto = []

                for r in res:
                    for s in m_dict[d]:
                        proto.append(r + s)
                res = proto

        return res


digits = "23"

s = Solution()

res = s.letterCombinations(digits)
print(res)
