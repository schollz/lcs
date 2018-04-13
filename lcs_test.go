package lcs

import (
	"bytes"
	"errors"
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

type Replacement struct {
	From [2]int
	To   [2]int
}

func findReplacements(a, b string) []Replacement {
	l := lcs(a, b)
	a = "<" + a + ">"
	b = "<" + b + ">"
	l = "<" + l + ">"
	startI := 0
	endI := 1
	startJ := 0
	endJ := 1
	curLcs := 0
	replacements := []Replacement{}
	for {
		// end if they are both at the end
		if startI == len(a)-1 && startJ == len(b)-1 {
			break
		}
		if a[endI] == l[curLcs+1] && b[endJ] == l[curLcs+1] {
			// fmt.Printf("[%d,%d] - [%d,%d]\n", startI, endI, startJ, endJ)
			// fmt.Printf("'%s' - '%s'\n", string(a[startI:endI]), string(b[startJ:endJ]))
			if a[startI:endI] == b[startJ:endJ] {
				// no need to do anything
			} else {
				replacement := Replacement{
					From: [2]int{startI, endI - 1},
					To:   [2]int{startJ, endJ - 1},
				}
				replacements = append(replacements, replacement)
			}
			startI = endI
			startJ = endJ
			curLcs++
			continue
		}
		if a[endI] != l[curLcs+1] {
			endI++
		}
		if b[endJ] != l[curLcs+1] {
			endJ++
		}
	}
	return replacements
}

type Diff struct {
	Left    string
	Right   string
	Between string
}

func (d Diff) String() string {
	return fmt.Sprintf("'%s' '%s' '%s'", d.Left, d.Between, d.Right)
}

func getDiffs(a, b string) []Diff {
	replacements := findReplacements(a, b)
	diffs := make([]Diff, len(replacements))
	for diffI, replacement := range replacements {
		fmt.Println(replacement)
		d := Diff{
			Between: b[replacement.To[0]:replacement.To[1]],
		}
		if replacement.From[1] == 0 {
			d.Left = "" // start
		} else {
			for i := replacement.To[0] - 1; i >= 0; i-- {
				d.Left = b[i:replacement.To[0]]
				if strings.Index(b, d.Left)+len(d.Left) == replacement.To[0] {
					break
				}
			}
		}
		if replacement.To[0] == len(a) {
			d.Right = "" // end
		} else {
			// only search on the right side of the original
			for i := replacement.From[1] + 1; i < len(a); i++ {
				d.Right = a[replacement.From[1]:i]
				if strings.Index(a[replacement.From[0]:], d.Right)+len(a[:replacement.From[0]]) == replacement.From[1] {
					break
				}
			}
		}
		fmt.Println(d)
		diffs[diffI] = d
	}
	return diffs
}

func Patch(a string, diffs []Diff) (string, error) {
	for _, d := range diffs {
		fmt.Println(d)
		start := 0
		end := len(a)
		if d.Left != "" {
			start = strings.Index(a, d.Left) + len(d.Left)
		}
		if d.Right != "" {
			end = strings.Index(a[start:], d.Right) + len(a[:start])
		}
		if start < 0 || end < 0 {
			return "", errors.New("bad patch")
		}
		if end == 0 {
			// prepend beginning
			a = d.Between + a
		} else if start == len(a) {
			// append to end
			a = a + d.Between
		} else {
			a = a[:start] + d.Between + a[end:]
		}
		fmt.Println(a)
	}
	return a, nil
}

func TestPatch1(t *testing.T) {
	a := "1234557890550"
	b := "a55b1255555cd89"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(lcs(a, b))
	diffs := getDiffs(a, b)
	a2, err := Patch(a, diffs)
	assert.Nil(t, err)
	assert.Equal(t, b, a2)
}

func TestPatch2(t *testing.T) {
	a := "a the cat and the cow jumped over the moon"
	b := "the blue cow leaped over the moon into space"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(lcs(a, b))
	diffs := getDiffs(a, b)
	fmt.Println(diffs)
	a2, err := Patch(a, diffs)
	assert.Nil(t, err)
	assert.Equal(t, b, a2)
}

func TestPatch3(t *testing.T) {
	a := "a the cat and the cow jumped over the moon"
	b := "the blue cow leaped over the moon into space"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(lcs(a, b))
	diffs := getDiffs(a, b)
	fmt.Println(diffs)
	a2, err := Patch("aslkdfjalskdjf", diffs)
	assert.NotNil(t, err)
	assert.Equal(t, "", a2)
}

func TestLcs(t *testing.T) {
	a := "the cat sat on a lap"
	b := "the dog sat and flap"
	assert.Equal(t, "the  sat n lap", lcs(a, b))
}

func TestLcsByte(t *testing.T) {
	a := []byte("the cat sat on a lap")
	b := []byte("the dog sat and flap")
	assert.Equal(t, []byte("the  sat n lap"), lcsByte(a, b))
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
		}
		return ""
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

func lcsLengthByte(xs, ys []byte) []int {
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

func lcsByte(xs, ys []byte) []byte {
	nx := len(xs)
	ny := len(ys)

	if nx == 0 {
		return []byte(nil)
	} else if nx == 1 {
		if bytes.Contains(ys, xs) {
			return xs
		} else {
			return []byte(nil)
		}
	}

	i := nx / 2
	xb := xs[:i]
	xe := xs[i:]
	ll_b := lcsLengthByte(xb, ys)
	ll_e := lcsLengthByte(reverseByte(xe), reverseByte(ys))
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
	return append(lcsByte(xb, yb), lcsByte(xe, ye)...)
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
