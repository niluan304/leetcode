# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next

# 1.转化为数字相加，再list(str)转化为链表，然后还需要反序list[::-1]
# 2.l1[i]+l2[i]，但这需要考虑溢出的问题
# 3.重构加法，反序相加__add__


class Solution:
    def addTwoNumbers(self, l1, l2):
        len1 = len(l1)
        len2 = len(l2)
        len0 = abs(len1 - len2)
        if len1 > len2:
            l2 = l2 + len0 * [0]
            print(l1)            
            print(l2)
        else:
            l1 = l1 + len0 * [0]
            print(l1)            
            print(l2)
        flag = 0
        l_target = []
        # 如何考虑溢出的问题？
        for i in range(0,max(len1,len2)):
            x = l1[i] + l2[i]
            if x > 9:
                flag = 1
                x = x -10
            l_target.append(x)
        print(l_target)
                       
            
            
if __name__ == "__main__":
    l2 = [9,9,9,9,9,9,9]
    l1 = [9,9,9,9]
    solution = Solution()
    solution.addTwoNumbers(l1,l2)
                