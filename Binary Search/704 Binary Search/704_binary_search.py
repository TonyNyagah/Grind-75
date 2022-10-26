class Solution:
    def __init__(self, nums: list[int], target: int):
        self.nums = nums
        self.target = target

    def search(nums, target) -> int:
        low = 0
        high = len(nums) - 1

        while low <= high:
            mid = (low + high) // 2
            guess = nums[mid]
            if guess == target:
                return mid
            if guess > target:
                high = mid - 1
            else:
                low = mid + 1

        return -1


# should return 4
print(Solution.search(nums=[-1, 0, 3, 5, 9, 12], target=9))
