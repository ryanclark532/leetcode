



def isAnagram(s: str, t: str) -> bool:
    def getStringCount(s: str):
        res = {}
        for n in s:
            if res.get(n) is None:
                res[n] = 1
            else: 
                res[n]+=1 
        return res

    sc = getStringCount(s)
    st = getStringCount(t)
    print(sc)
    print(st)

    for x in sc:
        if st.get(x) is None or st[x] != sc[x]:
            return False

    return True



print(isAnagram("anagram", "nagaram"))


v = {
    "t": 1
}

print(v.get("x"))
