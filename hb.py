#!/usr/bin/python3
# Implementation of Hirschberg's algorithm
# follows https://en.wikipedia.org/wiki/Hirschberg%27s_algorithm
import numpy
import copy

def Ins(x):
    return -1


def Del(x):
    return -2


def Sub(x, y):
    if x == y:
        return 3
    return -1


def NWScore(x, y):
    rowPrev = numpy.zeros(len(y)+1)
    for j, yj in enumerate(' ' + y):
        if j == 0:
            continue
        rowPrev[j] = rowPrev[j-1] + Ins(yj)

    rowCur = numpy.zeros(len(y)+1)
    lastLine = numpy.zeros(len(y)+1)
    for i, xi in enumerate(' ' + x):
        if i == 0:
            continue
        rowCur[0] = rowPrev[0] + Del(xi)
        for j, yj in enumerate(' ' + y):
            if j == 0:
                continue
            scoreSub = rowPrev[j - 1] + Sub(xi, yj)
            scoreDel = rowPrev[j] + Del(xi)
            scoreIns = rowCur[j - 1] + Ins(yj)
            rowCur[j] = max([scoreIns, scoreDel, scoreSub])
        rowPrev = copy.deepcopy(rowCur)
    for j, _ in enumerate(' ' + y):
        lastLine[j] = rowCur[j]

    return lastLine


def Hirschberg(x, y):
    z = ""
    w = ""
    if len(x) == 0:
        for _, yi in enumerate(y):
            z = z + '-'
            w = w + yi
    elif len(y) == 0:
        for _, xi in enumerate(x):
            z = z + xi
            w = w + '-'
    elif len(x) == 1:
        found = False
        for _, yi in enumerate(y):
            w = w + yi
            if yi == x and not found:
                z = z + x
                found = True
            else:
                z = z + '-'
        if not found:
            z = x + z[1:]
    elif len(y) == 1:
        found = False
        for _, xi in enumerate(x):
            z = z + xi
            if xi == y and not found:
                w = w + y
                found = True
            else:
                w = w + '-'
        if not found:
            w = y + w[1:]
    else:
        xlen = len(x)
        xmid = len(x) // 2
        ylen = len(y)
        scoreL = NWScore(x[:xmid], y)
        scoreR = NWScore(x[xmid:][::-1], y[::-1])
        revScoreR = scoreR[::-1]

        ymid = 0
        ymax = 0
        for i, _ in enumerate(scoreL):
            if scoreL[i] + revScoreR[i] > ymax:
                ymax = int(scoreL[i] + revScoreR[i])
                ymid = i
        z1, w1 = Hirschberg(x[:xmid], y[:ymid])
        z2, w2 = Hirschberg(x[xmid:], y[ymid:])
        z = z1 + z2
        w = w1 + w2
    return z, w

"""
        T   A   T   G   C  <-Y
    0  -2  -4  -6  -8 -10
 A -2  -1   0  -2  -4  -6
 G -4  -3  -2 *-1*  0  -2
 T -6  -2  -4   0  -2  -1
 A -8  -4   0  -2  -1  -3

 ^
 X
"""

scoreSub = 0 + -1
scoreDel = -2 + -2
scoreIns = -2 + -2

"""
[-2. -1.  0. -2. -4. -6.]
[-4. -3. -2.*-3.*-1. -2.]
[-6. -4. -4. -2. -3. -4.]
[-8. -6. -4. -4. -5. -6.]
[-8. -6. -4. -4. -5.  0.]
"""

print(NWScore("AGTA", "TATGC"))
print(NWScore("ACGC", "CGTAT"))

z,w = Hirschberg("AGTACGCA", "TATGC")
print(z)
print(w)


z,w = Hirschberg("The cat jumped over the moon", "The blue cat jumped over the fence")
print(z)
print(w)


