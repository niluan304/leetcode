class Solution:
    def shipWithinDays(self, weights: List[int], D: int) -> int:
        # 确定二分查找左右边界
        left, right = max(weights), sum(weights)
        while left < right:
            mid = (left + right) // 2
            # need 为需要运送的天数
            # cur 为当前这一天已经运送的包裹重量之和
            need, cur = 1, 0
            for weight in weights:
                if cur + weight > mid:
                    need += 1
                    cur = 0
                cur += weight

            if need <= D:
                right = mid
            else:
                left = mid + 1

        return left