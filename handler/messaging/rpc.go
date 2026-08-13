package messaging

import (
	"drylang/core"
	"net"
	"net/rpc/jsonrpc"
	"strconv"
)

type RPCDummy struct{}

// Very simplified RPC
func BuiltinRpc(v core.VMCore, args []core.Value, line, col int) (core.Value, error) {
	if len(args) < 1 || args[0].Type != core.ValString {
		return core.UnknownValue, v.Errorf("want rpc(method, ...args)")
	}
	method := args[0].Data.(string)

	switch method {
	case "serve":
		// rpc("serve", port)
		if len(args) != 2 || args[1].Type != core.ValNumber {
			return core.UnknownValue, v.Errorf("rpc.serve wants (port)")
		}
		port := int(args[1].Data.(float64))
		
		go func() {
			l, _ := net.Listen("tcp", ":"+strconv.Itoa(port))
			for {
				conn, err := l.Accept()
				if err != nil {
					continue
				}
				go jsonrpc.ServeConn(conn)
			}
		}()
		return core.Value{Type: core.ValBool, Data: true}, nil

	case "call":
		// rpc("call", host, port, remoteMethod, argStr)
		if len(args) != 5 {
			return core.UnknownValue, v.Errorf("rpc.call wants (host, port, remoteMethod, argStr)")
		}
		if args[1].Type != core.ValString || args[2].Type != core.ValNumber || args[3].Type != core.ValString || args[4].Type != core.ValString {
			return core.UnknownValue, v.Errorf("rpc.call args must be (string, number, string, string)")
		}

		host := args[1].Data.(string)
		port := int(args[2].Data.(float64))
		remoteMethod := args[3].Data.(string)
		argStr := args[4].Data.(string)

		client, err := jsonrpc.Dial("tcp", host+":"+strconv.Itoa(port))
		if err != nil {
			return core.UnknownValue, v.Errorf("rpc.call dial error: %s", err)
		}
		defer client.Close()

		var reply string
		err = client.Call(remoteMethod, &argStr, &reply)
		if err != nil {
			return core.UnknownValue, v.Errorf("rpc.call exec error: %s", err)
		}

		return core.Value{Type: core.ValString, Data: reply}, nil

	default:
		return core.UnknownValue, v.Errorf("unknown rpc method: %s", method)
	}
}
