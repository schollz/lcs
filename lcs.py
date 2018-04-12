import itertools

def lcs_lens(xs, ys):
    curr = list(itertools.repeat(0, 1 + len(ys)))
    for x in xs:
        prev = list(curr)
        for i, y in enumerate(ys):
            if x == y:
                curr[i + 1] = prev[i] + 1
            else:
                curr[i + 1] = max(curr[i], prev[i + 1])
    return curr

def lcs(xs, ys):
    nx, ny = len(xs), len(ys)
    if nx == 0:
        return []
    elif nx == 1:
        return [xs[0]] if xs[0] in ys else []
    else:
        i = nx // 2
        print("i",i)
        xb, xe = xs[:i], xs[i:]
        print("xb",xb)
        print("xe",xe)
        ll_b = lcs_lens(xb, ys)
        print(xb,ys,ll_b)
        print("ll_b", ll_b)
        ll_e = lcs_lens(xe[::-1], ys[::-1])
        print("ll_e", ll_e)
        for j in range(ny+1):
            print(ll_b[j] + ll_e[ny - j])
        val = -1
        k = 0
        for j in range(ny+1):
            print(j,ll_b[j] + ll_e[ny-j])
            if ll_b[j] + ll_e[ny-j] >= val:
                val = ll_b[j] + ll_e[ny-j]
                k = j
        print('val,k')
        print(val,k)
        val, k = max((ll_b[j] + ll_e[ny - j], j)
                    for j in range(ny + 1))
        print('actual')
        print(val,k)
        yb, ye = ys[:k], ys[k:]
        return lcs(xb, yb) + lcs(xe, ye)


print(lcs_lens("CAT","SPLAT"))
print(lcs("CAT","SPLAT"))