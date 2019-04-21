module github.com/ubunifupay/transaction

go 1.12

require (
	github.com/MAKOSCAFEE/malengo-pay v0.0.0-20190421104451-a23831b15d53
	github.com/golang/protobuf v1.3.1
	github.com/ubunifupay/balance v0.0.0-20190421135647-20784d982d9f
	golang.org/x/net v0.0.0-20190420063019-afa5a82059c6
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/grpc v1.20.1
)

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190418165655-df01cb2cc480
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190419195159-b8972e603456
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/net => github.com/golang/net v0.0.0-20190420063019-afa5a82059c6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190412183630-56d357773e84
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190419153524-e8e3143a4f4a
	golang.org/x/text => github.com/golang/text v0.3.0
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190420181800-aa740d480789
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190418145605-e7d98fc518a7
	google.golang.org/grpc => github.com/grpc/grpc-go v1.20.1
)
