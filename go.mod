module wasm-example

go 1.22

toolchain go1.22.0 // https://github.com/golang/go/issues/62278#issuecomment-1693538776

require (
	github.com/goplus/gop v1.2.6
	github.com/goplus/igop v0.25.0
	github.com/goplus/mod v0.13.12
	golang.org/x/tools v0.24.0
)

require (
	github.com/ajstarks/svgo v0.0.0-20210927141636-6d70534b1098 // indirect
	github.com/esimov/stackblur-go v1.0.1-0.20190121110005-00e727e3c7a9 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20220320163800-277f93cfa958 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20200217142428-fce0ec30dd00 // indirect
	github.com/goplus/canvas v0.1.0 // indirect
	github.com/goplus/gogen v1.15.2 // indirect
	github.com/goplus/reflectx v1.2.2 // indirect
	github.com/goplus/spx v1.0.0 // indirect
	github.com/hajimehoshi/ebiten/v2 v2.3.4 // indirect
	github.com/hajimehoshi/go-mp3 v0.3.3 // indirect
	github.com/hajimehoshi/oto/v2 v2.1.0 // indirect
	github.com/jezek/xgb v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/qiniu/audio v0.2.1 // indirect
	github.com/qiniu/x v1.13.10 // indirect
	github.com/srwiley/oksvg v0.0.0-20221011165216-be6e8873101c // indirect
	github.com/srwiley/rasterx v0.0.0-20210519020934-456a8d69b780 // indirect
	github.com/timandy/routine v1.1.1 // indirect
	github.com/visualfc/funcval v0.1.4 // indirect
	github.com/visualfc/gid v0.1.0 // indirect
	github.com/visualfc/goembed v0.3.2 // indirect
	github.com/visualfc/xtype v0.2.0 // indirect
	golang.org/x/exp v0.0.0-20190731235908-ec7cb31e5a56 // indirect
	golang.org/x/image v0.19.0 // indirect
	golang.org/x/mobile v0.0.0-20240806205939-81131f6468ab // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.23.0 // indirect
	golang.org/x/text v0.17.0 // indirect
)

// https://github.com/nighca/mod/tree/web
replace github.com/goplus/mod v0.13.12 => github.com/nighca/mod v0.0.0-20240805065729-b50535825ae2

//https://github.com/callme-taota/gogen/tree/web
//replace github.com/goplus/gogen v1.15.2 => github.com/callme-taota/gogen v0.0.0-20240807020116-1193cc00cd75
//replace github.com/goplus/gogen v1.15.2 => ../../gop/gogen

replace github.com/goplus/igop v0.25.0 => github.com/callme-taota/igop v0.0.3
//replace github.com/goplus/igop v0.25.0 => ../../gop/igop
