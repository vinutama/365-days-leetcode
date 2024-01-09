"""
# Definition for a binary tree node.
# class TreeNode(object):
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right

"""


class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right


class Solution(object):
    def createBST(self, input):
        if not input:
            return

        root = TreeNode(input[0])
        queue = [root]
        i = 1

        while queue:
            node = queue.pop(0)

            if i < len(input) and input[i] is not None:
                node.left = TreeNode(input[i])
                queue.append(node.left)

            i += 1
            if i < len(input) and input[i] is not None:
                node.right = TreeNode(input[i])
                queue.append(node.right)

            i += 1

        return root

    def leafSimilar(self, root1, root2):
        """
        :type root1: TreeNode
        :type root2: TreeNode
        :rtype: bool
        """
        root1 = self.createBST(root1)
        root2 = self.createBST(root2)
        res1 = self.search_latest_leaf(root1, [])
        res2 = self.search_latest_leaf(root2, [])

        return res1 == res2

    def search_latest_leaf(self, root, res):
        if not root:
            return res

        if root.left or root.right:
            self.search_latest_leaf(root.left, res)
            self.search_latest_leaf(root.right, res)
        else:
            # found edge leaf
            res.append(root.val)

        return res


sol = Solution()
print(
    sol.leafSimilar(
        [3, 5, 1, 6, 2, 9, 8, None, None, 7, 4],
        [3, 5, 1, 6, 7, 4, 2, None, None, None, None, None, None, 9, 8],
    )
)

print(sol.leafSimilar([1, 2, 3], [1, 3, 2]))
print(sol.leafSimilar([1, 2], [2, 2]))
