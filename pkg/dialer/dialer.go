package dialer

import (
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type DialFunc = func(context.Context, string) (net.Conn, error)

// NewDialer returns a DialFunc that can be used in grpc.WithContextDialer to
func NewDialer() DialFunc {
	return func(ctx context.Context, target string) (net.Conn, error) {
		dialer := &net.Dialer{}
		net, addr := parseDialTarget(target)

		deadline, ok := ctx.Deadline()
		if ok {
			grpclog.Infof("Dialing %s with timeout %s", target, time.Until(deadline))
		} else {
			grpclog.Infof("Dialing %s without deadline", target)
		}

		conn, err := dialer.DialContext(ctx, net, addr)
		if err != nil {
			grpclog.Warningf("Dial %s failed: %s", target, err)
			return nil, err
		}
		grpclog.Infof("Dial %s successfully connected to: %s", target, conn.RemoteAddr())
		return conn, nil
	}
}

const grpcUA = "grpc-go/" + grpc.Version
