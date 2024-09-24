




def isValid(s):
    if len(s) == 1:
        return False

    def isMatchingOpening(s, y):
        if y is None:
            return s == "[" or s == "{" or s == "("
        if s=="{":
            return y=="}"
        elif s == "[":
            return y== "]"
        else:
            return y== ")"

    queue = []
    for c in s:
        if isMatchingOpening(c, None):
            queue.append(c)
        else:
            l = len(queue) -1
            if l >= 0 and isMatchingOpening(queue[l], c):
                queue.pop()
            else:
                return False
    return len(queue) == 0



print(isValid("()"))