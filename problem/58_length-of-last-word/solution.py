class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        try:
            # strip 砍掉首末的特定字符
            # lstrip left / right rstrip
            return(len(s.strip(' ').split(' ')[-1]))
        except:
            return 0
    

if __name__ == "__main__":
    test = Solution()
    test.lengthOfLastWord('   Hello    World   ')