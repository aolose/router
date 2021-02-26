package anysrv

type reqPath struct {
	length int
	deep   int
	start  *[128]int
	end    *[128]int
}
