m_string = "pwwkew"


def longest_substring(m_string: str) -> str:

    s = ""

    i = 0
    j = 1

    while i < len(m_string) and j < len(m_string):

        if len(m_string[i:j]) > len(s):
            s = m_string[i:j]

        if m_string[j] not in m_string[i:j]:
            j += 1
        else:
            i += 1

    return s


print(longest_substring(m_string))
