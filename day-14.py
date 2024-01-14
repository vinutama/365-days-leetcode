"""
1657. Determine if Two Strings are Close

Two strings are considered close if you can attain one from the other using the following operations:

Operation 1: Swap any two existing characters.
For example, abcde -> aecdb
Operation 2: Transform every occurrence of one existing character into another existing character, and do the same with the other character.
For example, aacabb -> bbcbaa (all a's turn into b's, and all b's turn into a's)
You can use the operations on either string as many times as necessary.

Given two strings, word1 and word2, return true if word1 and word2 are close, and false otherwise.

 

Example 1:

Input: word1 = "abc", word2 = "bca"
Output: true
Explanation: You can attain word2 from word1 in 2 operations.
Apply Operation 1: "abc" -> "acb"
Apply Operation 1: "acb" -> "bca"
Example 2:

Input: word1 = "a", word2 = "aa"
Output: false
Explanation: It is impossible to attain word2 from word1, or vice versa, in any number of operations.
Example 3:

Input: word1 = "cabbba", word2 = "abbccc"
Output: true
Explanation: You can attain word2 from word1 in 3 operations.
Apply Operation 1: "cabbba" -> "caabbb"
Apply Operation 2: "caabbb" -> "baaccc"
Apply Operation 2: "baaccc" -> "abbccc"
 

Constraints:

1 <= word1.length, word2.length <= 105
word1 and word2 contain only lowercase English letters.
"""


class Solution:
    def closeStrings(self, word1: str, word2: str) -> bool:
        from collections import Counter

        if len(word1) != len(word2):
            return False

        word1map, word2map = Counter(word1), Counter(word2)
        cword1, cword2 = [], []

        for (key1, val1), (_, val2) in zip(word1map.items(), word2map.items()):
            if key1 not in word2map:
                return False

            cword1.append(val1)
            cword2.append(val2)

        cword1.sort()
        cword2.sort()

        return cword1 == cword2


sol = Solution()

print(sol.closeStrings("abc", "bca"))
print(sol.closeStrings("a", "aa"))
print(sol.closeStrings("cabbba", "abbccc"))
print(sol.closeStrings("cabbba", "aabbss"))
print(sol.closeStrings("uau", "ssx"))
print(sol.closeStrings("aaabbbbccddeeeeefffff", "aaaaabbcccdddeeeeffff"))
