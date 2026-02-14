class Solution(object):
    def lengthOfLongestSubstring(self, s):
        """
        :type s: str
        :rtype: int
        """
        last = {}
        left = 0
        best = 0

        for right , ch in enumerate(s):
            if ch in last and last[ch] >= left:
                left = last[ch] + 1

            last[ch] = right
            curr_len = right - left + 1
            if curr_len > best:
                best = curr_len
        return best