while True:
    w, d = [int(i) for i in input().split()]

    if w == 0 and d == 0:
        break

    w, n = divmod(w, d)

    print(f"{w} {n} / {d}")