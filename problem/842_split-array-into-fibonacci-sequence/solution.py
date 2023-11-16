# https://leetcode-cn.com/problems/split-array-into-fibonacci-sequence/

class Solution:
#    def __init__(self, confirm=[]):
#       self.confirm = confirm
    
    def splitIntoFibonacci(self, S: str) -> list:
        n = len(S)
        if S == '0'*n:
            return [0]*n

        # confirm 确认是否有符合要求的斐波那契数列
        confirm = []

        # 排除全0数列后，Fibonacci数至少三段
        for i in range(1,int(n/2)+1):
            if int(S[n-i:n]) > 2**31-1:
                return []
            fib = []
            # 日常丢失 range函数
            for j in range(i+1,2*i+1):
                def recursion(i,j,S):

                    # 查看 i,j,[n+2] = [n+1] +[n] 的关系

                    diff_val = str(int(S[n-i:n]) - int(S[n-j:n-i]))
                    #print('i:%d [n+2]:%s  j:%d [n+1]:%s diff:%s' %(n-i,S[n-i:n],n-j,S[n-j:n-i],diff_val))
                    x = len(diff_val)
                    # 可能会出现 S[n-j-x:n-j] 不存在
                    # 这段只能找出 [n+2] = [n+1] +[n]
                    if diff_val == (S[n-j-x:n-j]):
                        temp = (S[n-j:n-i]) + (S[n-i:n])
                        temp_n = int(S[n-i:n])
                        temp_m = int(S[n-j:n-i])
                        temp_x = 0
                        #print('n+2:%d,n+1:%d,temp:%s,S[n]:%s' %(temp_n,temp_m,temp,S[n-len(temp):n]))
                        while temp == S[n-len(temp):n]:
                            fib.append(temp_n)
                            if n - len(temp) == 0:
                                fib.append(temp_x)
                                confirm.append(1)
                                #print('a=b, temp=%s,m=%d,x=%d,fib=%s' %(temp,temp_m,temp_x,fib))
                                return fib
                            else:
                                temp_x = temp_n - temp_m
                                temp = str(temp_x) + temp
                                temp_n = temp_m
                                temp_m = temp_x
                                #print('else:',temp_n,temp_m,temp_x,fib)
                        '''
                        fib.append(int(S[n-i:n]))
                        #print(diff_val,S)
                        #print('n:%d,j:%d,x:%d,fib[]:%s' %(n,j,x,fib))
                        
                        if (n-j-x) == 0:
                            fib.append(int(S[n-j:n-i]))
                            fib.append(int(S[n-j-x:n-j]))
                            confirm.append(1)
                            #print(fib[::-1])
                            #return(fib[::-1])

                        else:
                            recursion(j,j+x,S[0:n-i])
                        '''
                    #print(n-j-x,diff_val,i,int(n/2))
                    if (n-j-x <= 0):
                        confirm.append(0)
                recursion(i,j,S)
                #print(fib[::-1],confirm)

                if 1 in confirm:
                    return(fib[::-1])

        if 1 not in confirm:
            #print('Not Fibonacci')
            return []

if __name__ == "__main__":
    a = '02246101626'
    b = '123456579'
    test = Solution()
    print(test.splitIntoFibonacci(a), '\n'+'*'*10)
    print(test.splitIntoFibonacci(b), '\n'+'*'*10)