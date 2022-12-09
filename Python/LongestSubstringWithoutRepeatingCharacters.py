class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        left = right = result = 0
        while right < len(s): 

            # check if window contains the next char
            while s[right] in s[left:right]:
                left +=1

            # update the result, if longest
            if right - left + 1 > result:
                result = right- left +1
                
            right +=1
        return result