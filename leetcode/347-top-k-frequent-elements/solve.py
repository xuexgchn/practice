# https://leetcode-cn.com/problems/top-k-frequent-elements/

class Solution:
    def topKFrequent(self, nums: list, k: int) -> list:
        import queue
        
        dic = {}
        for i in nums:
            if i not in dic:
                dic[i] = 1
            else:
                dic[i] = dic[i] + 1
        
        que = queue.PriorityQueue()
        for i in dic:
            que.put([dic[i], i])
        ans = []
        while que.qsize() > 0:
            s = que.get()
            ans.append(s[1])
        return ans[::-1][:k]

assert Solution().topKFrequent([1,1,1,2,2,3], 2) == [1,2]
assert Solution().topKFrequent([1], 1) == [1]
assert Solution().topKFrequent([-1, -1], 1) == [-1]