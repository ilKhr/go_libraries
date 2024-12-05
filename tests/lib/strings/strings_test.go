package strings_test

// TODO: draw conclusions after all methods have been used
/*
any
func
неизменяемый размер reader
Seek и ReadAt
*/

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClone(t *testing.T) {
	a := ""
	s := "abc"

	assert.NotEqual(t, a, s)

	a = strings.Clone(s)

	assert.Equal(t, a, s)
}

func TestCompare(t *testing.T) {
	fi := "abc"
	se := "def"
	th := "abc"

	assert.Equal(t, strings.Compare(fi, th), 0)
	assert.Equal(t, strings.Compare(fi, se), -1)
	assert.Equal(t, strings.Compare(se, fi), 1)

}

func TestContains(t *testing.T) {
	s := "abcde"

	assert.Equal(t, strings.Contains(s, "bcd"), true)
	assert.Equal(t, strings.Contains(s, "ae"), false)
	assert.Equal(t, strings.Contains(s, "false"), false)

}

func TestContainsAny(t *testing.T) {
	s := "abcde"

	assert.Equal(t, strings.ContainsAny(s, "bcd"), true)
	assert.Equal(t, strings.ContainsAny(s, "ae"), true)
	assert.Equal(t, strings.ContainsAny(s, "dk"), true)
	assert.Equal(t, strings.ContainsAny(s, "BCD"), false)
	assert.Equal(t, strings.ContainsAny(s, "lj"), false)

}

func TestContainsFunc(t *testing.T) {
	fi := "abcde"
	se := "ijklm"

	checkRBelongInterval := func(r rune) bool {
		rs := string(r)

		left := strings.Compare(rs, "c")
		right := strings.Compare(rs, "h")

		if left >= 0 && right <= 0 {
			return true
		}

		return false
	}

	assert.Equal(t, strings.ContainsFunc(fi, checkRBelongInterval), true)
	assert.Equal(t, strings.ContainsFunc(se, checkRBelongInterval), false)
}

func TestContainsRune(t *testing.T) {
	s := "abcde"

	assert.Equal(t, strings.ContainsRune(s, 'c'), true)
	assert.Equal(t, strings.ContainsRune(s, 'j'), false)
}

func TestCount(t *testing.T) {
	s := "aabaacde"

	assert.Equal(t, strings.Count(s, "aa"), 2)
	assert.Equal(t, strings.Count(s, "a"), 4)
	assert.Equal(t, strings.Count(s, ""), 9)
}

func TestCut(t *testing.T) {
	s := "aabbccddeeff"

	bef1, af1, status1 := strings.Cut(s, "cc")

	bef2, af2, status2 := strings.Cut(s, "kk")

	require.True(t, status1)
	assert.Equal(t, bef1, "aabb")
	assert.Equal(t, af1, "ddeeff")

	require.False(t, status2)
	assert.Equal(t, bef2, s)
	assert.Equal(t, af2, "")
}

func TestCutPrefix(t *testing.T) {
	s := "aabbccddeeff"

	af1, status1 := strings.CutPrefix(s, "cc")

	af2, status2 := strings.CutPrefix(s, "kk")

	af3, status3 := strings.CutPrefix(s, "")

	af4, status4 := strings.CutPrefix(s, "aa")

	require.False(t, status1)
	assert.Equal(t, af1, s)

	require.False(t, status2)
	assert.Equal(t, af2, s)

	require.True(t, status3)
	assert.Equal(t, af3, s)

	require.True(t, status4)
	assert.Equal(t, af4, "bbccddeeff")
}

func TestCutSuffix(t *testing.T) {
	s := "aabbccddeeff"

	bf1, status1 := strings.CutSuffix(s, "cc")

	bf2, status2 := strings.CutSuffix(s, "kk")

	bf3, status3 := strings.CutSuffix(s, "")

	bf4, status4 := strings.CutSuffix(s, "ff")

	require.False(t, status1)
	assert.Equal(t, bf1, s)

	require.False(t, status2)
	assert.Equal(t, bf2, s)

	require.True(t, status3)
	assert.Equal(t, bf3, s)

	require.True(t, status4)
	assert.Equal(t, bf4, "aabbccddee")
}

func TestEqualFold(t *testing.T) {
	s := "aabbccddeeff"

	res1 := strings.EqualFold(s, "AABBCCDDEEFF")
	res2 := strings.EqualFold(s, "AAbbCCddEEff")
	res3 := strings.EqualFold(s, "aa")

	require.True(t, res1)
	require.True(t, res2)
	require.False(t, res3)
}
func TestFields(t *testing.T) {
	s := "aabbcc ddeeff"

	res1 := strings.Fields(s)
	res2 := strings.Fields("")

	require.Equal(t, res1, []string{"aabbcc", "ddeeff"})
	require.Equal(t, res2, []string{})
}

func TestFieldsFunc(t *testing.T) {
	s := "test1;test2;test3"

	// when true - make cut
	f := func(r rune) bool {
		return !unicode.IsLetter(r)
	}

	res1 := strings.FieldsFunc(s, f)

	require.Equal(t, res1, []string{"test", "test", "test"})
}

func TestHasPrefix(t *testing.T) {
	s := "Bearer 12345"

	res1 := strings.HasPrefix(s, "Bearer")
	res2 := strings.HasPrefix(s, "BEARER")

	require.Equal(t, res1, true)
	require.Equal(t, res2, false)
}

func TestHasSuffix(t *testing.T) {
	s := "Bearer 12345"

	res1 := strings.HasSuffix(s, "45")
	res2 := strings.HasSuffix(s, "12")

	require.Equal(t, res1, true)
	require.Equal(t, res2, false)
}

func TestIndex(t *testing.T) {
	s := "Bearer 12345"

	res1 := strings.Index(s, "45")
	res2 := strings.Index(s, "12")
	res3 := strings.Index(s, "9")

	require.Equal(t, res1, 10)
	require.Equal(t, res2, 7)
	require.Equal(t, res3, -1)

}

func TestIndexAny(t *testing.T) {
	s := "Bearer 12345"

	res1 := strings.IndexAny(s, "9876B")
	res2 := strings.IndexAny(s, "9876b")
	res3 := strings.IndexAny(s, "LKP")

	require.Equal(t, 0, res1)
	require.Equal(t, -1, res2)
	require.Equal(t, -1, res3)

}

func TestIndexByte(t *testing.T) {
	s := "Bearer 12345"

	res1 := strings.IndexByte(s, '4')
	res2 := strings.IndexByte(s, '9')

	require.Equal(t, 10, res1)
	require.Equal(t, -1, res2)
}

func TestIndexFunc(t *testing.T) {
	s := "Bearer 🎈 12345"
	s2 := "Bearer 12345"

	f := func(r rune) bool {
		return (r >= 0x1F300 && r <= 0x1F5FF) // Miscellaneous Symbols and Pictographs
	}

	res1 := strings.IndexFunc(s, f)
	res2 := strings.IndexFunc(s2, f)

	require.Equal(t, 7, res1)
	require.Equal(t, -1, res2)
}

func TestIndexRune(t *testing.T) {
	s := "Bearer 12345"

	res1 := strings.IndexRune(s, '4')
	res2 := strings.IndexRune(s, '9')

	require.Equal(t, 10, res1)
	require.Equal(t, -1, res2)
}

func TestJoin(t *testing.T) {
	s := []string{"1", "2", "3", "4"}

	res1 := strings.Join(s, ";")
	res2 := strings.Join(s, "")

	require.Equal(t, "1;2;3;4", res1)
	require.Equal(t, "1234", res2)
}

func TestLastIndex(t *testing.T) {
	s := "012345345123"

	res1 := strings.LastIndex(s, "345")
	res2 := strings.LastIndex(s, "123")
	res3 := strings.LastIndex(s, "9")

	require.Equal(t, 6, res1)
	require.Equal(t, 9, res2)
	require.Equal(t, -1, res3)
}

func TestLastIndexAny(t *testing.T) {
	s := "Bearer 12345B"

	res1 := strings.LastIndexAny(s, "9876B")
	res2 := strings.LastIndexAny(s, "9876b")
	res3 := strings.LastIndexAny(s, "LKP")

	require.Equal(t, 12, res1)
	require.Equal(t, -1, res2)
	require.Equal(t, -1, res3)
}

func TestLastIndexByte(t *testing.T) {
	s := "Bearer 12345"

	res1 := strings.LastIndexByte(s, '4')
	res2 := strings.LastIndexByte(s, '9')

	require.Equal(t, 10, res1)
	require.Equal(t, -1, res2)
}

func TestLastIndexFunc(t *testing.T) {
	s := "Bearer 🎈 12345"
	s2 := "Bearer 12345"

	f := func(r rune) bool {
		return (r >= 0x1F300 && r <= 0x1F5FF) // Miscellaneous Symbols and Pictographs
	}

	res1 := strings.LastIndexFunc(s, f)
	res2 := strings.LastIndexFunc(s2, f)

	require.Equal(t, 7, res1)
	require.Equal(t, -1, res2)
}

func TestMap(t *testing.T) {
	s := "expected string"

	f := func() func(rune) rune {

		isUpper := true

		return func(r rune) rune {
			var res rune

			if isUpper {
				res = unicode.ToUpper(r)
			} else {
				res = unicode.ToLower(r)
			}

			isUpper = !isUpper

			return res

		}
	}

	res := strings.Map(f(), s)

	require.Equal(t, "ExPeCtEd sTrInG", res)
}

func TestRepeat(t *testing.T) {

	res1 := strings.Repeat("k", 5)
	res2 := strings.Repeat("", 6)
	res3 := strings.Repeat("p", 0)

	require.Equal(t, "kkkkk", res1)
	require.Equal(t, "", res2)
	require.Equal(t, "", res3)
}

func TestReplace(t *testing.T) {

	s := "replaceTHISwordandTHISandTHIS"

	res1 := strings.Replace(s, "THIS", "this", -1)
	res2 := strings.Replace(s, "THIS", "this", 0)
	res3 := strings.Replace(s, "THIS", "this", 1)
	res4 := strings.Replace(s, "THIS", "this", 2)
	res5 := strings.Replace(s, "", ";", 4)

	require.Equal(t, "replacethiswordandthisandthis", res1)
	require.Equal(t, s, res2)
	require.Equal(t, "replacethiswordandTHISandTHIS", res3)
	require.Equal(t, "replacethiswordandthisandTHIS", res4)
	require.Equal(t, ";r;e;p;laceTHISwordandTHISandTHIS", res5)
}

func TestReplaceAll(t *testing.T) {

	s := "setsemilicon"

	res1 := strings.ReplaceAll(s, "", ";")
	res2 := strings.ReplaceAll(s, "semilicon", ";")

	require.Equal(t, ";s;e;t;s;e;m;i;l;i;c;o;n;", res1)
	require.Equal(t, "set;", res2)
}

func TestSplit(t *testing.T) {

	s := "set;semilicon;separator"
	s2 := "kakava"

	res1 := strings.Split(s, ";")
	res2 := strings.Split(s2, "")
	res3 := strings.Split(s2, "NOT_EXISTS")

	require.Equal(t, []string{"set", "semilicon", "separator"}, res1)
	require.Equal(t, []string{"k", "a", "k", "a", "v", "a"}, res2)
	require.Equal(t, []string{s2}, res3)
}

func TestSplitAfter(t *testing.T) {

	s := "set;semilicon;separator"
	s2 := "kakava"

	res1 := strings.SplitAfter(s, ";")
	res2 := strings.SplitAfter(s2, "")
	res3 := strings.SplitAfter(s2, "NOT_EXISTS")

	require.Equal(t, []string{"set;", "semilicon;", "separator"}, res1)
	require.Equal(t, []string{"k", "a", "k", "a", "v", "a"}, res2)
	require.Equal(t, []string{s2}, res3)
}

func TestSplitAfterN(t *testing.T) {

	s := "set;semilicon;separator;1;2;3;4;5"
	s2 := "kakava"

	res1 := strings.SplitAfterN(s, ";", 3)
	res2 := strings.SplitAfterN(s2, "", 2)
	res3 := strings.SplitAfterN(s2, "NOT_EXISTS", -1)
	res4 := strings.SplitAfterN(s2, "", -1)

	require.Equal(t, []string{"set;", "semilicon;", "separator;1;2;3;4;5"}, res1)
	require.Equal(t, []string{"k", "akava"}, res2)
	require.Equal(t, []string{s2}, res3)
	require.Equal(t, []string{"k", "a", "k", "a", "v", "a"}, res4)
}

func TestSplitN(t *testing.T) {

	s := "set;semilicon;separator"
	s2 := "kakava"

	res1 := strings.SplitN(s, ";", 2)
	res2 := strings.SplitN(s2, "", 3)
	res3 := strings.SplitN(s2, "NOT_EXISTS", -1)

	require.Equal(t, []string{"set", "semilicon;separator"}, res1)
	require.Equal(t, []string{"k", "a", "kava"}, res2)
	require.Equal(t, []string{s2}, res3)
}

func TestToLower(t *testing.T) {

	s := "TEST"
	s2 := ""

	res1 := strings.ToLower(s)
	res2 := strings.ToLower(s2)

	require.Equal(t, "test", res1)
	require.Equal(t, "", res2)
}

func TestToLowerSpecial(t *testing.T) {

	s := "TEST"
	s2 := ""

	res1 := strings.ToLowerSpecial(unicode.AzeriCase, s)
	res2 := strings.ToLowerSpecial(unicode.AzeriCase, s2)

	require.Equal(t, "test", res1)
	require.Equal(t, "", res2)
}

func TestToTitle(t *testing.T) {

	s := "dünyanın ilk borsa yapısı Aizonai kabul edilir"
	s2 := ""

	res1 := strings.ToTitle(s)
	res2 := strings.ToTitle(s2)

	require.Equal(t, "DÜNYANIN ILK BORSA YAPISI AIZONAI KABUL EDILIR", res1)
	require.Equal(t, "", res2)
}

func TestToTitleSpecial(t *testing.T) {

	s := "dünyanın ilk borsa yapısı Aizonai kabul edilir"
	s2 := ""

	res1 := strings.ToTitleSpecial(unicode.AzeriCase, s)
	res2 := strings.ToTitleSpecial(unicode.AzeriCase, s2)

	require.Equal(t, "DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR", res1)
	require.Equal(t, "", res2)
}

func TestToUpper(t *testing.T) {

	s := "test"
	s2 := ""

	res1 := strings.ToUpper(s)
	res2 := strings.ToUpper(s2)

	require.Equal(t, "TEST", res1)
	require.Equal(t, "", res2)
}

func TestToUpperSpecial(t *testing.T) {

	s := "dünyanın ilk borsa yapısı Aizonai kabul edilir"
	s2 := ""

	res1 := strings.ToUpperSpecial(unicode.AzeriCase, s)
	res2 := strings.ToUpperSpecial(unicode.AzeriCase, s2)

	require.Equal(t, "DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR", res1)
	require.Equal(t, "", res2)
}

func TestToValidUTF8(t *testing.T) {

	s := "a\xffb\xfcc\xfa"

	res1 := strings.ToValidUTF8(s, "UNKNOWN")
	res2 := strings.ToValidUTF8(s, "")

	require.Equal(t, "aUNKNOWNbUNKNOWNcUNKNOWN", res1)
	require.Equal(t, "abc", res2)
}

func TestTrim(t *testing.T) {

	s := "     spaces trim    "
	s2 := ";;;;spaces;!trim!!!!"

	res1 := strings.Trim(s, " ")
	res2 := strings.Trim(s2, ";!")

	require.Equal(t, "spaces trim", res1)
	require.Equal(t, "spaces;!trim", res2)
}

func TestTrimFunc(t *testing.T) {

	s := "lowerUppeRlower"

	f := func(r rune) bool {
		return !unicode.IsUpper(r)
	}

	res1 := strings.TrimFunc(s, f)

	require.Equal(t, "UppeR", res1)
}

func TestTrimLeft(t *testing.T) {

	s := "     spaces trimLeft    "
	s2 := ";;;;spaces;!trimLeft!!!!"

	res1 := strings.TrimLeft(s, " ")
	res2 := strings.TrimLeft(s2, ";!")

	require.Equal(t, "spaces trimLeft    ", res1)
	require.Equal(t, "spaces;!trimLeft!!!!", res2)
}

func TestTrimLeftFunc(t *testing.T) {

	s := "lowerUppeRlower"

	f := func(r rune) bool {
		return !unicode.IsUpper(r)
	}

	res1 := strings.TrimLeftFunc(s, f)

	require.Equal(t, "UppeRlower", res1)
}

func TestTrimPrefix(t *testing.T) {

	s := "Bearer 12345"

	res1 := strings.TrimPrefix(s, "Bearer\x20")

	require.Equal(t, "12345", res1)
}

func TestTrimRight(t *testing.T) {

	s := "     spaces trimRight    "
	s2 := ";;;;spaces;!trimRight!!!!"

	res1 := strings.TrimRight(s, " ")
	res2 := strings.TrimRight(s2, ";!")

	require.Equal(t, "     spaces trimRight", res1)
	require.Equal(t, ";;;;spaces;!trimRight", res2)
}

func TestTrimRightFunc(t *testing.T) {

	s := "lowerUppeRlower"

	f := func(r rune) bool {
		return !unicode.IsUpper(r)
	}

	res1 := strings.TrimRightFunc(s, f)

	require.Equal(t, "lowerUppeR", res1)
}

func TestTrimSpace(t *testing.T) {

	s := "     spaces trimSpace    "

	res1 := strings.TrimSpace(s)

	require.Equal(t, "spaces trimSpace", res1)
}

func TestTrimSuffix(t *testing.T) {

	s := "Bearer 12345 Token"

	res1 := strings.TrimSuffix(s, "\x20Token")

	require.Equal(t, "Bearer 12345", res1)
}

func TestBuilder(t *testing.T) {
	var b strings.Builder

	firstCap := b.Cap()

	b.Grow(8)

	secondCap := b.Cap()

	firstLen := b.Len()

	writeLenFirst, _ := b.Write([]byte{'\x66', '\x69', 'r', 's', 't', '\x20'})
	b.WriteByte('\x0a')
	writeLenSecond, _ := b.WriteRune('s')
	b.WriteString("econd")

	secondLen := b.Len()

	res := b.String()

	b.Reset()

	thirdLen := b.Len()
	thirdCap := b.Cap()

	b.Reset()

	require.Equal(t, 0, firstCap)
	require.Equal(t, 0, firstLen)
	require.Equal(t, 8, secondCap)
	require.Equal(t, 6, writeLenFirst)
	require.Equal(t, 1, writeLenSecond)
	require.Equal(t, 13, secondLen)
	require.Equal(t, "first \nsecond", res)
	require.Equal(t, 0, thirdLen)
	require.Equal(t, 0, thirdCap)
}

func TestReader(t *testing.T) {
	// Reader.s is not writable and len anytime will be 0
	// var r *strings.Reader

	val := string([]byte{'\x66', '\x69', 'r', 's', 't', '\x20', 's', 'e', 'c', 'o', 'n', 'd'})

	r := strings.NewReader(val)

	firstLen := r.Len()
	firstSize := r.Size()

	firstByte, _ := r.ReadByte()

	targetBuffer := make([]byte, len(val))

	firstRune, firstRuneSize, _ := r.ReadRune()

	writeLenFirst, _ := r.ReadAt(targetBuffer, 6)

	writeLenThird, _ := r.Read(targetBuffer)

	r.Reset(val)

	targetBufferSecond := make([]byte, 2)

	position, _ := r.Seek(2, io.SeekStart)

	lenFourth := r.Len()
	secondSize := r.Size()

	n, _ := r.Read(targetBufferSecond)

	lenFifth := r.Len()

	_ = r.UnreadByte()
	lenSixth := r.Len()
	_, _, _ = r.ReadRune()
	lenSeventh := r.Len()
	_ = r.UnreadRune()
	lenEigth := r.Len()

	r.Reset(val)
	// _, _ = r.WriteTo(os.Stdout)

	_, _ = r.WriteTo(&bytes.Buffer{})

	require.Equal(t, 12, firstLen)
	require.Equal(t, int64(12), firstSize)
	require.Equal(t, byte('\x66'), firstByte)
	require.Equal(t, '\x69', firstRune)
	require.Equal(t, 1, firstRuneSize)
	require.Equal(t, 6, writeLenFirst)
	require.Equal(t, []byte{'r', 's', 't', '\x20', 's', 'e', 'c', 'o', 'n', 'd', '\x00', '\x00'}, targetBuffer)
	require.Equal(t, 10, writeLenThird)
	require.Equal(t, 10, lenFourth)
	require.Equal(t, int64(12), secondSize)
	require.Equal(t, 2, n)
	require.Equal(t, 8, lenFifth)
	require.Equal(t, []byte{'r', 's'}, targetBufferSecond)
	require.Equal(t, int64(2), position)
	require.Equal(t, 9, lenSixth)
	require.Equal(t, 8, lenSeventh)
	require.Equal(t, 9, lenEigth)
}

func TestReplacer(t *testing.T) {

	r := strings.NewReplacer("apple", "A", "banana", "B", "cake", "C")

	s := strings.ToLower("Apple and banana and cake")

	res := r.Replace(s)

	_, _ = r.WriteString(&bytes.Buffer{}, s)

	require.Equal(t, "A and B and C", res)
}
