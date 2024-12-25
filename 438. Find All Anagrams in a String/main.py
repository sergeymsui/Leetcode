from collections import defaultdict

s = "abab"
p = "ab"


class Solution:

    def compare(self, s_dict, p_dict):

        for k in p_dict:
            if s_dict[k] != p_dict[k]:
                return False

        return True

    def findAnagrams(self, s: str, p: str):

        res = []

        if len(s) < len(p):
            return res

        p_len = len(p)
        p_dict = defaultdict(int)
        s_dict = defaultdict(int)

        for i, e in enumerate(p):
            p_dict[e] += 1
            s_dict[s[i]] += 1

        i = 0
        j = i + p_len - 1

        while j < len(s):

            if self.compare(s_dict, p_dict):
                res.append(i)

            i += 1
            j += 1

            if j < len(s):
                s_dict[s[j]] += 1
                s_dict[s[i - 1]] -= 1

        return res


r = Solution()
print(r.findAnagrams(s, p))
