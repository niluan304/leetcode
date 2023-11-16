# -*- coding: utf-8 -*-
"""
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
"""
class Solution:
    def lengthOfLongestSubstring(self, s: str):
        str_1 = ''
        mid = ''
        if len(s) == 1:
            return 1
        for k, v in enumerate(s):
            # if str(v) not in  and k+1 == len(s):                
            if str(v) not in mid:
                mid += str(v)
                if k+1 == len(s) and len(mid) > len(str_1):
                    str_1 = mid
            
            else:
                v1 = mid.find(v) + 1
                # print('k:',k,'\tv1:',v1,'\tlen(mid):',len(mid))
                # print('str_1:', str_1, 'mid:', mid)
                if len(mid) > len(str_1):
                    str_1 = mid
                mid = s[k-len(mid)+v1:k+1]
                # print('str_1:', str_1, 'mid:', mid)
            
        return len(str_1)

if __name__ == '__main__':
    list_s=[
        "abcabcbb",
        "bbbbb",
        "pwwkew",
        " ",
        "a",
        "",
        "au",
        "asdf",
        "cdd",
        "abcdd"
    ]
    test = Solution()
    
    for i in list_s:
        print(test.lengthOfLongestSubstring(i),'←',i)