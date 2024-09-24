def twoSum(nums, target):
    a = dict()
    for i, x in enumerate(nums):
       a[x] = i

    for i, x in enumerate(nums):
        complement = target - x
        if complement in a:
            val = a[complement]
            if val and val !=i:
                return [a[complement],i]


print(twoSum([2,5,5,11], 10))
