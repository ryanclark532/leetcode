def search(nums, target) -> int:

    def s(arr, low, high, x):

        if high >= low:

            mid = (high + low) // 2

            if arr[mid] == x:
                return mid

            elif arr[mid] > x:
                return s(arr, low, mid-1, x)
            else:
                return s(arr, mid+1, high, x)
        else:
            return -1


    return s(nums, 0, len(nums)-1, target)



print(search([-1,0,3,5,9,12], 9))
