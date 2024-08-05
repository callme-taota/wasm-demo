module wasm-example

go 1.22

toolchain go1.22.0 // https://github.com/golang/go/issues/62278#issuecomment-1693538776

require (
	github.com/goplus/gogen v1.15.2
	github.com/goplus/gop v1.2.6
	github.com/goplus/igop v0.25.0
	github.com/goplus/mod v0.13.12
)

require (
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/goplus/reflectx v1.2.2 // indirect
	github.com/qiniu/x v1.13.10 // indirect
	github.com/timandy/routine v1.1.1 // indirect
	github.com/visualfc/funcval v0.1.4 // indirect
	github.com/visualfc/gid v0.1.0 // indirect
	github.com/visualfc/goembed v0.3.2 // indirect
	github.com/visualfc/xtype v0.2.0 // indirect
	golang.org/x/mod v0.19.0 // indirect
	golang.org/x/tools v0.19.0 // indirect
)

// https://github.com/nighca/mod/tree/web
replace github.com/goplus/mod v0.13.12 => github.com/nighca/mod v0.0.0-20240805065729-b50535825ae2
