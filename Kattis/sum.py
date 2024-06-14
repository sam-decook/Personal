from sys import stdin

for line in stdin:
    nums = [int(i) for i in line.split()]
    total = sum(nums)
    for n in nums:
        if total - n == n:
            print(n)
            break