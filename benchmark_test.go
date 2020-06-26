package exiftool

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

var (
	dir        = "../test/img/"
	benchmarks = []struct {
		name     string
		fileName string
	}{
		{".CR2/GPS", "2.CR2"},
		{".CR2/7D", "7D2.CR2"},
		{".CR3", "1.CR3"},
		{".JPG/GPS", "17.jpg"},
		{".HEIC", "1.heic"},
		{".GoPro/6", "hero6.jpg"},
		{".NEF/Nikon", "2.NEF"},
		{".ARW/Sony", "2.ARW"},
		{".DNG/Adobe", "1.DNG"},
		{".JPG/NoExif", "20.jpg"},
	}

	sampleDir     = "samples/"
	testFilenames = []string{
		"honor20.jpg",
		"hero6.jpg",
		"1.CR2",
		"3.CR2",
		"350D.CR2",
		"XT1.CR2",
		"60D.CR2",
		"6D.CR2",
		"7D.CR2",
		"90D.cr3",
		"2.CR3",
		"1.CR3",
		"1.jpg",
		"2.jpg",
		"1.NEF",
		"2.NEF",
		"3.NEF",
		"1.ARW",
		"2.ARW",
		"4.RW2",
		"hero6.gpr",
		"4.webp",
		"20.jpg",
	}
)

func BenchmarkParseExif100(b *testing.B) {
	for _, bm := range testFilenames {
		b.Run(bm, func(b *testing.B) {
			f, err := os.Open(sampleDir + bm)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			buf, _ := ioutil.ReadAll(f)
			cb := bytes.NewReader(buf)

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err = ParseExif(cb)
				if err != nil {
					if err != ErrNoExif {
						b.Fatal(err)
					}
				}
			}
		})
	}
}

//BenchmarkParseExif100/honor20.jpg-8         	   41763	     26779 ns/op	    8791 B/op	      38 allocs/op
//BenchmarkParseExif100/hero6.jpg-8           	   56798	     21478 ns/op	    8036 B/op	      31 allocs/op
//BenchmarkParseExif100/1.CR2-8               	   25717	     43625 ns/op	   13215 B/op	      57 allocs/op
//BenchmarkParseExif100/3.CR2-8               	   31881	     37863 ns/op	   12144 B/op	      53 allocs/op
//BenchmarkParseExif100/350D.CR2-8            	   39760	     30335 ns/op	   10445 B/op	      46 allocs/op
//BenchmarkParseExif100/XT1.CR2-8             	   39951	     30473 ns/op	   10443 B/op	      46 allocs/op
//BenchmarkParseExif100/60D.CR2-8             	   28388	     41794 ns/op	   12592 B/op	      52 allocs/op
//BenchmarkParseExif100/6D.CR2-8              	   26862	     43931 ns/op	   13182 B/op	      57 allocs/op
//BenchmarkParseExif100/7D.CR2-8              	   27742	     43892 ns/op	   13215 B/op	      57 allocs/op
//BenchmarkParseExif100/90D.cr3-8             	  152034	      8243 ns/op	    5158 B/op	      17 allocs/op
//BenchmarkParseExif100/2.CR3-8               	  148568	      8020 ns/op	    5157 B/op	      17 allocs/op
//BenchmarkParseExif100/1.CR3-8               	  152665	      8190 ns/op	    5158 B/op	      17 allocs/op
//BenchmarkParseExif100/1.jpg-8               	   65166	     18525 ns/op	    6755 B/op	      30 allocs/op
//BenchmarkParseExif100/2.jpg-8               	   52927	     22747 ns/op	    8421 B/op	      33 allocs/op
//BenchmarkParseExif100/1.NEF-8               	   23182	     47926 ns/op	   13600 B/op	      61 allocs/op
//BenchmarkParseExif100/2.NEF-8               	   22272	     51129 ns/op	   16674 B/op	      63 allocs/op
//BenchmarkParseExif100/3.NEF-8               	   20660	     58233 ns/op	   17008 B/op	      67 allocs/op
//BenchmarkParseExif100/1.ARW-8               	   31310	     38229 ns/op	   11930 B/op	      56 allocs/op
//BenchmarkParseExif100/2.ARW-8               	   30970	     38146 ns/op	   11932 B/op	      57 allocs/op
//BenchmarkParseExif100/4.RW2-8               	   36390	     32900 ns/op	    8199 B/op	      31 allocs/op
//BenchmarkParseExif100/hero6.gpr-8           	   33552	     35745 ns/op	   13603 B/op	      39 allocs/op
//BenchmarkParseExif100/4.webp-8              	     826	   1536155 ns/op	    4432 B/op	       5 allocs/op
//BenchmarkParseExif100/20.jpg-8              	     472	   2536924 ns/op	    4432 B/op	       5 allocs/op

func BenchmarkScanExif100(b *testing.B) {
	for _, bm := range testFilenames {
		b.Run(bm, func(b *testing.B) {
			f, err := os.Open(sampleDir + bm)
			if err != nil {
				panic(err)
			}
			defer f.Close()
			buf, _ := ioutil.ReadAll(f)
			cb := bytes.NewReader(buf)

			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_, err = ScanExif(cb)
				if err != nil {
					if err != ErrNoExif {
						b.Fatal(err)
					}
				}
			}
		})
	}
}

//BenchmarkScanExif100/honor20.jpg-8         	   36843	     29699 ns/op	   17050 B/op	      40 allocs/op
//BenchmarkScanExif100/hero6.jpg-8           	   46316	     25988 ns/op	   35238 B/op	      33 allocs/op
//BenchmarkScanExif100/1.CR2-8               	   24404	     48578 ns/op	   13212 B/op	      57 allocs/op
//BenchmarkScanExif100/3.CR2-8               	   28908	     40710 ns/op	   12145 B/op	      53 allocs/op
//BenchmarkScanExif100/350D.CR2-8            	   39280	     31519 ns/op	   10445 B/op	      46 allocs/op
//BenchmarkScanExif100/XT1.CR2-8             	   37731	     31582 ns/op	   10444 B/op	      46 allocs/op
//BenchmarkScanExif100/60D.CR2-8             	   27439	     43459 ns/op	   12593 B/op	      52 allocs/op
//BenchmarkScanExif100/6D.CR2-8              	   26264	     45286 ns/op	   13185 B/op	      57 allocs/op
//BenchmarkScanExif100/7D.CR2-8              	   26625	     46062 ns/op	   13216 B/op	      57 allocs/op
//BenchmarkScanExif100/90D.cr3-8             	  131457	      8244 ns/op	    5157 B/op	      17 allocs/op
//BenchmarkScanExif100/2.CR3-8               	  149314	      8345 ns/op	    5157 B/op	      17 allocs/op
//BenchmarkScanExif100/1.CR3-8               	  138854	      8470 ns/op	    5157 B/op	      17 allocs/op
//BenchmarkScanExif100/1.jpg-8               	   52980	     22424 ns/op	   31394 B/op	      32 allocs/op
//BenchmarkScanExif100/2.jpg-8               	   49075	     23813 ns/op	   16679 B/op	      35 allocs/op
//BenchmarkScanExif100/1.NEF-8               	   24420	     50230 ns/op	   13598 B/op	      61 allocs/op
//BenchmarkScanExif100/2.NEF-8               	   22437	     53125 ns/op	   16671 B/op	      62 allocs/op
//BenchmarkScanExif100/3.NEF-8               	   20294	     58299 ns/op	   17008 B/op	      67 allocs/op
//BenchmarkScanExif100/1.ARW-8               	   30277	     39593 ns/op	   11928 B/op	      56 allocs/op
//BenchmarkScanExif100/2.ARW-8               	   29332	     40165 ns/op	   11932 B/op	      56 allocs/op
//BenchmarkScanExif100/4.RW2-8               	   34719	     34740 ns/op	    8202 B/op	      31 allocs/op
//BenchmarkScanExif100/hero6.gpr-8           	   31630	     38285 ns/op	   13606 B/op	      39 allocs/op
//
//BenchmarkScanExif100/20.jpg-8              	 1249808	       996 ns/op	    4496 B/op	       5 allocs/op
