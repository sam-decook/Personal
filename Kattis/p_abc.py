idx= {"A": 0, "B": 1, "C": 2}

nums = input().split()
letters = input()

nums.sort()

ans = [nums[idx[letters[i]]] for i in range(3)]
print(f"{ans[0]} {ans[1]} {ans[2]}")