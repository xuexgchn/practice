# https://leetcode-cn.com/problems/first-unique-character-in-a-string/

class Solution:
    def firstUniqChar(self, s: str) -> int:
        dic = {}
        for i in s:
            if i not in dic:
                dic[i] = 1
            else:
                dic[i] = dic[i] + 1

        for i in range(len(s)):
            if dic[s[i]] == 1:
                return i        
        return -1

assert(Solution().firstUniqChar("leetcode")) == 0
assert(Solution().firstUniqChar("loveleetcode")) == 2