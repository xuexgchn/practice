class Solution:
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False
        y = x
        ans = 0
        while x > 0:
            ans = ans * 10 + x % 10
            x //= 10
        if ans == y:
            return True
        return False

assert Solution().isPalindrome(121) == True
assert Solution().isPalindrome(-121) == False
assert Solution().isPalindrome(10) == False
