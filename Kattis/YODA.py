arr1 = [int(i) for i in input()]
arr2 = [int(i) for i in input()]

while len(arr1) < len(arr2):
    arr1 = [0] + arr1
while len(arr1) > len(arr2):
    arr2 = [0] + arr2

colArr1 = []
colArr2 = []

for a1, a2 in zip(arr1, arr2):
    if a1 > a2:
        colArr1.append(a1)
    elif a1 < a2:
        colArr2.append(a2)
    else:
        colArr1.append(a1)
        colArr2.append(a2)

if len(colArr1) > 0:
    s = int(''.join([str(i) for i in colArr1]))
    print(f"{s}")
else:
    print("YODA")


if len(colArr2) > 0:
    s = int(''.join([str(i) for i in colArr2]))
    print(f"{s}")
else:
    print("YODA")
