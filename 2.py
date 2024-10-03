from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


def addTwoNumbers(l1: Optional[ListNode], l2: Optional[ListNode]):
    l1q = []
    while l1:
        l1q.append(l1.val)
        l1 = l1.next

    val1 = int(''.join(map(str, l1q)))
    

    l2q = []
    while l2:
        l2q.append(l2.val)
        l2 = l2.next


    val2 = int(''.join(map(str, l2q)))

    total = val1 + val2

    return [int(digit) for digit in str(number)[::-1]]




