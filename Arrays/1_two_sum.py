class Solution:
    def __init__(self, nums: list[int], target: int):
        self.nums = nums
        self.target = target

    def twoSum(nums, target) -> list[int]:
        hash_map = {}
        for i in range(len(nums)):
            if nums[i] in hash_map:
                return [i, hash_map[nums[i]]]
            else:
                hash_map[target - nums[i]] = i


# should return [3, 2]
print(Solution.twoSum(nums=[3, 5, 7, 8, 1, 2], target=15))
