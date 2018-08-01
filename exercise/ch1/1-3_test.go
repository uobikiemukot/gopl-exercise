package echo

import (
	"io"
	"os"
	"testing"
)

func setup(b *testing.B) ([]string, io.Writer) {
	var data []string = []string{
		"backened",
		"anetholes",
		"aulas",
		"argufying",
		"a-dangle",
		"aribine",
		"affectious",
		"aggerate",
		"antagonisms",
		"adoptianist",
		"anabolite",
		"Ambrosio",
	}

	for i := 0; i < 8; i++ {
		data = append(data, data...)
	}

	devnull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal("cannot open null device")
	}

	return data, devnull
}

func BenchmarkEchoes(b *testing.B) {
	data, w := setup(b)

	b.Run("echo1", func(b *testing.B) { EchoN(b, w, data, echo1) })
	b.Run("echo2", func(b *testing.B) { EchoN(b, w, data, echo2) })
	b.Run("echo3", func(b *testing.B) { EchoN(b, w, data, echo3) })
}

//func EchoN(b *testing.B, w io.Writer, data []string, f func(w io.Writer, data []string)) {
func EchoN(b *testing.B, w io.Writer, data []string, f func(io.Writer, []string)) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		f(w, data)
	}
}
