

class TreeNode:
     def __init__(self, x):
         self.val = x
         self.left = None
         self.right = None

def lowestCommonAncestor(root, p, q) -> TreeNode:

    while True:
        if root.val > p.val and root.val > q.val:
            root = root.left

        elif root.val < p.val and root.val < q.val:
            root = root.right

        else:
            return root
