# Anysrv
Anysrv is a high-performance web server with a clean API.

### Benchmark 
```
#GithubAPI Routes: 203
   Aero: 198944 Bytes
   Srv: 82968 Bytes
   Gin: 59328 Bytes
   HttpRouter: 37144 Bytes

#GPlusAPI Routes: 13
   Aero: 26328 Bytes
   Srv: 7424 Bytes
   Gin: 4464 Bytes
   HttpRouter: 2808 Bytes

#ParseAPI Routes: 26
   Aero: 28472 Bytes
   Srv: 9792 Bytes
   Gin: 7808 Bytes
   HttpRouter: 5072 Bytes

#Static Routes: 157
   Aero: 34536 Bytes
   Srv: 7976 Bytes
   Gin: 34984 Bytes
   HttpRouter: 21712 Bytes

goos: windows
goarch: amd64
pkg: github.com/julienschmidt/go-http-routing-benchmark
BenchmarkAero_Param
BenchmarkAero_Param              	26088033	        41.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_Param
BenchmarkSrv_Param               	31577617	        33.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param
BenchmarkGin_Param               	22655956	        50.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Param5
BenchmarkAero_Param5             	19345945	        61.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_Param5
BenchmarkSrv_Param5              	24487646	        46.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param5
BenchmarkGin_Param5              	12250497	        85.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_Param5
BenchmarkHttpRouter_Param5       	15384674	        74.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Param20
BenchmarkAero_Param20            	 9231194	       123 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_Param20
BenchmarkSrv_Param20             	13335214	        88.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Param20
BenchmarkGin_Param20             	 6779691	       173 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_Param20
BenchmarkHttpRouter_Param20      	 7057279	       174 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParamWrite
BenchmarkAero_ParamWrite         	16903669	        70.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_ParamWrite
BenchmarkSrv_ParamWrite          	18758469	        63.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParamWrite
BenchmarkGin_ParamWrite          	11879787	       103 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParamWrite
BenchmarkHttpRouter_ParamWrite   	15384674	        78.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubStatic
BenchmarkAero_GithubStatic       	27899774	        41.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_GithubStatic
BenchmarkSrv_GithubStatic        	52094185	        21.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubStatic
BenchmarkGin_GithubStatic        	19049100	        62.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GithubStatic
BenchmarkHttpRouter_GithubStatic 	33327129	        35.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GithubAll
BenchmarkAero_GithubAll          	   72705	     16684 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_GithubAll
BenchmarkSrv_GithubAll           	   74535	     15509 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GithubAll
BenchmarkGin_GithubAll           	   51500	     22291 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GithubAll
BenchmarkHttpRouter_GithubAll    	   67416	     17785 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlusStatic
BenchmarkAero_GPlusStatic        	33325371	        35.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_GPlusStatic
BenchmarkSrv_GPlusStatic         	74959708	        15.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusStatic
BenchmarkGin_GPlusStatic         	24490495	        50.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlusStatic
BenchmarkHttpRouter_GPlusStatic  	52255249	        21.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlusParam
BenchmarkAero_GPlusParam         	20676214	        57.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_GPlusParam
BenchmarkSrv_GPlusParam          	22657796	        52.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusParam
BenchmarkGin_GPlusParam          	16439324	        68.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlusParam
BenchmarkHttpRouter_GPlusParam   	20021322	        57.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlus2Params
BenchmarkAero_GPlus2Params       	14110209	        82.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_GPlus2Params
BenchmarkSrv_GPlus2Params        	13480237	        86.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlus2Params
BenchmarkGin_GPlus2Params        	14116334	        84.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlus2Params
BenchmarkHttpRouter_GPlus2Params 	15575312	        74.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_GPlusAll
BenchmarkAero_GPlusAll           	 1449177	       807 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_GPlusAll
BenchmarkSrv_GPlusAll            	 1711742	       676 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_GPlusAll
BenchmarkGin_GPlusAll            	 1235826	       936 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_GPlusAll
BenchmarkHttpRouter_GPlusAll     	 1544367	       772 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParseStatic
BenchmarkAero_ParseStatic        	30757164	        38.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_ParseStatic
BenchmarkSrv_ParseStatic         	47976778	        23.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseStatic
BenchmarkGin_ParseStatic         	23992849	        51.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParseStatic
BenchmarkHttpRouter_ParseStatic  	54423249	        21.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParseParam
BenchmarkAero_ParseParam         	23536240	        49.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_ParseParam
BenchmarkSrv_ParseParam          	23545662	        46.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseParam
BenchmarkGin_ParseParam          	20012139	        59.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParseParam
BenchmarkHttpRouter_ParseParam   	24994948	        48.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_Parse2Params
BenchmarkAero_Parse2Params       	19670905	        60.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_Parse2Params
BenchmarkSrv_Parse2Params        	20681132	        56.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_Parse2Params
BenchmarkGin_Parse2Params        	17406490	        66.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_Parse2Params
BenchmarkHttpRouter_Parse2Params 	19669873	        59.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_ParseAll
BenchmarkAero_ParseAll           	  853605	      1317 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_ParseAll
BenchmarkSrv_ParseAll            	 1000000	      1127 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_ParseAll
BenchmarkGin_ParseAll            	  705715	      1717 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_ParseAll
BenchmarkHttpRouter_ParseAll     	 1023916	      1179 ns/op	       0 B/op	       0 allocs/op
BenchmarkAero_StaticAll
BenchmarkAero_StaticAll          	  134827	      8292 ns/op	       0 B/op	       0 allocs/op
BenchmarkSrv_StaticAll
BenchmarkSrv_StaticAll           	  137894	      8427 ns/op	       0 B/op	       0 allocs/op
BenchmarkGin_StaticAll
BenchmarkGin_StaticAll           	   72290	     16227 ns/op	       0 B/op	       0 allocs/op
BenchmarkHttpRouter_StaticAll
BenchmarkHttpRouter_StaticAll    	  131880	      9168 ns/op	       0 B/op	       0 allocs/op
PASS

Process finished with exit code 0

```