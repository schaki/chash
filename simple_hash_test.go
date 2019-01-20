package simplehash

import (
	"fmt"
	"testing"
)

func TestStringRetrieval(t *testing.T) {
	cases := []struct {
		Contents []string
	}{
		{[]string{""}},
		{[]string{"some content", "more contents"}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			d := Data{}
			for _, content := range tc.Contents {
				k, err := d.Put(content)
				if err != nil {
					t.Fatalf("failed putting content %s\n", content)
				}
				c := d.Get(k)
				if c != content {
					t.Fatalf("content does not match\n%s\n%s\n", c, content)
				}
			}
		})
	}
}

func TestIntRetrieval(t *testing.T) {
	cases := []struct {
		Contents []int
	}{
		{[]int{0}},
		{[]int{1, 2, 3}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			d := Data{}
			for _, content := range tc.Contents {
				k, err := d.Put(content)
				if err != nil {
					t.Fatalf("failed putting content %d\n", content)
				}
				c := d.Get(k)
				if c != content {
					t.Fatalf("content does not match\n%d\n%d\n", c, content)
				}
			}
		})
	}
}

func TestSliceRetrieval(t *testing.T) {
	cases := []struct {
		Contents [][]int
	}{
		{[][]int{{10, 100}}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			d := Data{}
			for _, content := range tc.Contents {
				k, err := d.Put(content)
				if err != nil {
					t.Fatalf("failed putting content %d %s\n", content, err)
				}
				c := d.Get(k)
				if len(c.([]int)) != len(content) {
					t.Fatalf("content does not match\n%d\n%d\n", c, content)
				}
			}
		})
	}
}

type TestStruct struct {
	S     string
	I     int
	Items []string
}

func TestStructRetrieval(t *testing.T) {
	ts := &TestStruct{
		S:     "test",
		I:     100,
		Items: []string{"cheese", "burgers", "and", "fries"},
	}
	cases := []struct {
		Contents []TestStruct
	}{
		{[]TestStruct{*ts}},
	}

	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			d := Data{}
			for _, content := range tc.Contents {
				k, err := d.Put(content)
				if err != nil {
					t.Fatalf("failed putting content %v %s\n", content, err)
				}
				c := d.Get(k)
				if c.(TestStruct).S != ts.S {
					t.Fatalf("content does not match\n%v\n%v\n", c, content)
				}
			}
		})
	}
}

func ensureData(s *int) (Data, []uint64) {
	d := Data{}
	i := 100
	r := []uint64{}
	if s != nil {
		i = *s
	}
	for n := 0; n < i; n++ {
		k, _ := d.Put(fmt.Sprintf("%d", n))
		r = append(r, k)
	}
	return d, r
}

func BenchmarkPut(b *testing.B) {
	d := Data{}
	for n := 0; n < b.N; n++ {
		d.Put(fmt.Sprintf("%d", n))
	}
}

func benchmarkGet(i int, b *testing.B) {
	d, keys := ensureData(&i)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		d.Get(keys[n%i])
	}
}

func BenchmarkGet10(b *testing.B)   { benchmarkGet(10, b) }
func BenchmarkGet100(b *testing.B)  { benchmarkGet(100, b) }
func BenchmarkGet1000(b *testing.B) { benchmarkGet(1000, b) }
