# Anysrv

Anysrv is a high-performance web server with a clean API.

### Benchmark

Anysrv vs Aero vs Gin vs HttpRouter

```
GOROOT=C:\Program Files\Go #gosetup
GOPATH=C:\Users\ufota\go #gosetup
"C:\Program Files\Go\bin\go.exe" test -c -o C:\Users\ufota\AppData\Local\Temp\___gobench_go_http_routing_benchmark.exe go-http-routing-benchmark #gosetup
C:\Users\ufota\AppData\Local\Temp\___gobench_go_http_routing_benchmark.exe -test.v -test.bench . -test.run ^$ #gosetup
#GithubAPI Routes: 203
   Aero: 194544 Bytes
   Anysrv: 76120 Bytes
   Gin: 59328 Bytes
   HttpRouter: 37144 Bytes

#GPlusAPI Routes: 13
   Aero: 26040 Bytes
   Anysrv: 6944 Bytes
   Gin: 4464 Bytes
   HttpRouter: 2808 Bytes

#ParseAPI Routes: 26
   Aero: 28488 Bytes
   Anysrv: 9128 Bytes
   Gin: 7808 Bytes
   HttpRouter: 5072 Bytes

#Static Routes: 157
   Aero: 34536 Bytes
   Anysrv: 7976 Bytes
   Gin: 34984 Bytes
   HttpRouter: 21712 Bytes

goos: windows
goarch: amd64
pkg: go-http-routing-benchmark
BenchmarkAero_Param
BenchmarkAero_Param              	28569115	        40.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_Param
BenchmarkAnysrv_Param            	35272020	        33.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param
BenchmarkGin_Param               	23074482	        54.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Param5
BenchmarkAero_Param5             	19355119	        60.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_Param5
BenchmarkAnysrv_Param5           	23055109	        47.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param5
BenchmarkGin_Param5              	13947454	        85.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_Param5
BenchmarkHttpRouter_Param5       	15999061	        73.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Param20
BenchmarkAero_Param20            	 9602326	       123 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_Param20
BenchmarkAnysrv_Param20          	13185435	        86.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param20
BenchmarkGin_Param20             	 6451747	       176 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_Param20
BenchmarkHttpRouter_Param20      	 7741945	       162 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParamWrite
BenchmarkAero_ParamWrite         	16894150	        69.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_ParamWrite
BenchmarkAnysrv_ParamWrite       	20014342	        58.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParamWrite
BenchmarkGin_ParamWrite          	11764578	       104 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParamWrite
BenchmarkHttpRouter_ParamWrite   	16438310	        70.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubStatic
BenchmarkAero_GithubStatic       	30000300	        41.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_GithubStatic
BenchmarkAnysrv_GithubStatic     	49994167	        22.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubStatic
BenchmarkGin_GithubStatic        	19364582	        64.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GithubStatic
BenchmarkHttpRouter_GithubStatic 	27911325	        40.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubAll
BenchmarkAero_GithubAll          	   59113	     17459 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_GithubAll
BenchmarkAnysrv_GithubAll        	   72266	     16384 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll
BenchmarkGin_GithubAll           	   52856	     22305 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GithubAll
BenchmarkHttpRouter_GithubAll    	   67410	     17208 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlusStatic
BenchmarkAero_GPlusStatic        	36347334	        34.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_GPlusStatic
BenchmarkAnysrv_GPlusStatic      	74922267	        15.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusStatic
BenchmarkGin_GPlusStatic         	22235521	        51.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlusStatic
BenchmarkHttpRouter_GPlusStatic  	54511755	        22.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlusParam
BenchmarkAero_GPlusParam         	20354541	        56.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_GPlusParam
BenchmarkAnysrv_GPlusParam       	22203430	        54.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusParam
BenchmarkGin_GPlusParam          	17390648	        68.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlusParam
BenchmarkHttpRouter_GPlusParam   	18748417	        61.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlus2Params
BenchmarkAero_GPlus2Params       	14992989	        79.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_GPlus2Params
BenchmarkAnysrv_GPlus2Params     	13043946	        88.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlus2Params
BenchmarkGin_GPlus2Params        	13476331	        86.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlus2Params
BenchmarkHttpRouter_GPlus2Params 	14812620	        77.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlusAll
BenchmarkAero_GPlusAll           	 1498042	       823 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_GPlusAll
BenchmarkAnysrv_GPlusAll         	 1685386	       700 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusAll
BenchmarkGin_GPlusAll            	 1233279	       967 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlusAll
BenchmarkHttpRouter_GPlusAll     	 1524744	       802 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParseStatic
BenchmarkAero_ParseStatic        	32423581	        36.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_ParseStatic
BenchmarkAnysrv_ParseStatic      	49963152	        24.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseStatic
BenchmarkGin_ParseStatic         	21812906	        53.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParseStatic
BenchmarkHttpRouter_ParseStatic  	54540991	        20.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParseParam
BenchmarkAero_ParseParam         	23996112	        50.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_ParseParam
BenchmarkAnysrv_ParseParam       	26063385	        46.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseParam
BenchmarkGin_ParseParam          	21060464	        59.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParseParam
BenchmarkHttpRouter_ParseParam   	25022362	        48.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Parse2Params
BenchmarkAero_Parse2Params       	20689369	        58.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_Parse2Params
BenchmarkAnysrv_Parse2Params     	21424974	        55.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Parse2Params
BenchmarkGin_Parse2Params        	17144007	        69.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_Parse2Params
BenchmarkHttpRouter_Parse2Params 	19364426	        60.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParseAll
BenchmarkAero_ParseAll           	  857246	      1321 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_ParseAll
BenchmarkAnysrv_ParseAll         	 1000000	      1171 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseAll
BenchmarkGin_ParseAll            	  705922	      1673 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParseAll
BenchmarkHttpRouter_ParseAll     	 1000000	      1133 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_StaticAll
BenchmarkAero_StaticAll          	  141168	      8338 ns/op	       0 B/op	       0 allocs/op
BenchmarkAnysrv_StaticAll
BenchmarkAnysrv_StaticAll        	  137895	      8521 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_StaticAll
BenchmarkGin_StaticAll           	   70582	     16336 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_StaticAll
BenchmarkHttpRouter_StaticAll    	  124977	      9361 ns/op	       0 B/op	       0 allocs/op
PASS

Process finished with exit code 0

```