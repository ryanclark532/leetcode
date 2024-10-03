class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


def invertTree(root: Optional[TreeNode]) -> Optional[TreeNode]:
    def invert(node: TreeNode):
        node.left, node.right = node.right, node.left
        invert(node.left)



    return root
