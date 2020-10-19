// Written by Gon Yi

package chubbylog_test

import (
	"github.com/orangenumber/chubbylog"
	"os"
	"testing"
)

// =====================================================================================================================
// A CHUBBY LOGGER
// =====================================================================================================================
func Benchmark_ChubbyLog_Printf(b *testing.B) {
	b.StartTimer()
	out, err := os.Create("./tmp/chubby.txt")
	if err != nil {
		println(err.Error())
	}
	x := chubbylog.New(out, chubbylog.F_DATE|chubbylog.F_TIME|chubbylog.F_PREFIX)

	for i := 0; i < b.N; i++ {
		switch i % 3 {
		case 0:
			x.Infof("info it..: %d\n", i)
		case 1:
			x.Warnf("warn it..: %d\n", i)
		case 2:
			x.Errorf("error it..: %d\n", i)
		}
	}
	b.StopTimer()
	b.ReportAllocs()
}
func Benchmark_ChubbyLog_Printf_Buf(b *testing.B) {
	b.StartTimer()
	outInfo, _ := os.Create("./tmp/chubby_info_buf.txt")
	outWarn, _ := os.Create("./tmp/chubby_warn_buf.txt")
	outError, _ := os.Create("./tmp/chubby_error_buf.txt")

	x := chubbylog.New(nil, chubbylog.F_STD|chubbylog.F_USE_BUF_2K)
	x.GetInfo().SetOutput(outInfo)
	x.GetWarn().SetOutput(outWarn)
	x.GetError().SetOutput(outError)
	x.GetFatal().SetFlag(chubbylog.F_STD)

	for i := 0; i < b.N; i++ {
		switch i % 3 {
		case 0:
			x.Infof("info it..: %d\n", i)
		case 1:
			x.Warnf("warn it..: %d\n", i)
		case 2:
			x.Errorf("error it..: %d\n", i)
		}
	}
	b.StopTimer()
	b.ReportAllocs()
}
