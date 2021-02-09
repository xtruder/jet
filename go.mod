module github.com/go-jet/jet/v2

go 1.11

require (
	github.com/bradleyjkemp/cupaloy v2.3.0+incompatible
	github.com/go-jet/jet-test-data v0.0.0-20200524155528-ed53a505eb73
	github.com/go-sql-driver/mysql v1.5.0
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/google/uuid v1.2.0
	github.com/kr/pretty v0.1.0 // indirect
	github.com/lib/pq v1.9.0
	github.com/pkg/profile v1.5.0
	github.com/stretchr/testify v1.7.0
	github.com/xtruder/go-testparrot v0.0.0-20210130125745-5e9be570b589
	gopkg.in/check.v1 v1.0.0-20180628173108-788fd7840127 // indirect
)

replace github.com/go-jet/jet-test-data => /workspace/tests/testdata

replace github.com/stretchr/testify => github.com/posener/testify v1.1.5-0.20200314174129-64d5d85d1fa6
