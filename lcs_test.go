package lcs

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkLcs(b *testing.B) {
	a := "the cat sat on a lap"
	b2 := "the dog sat and flap"
	for i := 0; i < b.N; i++ {
		lcs(a, b2)
	}
}

func BenchmarkLcs2(b *testing.B) {
	a := "the cat sat on a lap"
	b2 := "the dog sat and flap"
	for i := 0; i < b.N; i++ {
		lcs2(a, b2)
	}
}

func TestLcs2(t *testing.T) {
	a := "the cat sat on a lap"
	b := "the dog sat and flap"
	assert.Equal(t, "the  sat n lap", lcs2(a, b))
}

func TestAlign(t *testing.T) {
	a := "<the cat jumped>"
	b := "<and the dog leaped>"
	l := lcs(a, b)
	fmt.Println(l)
	aPos := make([]int, len(l))
	bPos := make([]int, len(l))
	j := 0
	for i := range a {
		if a[i] == l[j] {
			aPos[j] = i
			j++
		}
	}
	j = 0
	for i := range b {
		if b[i] == l[j] {
			bPos[j] = i
			j++
		}
	}
	fmt.Println(aPos)
	fmt.Println(bPos)
}
func TestLcs(t *testing.T) {
	a := "the cat sat on a lap"
	b := "the dog sat and flap"
	assert.Equal(t, "the  sat n lap", lcs(a, b))
}

func lcs2(a, b string) string {
	arunes := []rune(a)
	brunes := []rune(b)
	aLen := len(arunes)
	bLen := len(brunes)
	lengths := make([][]int, aLen+1)
	for i := 0; i <= aLen; i++ {
		lengths[i] = make([]int, bLen+1)
	}
	// row 0 and column 0 are initialized to 0 already

	for i := 0; i < aLen; i++ {
		for j := 0; j < bLen; j++ {
			if arunes[i] == brunes[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}

	// read the substring out from the matrix
	s := make([]rune, 0, lengths[aLen][bLen])
	for x, y := aLen, bLen; x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, arunes[x-1])
			x--
			y--
		}
	}
	// reverse string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func lcsLength(xs, ys string) []int {
	curr := make([]int, 1+len(ys))
	prev := make([]int, 1+len(ys))
	for _, x := range xs {
		for j, y := range ys {
			prev[j+1] = curr[j+1]
			if x == y {
				curr[j+1] = prev[j] + 1
			} else {
				a := curr[j]
				if prev[j+1] > a {
					a = prev[j+1]
				}
				curr[j+1] = a
			}
		}
	}
	return curr
}

func lcs(xs, ys string) string {
	nx := len(xs)
	ny := len(ys)

	if nx == 0 {
		return ""
	} else if nx == 1 {
		if strings.Contains(ys, xs) {
			return xs
		} else {
			return ""
		}
	}

	i := nx / 2
	xb := xs[:i]
	xe := xs[i:]
	ll_b := lcsLength(xb, ys)
	ll_e := lcsLength(reverse(xe), reverse(ys))
	val := -1
	k := 0
	for j := 0; j < ny+1; j++ {
		if ll_b[j]+ll_e[ny-j] >= val {
			val = ll_b[j] + ll_e[ny-j]
			k = j
		}
	}
	yb := ys[:k]
	ye := ys[k:]
	return lcs(xb, yb) + lcs(xe, ye)
}

func reverse(s string) string {
	b := make([]byte, len(s))
	var j int = len(s) - 1
	for i := 0; i <= j; i++ {
		b[j-i] = s[i]
	}

	return string(b)
}

func reverseByte(s []byte) []byte {
	b := make([]byte, len(s))
	var j int = len(s) - 1
	for i := 0; i <= j; i++ {
		b[j-i] = s[i]
	}

	return b
}

func TestReverseByte(t *testing.T) {
	a := []byte{11, 22, 33}
	assert.Equal(t, []byte{33, 22, 11}, reverseByte(a))
}

// def lcs(xs, ys):
//     nx, ny = len(xs), len(ys)
//     if nx == 0:
//         return []
//     elif nx == 1:
//         return [xs[0]] if xs[0] in ys else []
//     else:
//         i = nx // 2
//         xb, xe = xs[:i], xs[i:]
//         ll_b = lcs_lens(xb, ys)
//         ll_e = lcs_lens(xe[::-1], ys[::-1])
//         _, k = max((ll_b[j] + ll_e[ny - j], j)
//                     for j in range(ny + 1))
//         yb, ye = ys[:k], ys[k:]
//         return lcs(xb, yb) + lcs(xe, ye)

func TestLcsLength(t *testing.T) {
	fmt.Println(lcsLength("CAT", "SPLAT"))
}
