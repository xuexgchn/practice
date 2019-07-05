class Solution:
    def isValid(self, s: str) -> bool:
        top = -1
        stack = [' ' for i in range(len(s))]
        index = 0
        while index < len(s):
            if s[index] == ')':
                if top >= 0 and stack[top] == '(':
                    top -= 1
                else:
                    return False
            elif s[index] == '}':
                if top >= 0 and stack[top] == '{':
                    top -= 1
                else:
                    return False
            elif s[index] == ']':
                if top >= 0 and stack[top] == '[':
                    top -= 1
                else:
                    return False
            else:
                top += 1
                stack[top] = s[index]
            index  += 1
        return top == -1

assert Solution().isValid("()") == True
assert Solution().isValid("()[]{}") == True
assert Solution().isValid("(]") == False
assert Solution().isValid("([)]") == False
assert Solution().isValid("{[]}") == True