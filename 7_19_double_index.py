# 快慢指针
# 不需要创建新数组
# 用快指针进行迭代，对每个元素进行条件验证
# 验证通过，重新赋值给slow指针内元素，相当于重排


class test():
    pass

def removeElement(self, nums, value):
    slow = 0
    n=len(nums)
    for fast in range(n):
        if nums[fast] != val:
            nums[slow]=nums[fast]
            slow +=1
    return slow


if  __name__ == "__main__":
    m=test()
    m.removeElement([1,2,3,3],3)