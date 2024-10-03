import re


def isPalindrome(s: str) -> bool:
    s = re.sub(r"[^a-zA-Z0-9]", "", s).lower()
    if s == "":
        return True

    i = 0
    x = len(s) - 1


    while i < x:
        if s[i] != s[x]:
            return False
        i += 1
        x -= 1
    return True


print(isPalindrome("race a car"))
