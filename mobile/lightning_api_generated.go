// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: rpc.proto
package lndmobile

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/golang/protobuf/proto"

	"github.com/lightningnetwork/lnd/lnrpc"
)

var (
	bufLightningLis = bufconn.Listen(100)
)

func getLightningConn() (*grpc.ClientConn, func(), error) {
	conn, err := bufLightningLis.Dial()
	if err != nil {
		return nil, nil, err
	}

	clientConn, err := grpc.Dial("",
		grpc.WithDialer(func(target string,
			timeout time.Duration) (net.Conn, error) {
			return conn, nil
		}),
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(10*time.Second),
	)
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	close := func() {
		conn.Close()
	}

	return clientConn, close, nil
}

// getLightningClient returns a client connection to the server listening
// on lis.
func getLightningClient() (lnrpc.LightningClient, func(), error) {
	clientConn, close, err := getLightningConn()
	if err != nil {
		return nil, nil, err
	}
	client := lnrpc.NewLightningClient(clientConn)
	return client, close, nil
}

// WalletBalance returns total unspent outputs(confirmed and unconfirmed), all
// confirmed unspent outputs and all unconfirmed unspent outputs under control
// of the wallet.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func WalletBalance(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.WalletBalanceRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.WalletBalanceRequest)
			return client.WalletBalance(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ChannelBalance returns the total funds available across all open channels
// in satoshis.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ChannelBalance(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ChannelBalanceRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ChannelBalanceRequest)
			return client.ChannelBalance(ctx, r)
		},
	}
	s.start(msg, callback)
}

// GetTransactions returns a list describing all the known transactions
// relevant to the wallet.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func GetTransactions(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.GetTransactionsRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.GetTransactionsRequest)
			return client.GetTransactions(ctx, r)
		},
	}
	s.start(msg, callback)
}

// SendCoins executes a request to send coins to a particular address. Unlike
// SendMany, this RPC call only allows creating a single output at a time. If
// neither target_conf, or sat_per_byte are set, then the internal wallet will
// consult its fee model to determine a fee for the default confirmation
// target.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func SendCoins(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.SendCoinsRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.SendCoinsRequest)
			return client.SendCoins(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ListUnspent returns a list of all utxos spendable by the wallet with a
// number of confirmations between the specified minimum and maximum.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ListUnspent(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ListUnspentRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ListUnspentRequest)
			return client.ListUnspent(ctx, r)
		},
	}
	s.start(msg, callback)
}

// SubscribeTransactions creates a uni-directional stream from the server to
// the client in which any newly discovered transactions relevant to the
// wallet are sent over.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func SubscribeTransactions(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &lnrpc.GetTransactionsRequest{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*lnrpc.GetTransactionsRequest)
			stream, err := client.SubscribeTransactions(ctx, r)
			if err != nil {
				close()
				return nil, nil, err
			}
			return &receiver{
				recv: func() (proto.Message, error) {
					return stream.Recv()
				},
			}, close, nil
		},
	}
	s.start(msg, rStream)
}

// SendMany handles a request for a transaction that creates multiple specified
// outputs in parallel. If neither target_conf, or sat_per_byte are set, then
// the internal wallet will consult its fee model to determine a fee for the
// default confirmation target.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func SendMany(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.SendManyRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.SendManyRequest)
			return client.SendMany(ctx, r)
		},
	}
	s.start(msg, callback)
}

// NewAddress creates a new address under control of the local wallet.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func NewAddress(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.NewAddressRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.NewAddressRequest)
			return client.NewAddress(ctx, r)
		},
	}
	s.start(msg, callback)
}

// SignMessage signs a message with this node's private key. The returned
// signature string is `zbase32` encoded and pubkey recoverable, meaning that
// only the message digest and signature are needed for verification.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func SignMessage(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.SignMessageRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.SignMessageRequest)
			return client.SignMessage(ctx, r)
		},
	}
	s.start(msg, callback)
}

// VerifyMessage verifies a signature over a msg. The signature must be
// zbase32 encoded and signed by an active node in the resident node's
// channel database. In addition to returning the validity of the signature,
// VerifyMessage also returns the recovered pubkey from the signature.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func VerifyMessage(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.VerifyMessageRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.VerifyMessageRequest)
			return client.VerifyMessage(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ConnectPeer attempts to establish a connection to a remote peer. This is at
// the networking level, and is used for communication between nodes. This is
// distinct from establishing a channel with a peer.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ConnectPeer(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ConnectPeerRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ConnectPeerRequest)
			return client.ConnectPeer(ctx, r)
		},
	}
	s.start(msg, callback)
}

// DisconnectPeer attempts to disconnect one peer from another identified by a
// given pubKey. In the case that we currently have a pending or active channel
// with the target peer, then this action will be not be allowed.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func DisconnectPeer(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.DisconnectPeerRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.DisconnectPeerRequest)
			return client.DisconnectPeer(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ListPeers returns a verbose listing of all currently active peers.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ListPeers(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ListPeersRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ListPeersRequest)
			return client.ListPeers(ctx, r)
		},
	}
	s.start(msg, callback)
}

// GetInfo returns general information concerning the lightning node including
// it's identity pubkey, alias, the chains it is connected to, and information
// concerning the number of open+pending channels.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func GetInfo(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.GetInfoRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.GetInfoRequest)
			return client.GetInfo(ctx, r)
		},
	}
	s.start(msg, callback)
}

// PendingChannels returns a list of all the channels that are currently
// considered "pending". A channel is pending if it has finished the funding
// workflow and is waiting for confirmations for the funding txn, or is in the
// process of closure, either initiated cooperatively or non-cooperatively.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func PendingChannels(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.PendingChannelsRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.PendingChannelsRequest)
			return client.PendingChannels(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ListChannels returns a description of all the open channels that this node
// is a participant in.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ListChannels(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ListChannelsRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ListChannelsRequest)
			return client.ListChannels(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ClosedChannels returns a description of all the closed channels that
// this node was a participant in.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ClosedChannels(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ClosedChannelsRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ClosedChannelsRequest)
			return client.ClosedChannels(ctx, r)
		},
	}
	s.start(msg, callback)
}

// OpenChannelSync is a synchronous version of the OpenChannel RPC call. This
// call is meant to be consumed by clients to the REST proxy. As with all
// other sync calls, all byte slices are intended to be populated as hex
// encoded strings.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func OpenChannelSync(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.OpenChannelRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.OpenChannelRequest)
			return client.OpenChannelSync(ctx, r)
		},
	}
	s.start(msg, callback)
}

// OpenChannel attempts to open a singly funded channel specified in the
// request to a remote peer. Users are able to specify a target number of
// blocks that the funding transaction should be confirmed in, or a manual fee
// rate to us for the funding transaction. If neither are specified, then a
// lax block confirmation target is used.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func OpenChannel(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &lnrpc.OpenChannelRequest{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*lnrpc.OpenChannelRequest)
			stream, err := client.OpenChannel(ctx, r)
			if err != nil {
				close()
				return nil, nil, err
			}
			return &receiver{
				recv: func() (proto.Message, error) {
					return stream.Recv()
				},
			}, close, nil
		},
	}
	s.start(msg, rStream)
}

// CloseChannel attempts to close an active channel identified by its channel
// outpoint (ChannelPoint). The actions of this method can additionally be
// augmented to attempt a force close after a timeout period in the case of an
// inactive peer. If a non-force close (cooperative closure) is requested,
// then the user can specify either a target number of blocks until the
// closure transaction is confirmed, or a manual fee rate. If neither are
// specified, then a default lax, block confirmation target is used.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func CloseChannel(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &lnrpc.CloseChannelRequest{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*lnrpc.CloseChannelRequest)
			stream, err := client.CloseChannel(ctx, r)
			if err != nil {
				close()
				return nil, nil, err
			}
			return &receiver{
				recv: func() (proto.Message, error) {
					return stream.Recv()
				},
			}, close, nil
		},
	}
	s.start(msg, rStream)
}

// AbandonChannel removes all channel state from the database except for a
// close summary. This method can be used to get rid of permanently unusable
// channels due to bugs fixed in newer versions of lnd. Only available
// when in debug builds of lnd.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func AbandonChannel(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.AbandonChannelRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.AbandonChannelRequest)
			return client.AbandonChannel(ctx, r)
		},
	}
	s.start(msg, callback)
}

// SendPayment dispatches a bi-directional streaming RPC for sending payments
// through the Lightning Network. A single RPC invocation creates a persistent
// bi-directional stream allowing clients to rapidly send payments through the
// Lightning Network with a single persistent connection.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced. The send stream can accept zero or more requests before it
// is closed.
func SendPayment(rStream RecvStream) (SendStream, error) {
	b := &biStreamHandler{
		newProto: func() proto.Message {
			return &lnrpc.SendRequest{}
		},
		biStream: func(ctx context.Context) (*receiver, *sender, func(), error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, nil, nil, err
			}

			stream, err := client.SendPayment(ctx)
			if err != nil {
				close()
				return nil, nil, nil, err
			}
			return &receiver{
					recv: func() (proto.Message, error) {
						return stream.Recv()
					},
				},
				&sender{
					send: func(req proto.Message) error {
						r := req.(*lnrpc.SendRequest)
						return stream.Send(r)
					},
					close: stream.CloseSend,
				}, close, nil
		},
	}
	return b.start(rStream)
}

// SendPaymentSync is the synchronous non-streaming version of SendPayment.
// This RPC is intended to be consumed by clients of the REST proxy.
// Additionally, this RPC expects the destination's public key and the payment
// hash (if any) to be encoded as hex strings.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func SendPaymentSync(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.SendRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.SendRequest)
			return client.SendPaymentSync(ctx, r)
		},
	}
	s.start(msg, callback)
}

// SendToRoute is a bi-directional streaming RPC for sending payment through
// the Lightning Network. This method differs from SendPayment in that it
// allows users to specify a full route manually. This can be used for things
// like rebalancing, and atomic swaps.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced. The send stream can accept zero or more requests before it
// is closed.
func SendToRoute(rStream RecvStream) (SendStream, error) {
	b := &biStreamHandler{
		newProto: func() proto.Message {
			return &lnrpc.SendToRouteRequest{}
		},
		biStream: func(ctx context.Context) (*receiver, *sender, func(), error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, nil, nil, err
			}

			stream, err := client.SendToRoute(ctx)
			if err != nil {
				close()
				return nil, nil, nil, err
			}
			return &receiver{
					recv: func() (proto.Message, error) {
						return stream.Recv()
					},
				},
				&sender{
					send: func(req proto.Message) error {
						r := req.(*lnrpc.SendToRouteRequest)
						return stream.Send(r)
					},
					close: stream.CloseSend,
				}, close, nil
		},
	}
	return b.start(rStream)
}

// SendToRouteSync is a synchronous version of SendToRoute. It Will block
// until the payment either fails or succeeds.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func SendToRouteSync(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.SendToRouteRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.SendToRouteRequest)
			return client.SendToRouteSync(ctx, r)
		},
	}
	s.start(msg, callback)
}

// AddInvoice attempts to add a new invoice to the invoice database. Any
// duplicated invoices are rejected, therefore all invoices *must* have a
// unique payment preimage.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func AddInvoice(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.Invoice{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.Invoice)
			return client.AddInvoice(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ListInvoices returns a list of all the invoices currently stored within the
// database. Any active debug invoices are ignored. It has full support for
// paginated responses, allowing users to query for specific invoices through
// their add_index. This can be done by using either the first_index_offset or
// last_index_offset fields included in the response as the index_offset of the
// next request. The reversed flag is set by default in order to paginate
// backwards. If you wish to paginate forwards, you must explicitly set the
// flag to false. If none of the parameters are specified, then the last 100
// invoices will be returned.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ListInvoices(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ListInvoiceRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ListInvoiceRequest)
			return client.ListInvoices(ctx, r)
		},
	}
	s.start(msg, callback)
}

// LookupInvoice attempts to look up an invoice according to its payment hash.
// The passed payment hash *must* be exactly 32 bytes, if not, an error is
// returned.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func LookupInvoice(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.PaymentHash{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.PaymentHash)
			return client.LookupInvoice(ctx, r)
		},
	}
	s.start(msg, callback)
}

// SubscribeInvoices returns a uni-directional stream (server -> client) for
// notifying the client of newly added/settled invoices. The caller can
// optionally specify the add_index and/or the settle_index. If the add_index
// is specified, then we'll first start by sending add invoice events for all
// invoices with an add_index greater than the specified value.  If the
// settle_index is specified, the next, we'll send out all settle events for
// invoices with a settle_index greater than the specified value.  One or both
// of these fields can be set. If no fields are set, then we'll only send out
// the latest add/settle events.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func SubscribeInvoices(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &lnrpc.InvoiceSubscription{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*lnrpc.InvoiceSubscription)
			stream, err := client.SubscribeInvoices(ctx, r)
			if err != nil {
				close()
				return nil, nil, err
			}
			return &receiver{
				recv: func() (proto.Message, error) {
					return stream.Recv()
				},
			}, close, nil
		},
	}
	s.start(msg, rStream)
}

// DecodePayReq takes an encoded payment request string and attempts to decode
// it, returning a full description of the conditions encoded within the
// payment request.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func DecodePayReq(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.PayReqString{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.PayReqString)
			return client.DecodePayReq(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ListPayments returns a list of all outgoing payments.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ListPayments(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ListPaymentsRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ListPaymentsRequest)
			return client.ListPayments(ctx, r)
		},
	}
	s.start(msg, callback)
}

// DeleteAllPayments deletes all outgoing payments from DB.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func DeleteAllPayments(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.DeleteAllPaymentsRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.DeleteAllPaymentsRequest)
			return client.DeleteAllPayments(ctx, r)
		},
	}
	s.start(msg, callback)
}

// DescribeGraph returns a description of the latest graph state from the
// point of view of the node. The graph information is partitioned into two
// components: all the nodes/vertexes, and all the edges that connect the
// vertexes themselves.  As this is a directed graph, the edges also contain
// the node directional specific routing policy which includes: the time lock
// delta, fee information, etc.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func DescribeGraph(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ChannelGraphRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ChannelGraphRequest)
			return client.DescribeGraph(ctx, r)
		},
	}
	s.start(msg, callback)
}

// GetChanInfo returns the latest authenticated network announcement for the
// given channel identified by its channel ID: an 8-byte integer which
// uniquely identifies the location of transaction's funding output within the
// blockchain.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func GetChanInfo(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ChanInfoRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ChanInfoRequest)
			return client.GetChanInfo(ctx, r)
		},
	}
	s.start(msg, callback)
}

// GetNodeInfo returns the latest advertised, aggregated, and authenticated
// channel information for the specified node identified by its public key.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func GetNodeInfo(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.NodeInfoRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.NodeInfoRequest)
			return client.GetNodeInfo(ctx, r)
		},
	}
	s.start(msg, callback)
}

// QueryRoutes attempts to query the daemon's Channel Router for a possible
// route to a target destination capable of carrying a specific amount of
// satoshis. The retuned route contains the full details required to craft and
// send an HTLC, also including the necessary information that should be
// present within the Sphinx packet encapsulated within the HTLC.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func QueryRoutes(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.QueryRoutesRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.QueryRoutesRequest)
			return client.QueryRoutes(ctx, r)
		},
	}
	s.start(msg, callback)
}

// GetNetworkInfo returns some basic stats about the known channel graph from
// the point of view of the node.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func GetNetworkInfo(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.NetworkInfoRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.NetworkInfoRequest)
			return client.GetNetworkInfo(ctx, r)
		},
	}
	s.start(msg, callback)
}

// StopDaemon will send a shutdown request to the interrupt handler, triggering
// a graceful shutdown of the daemon.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func StopDaemon(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.StopRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.StopRequest)
			return client.StopDaemon(ctx, r)
		},
	}
	s.start(msg, callback)
}

// SubscribeChannelGraph launches a streaming RPC that allows the caller to
// receive notifications upon any changes to the channel graph topology from
// the point of view of the responding node. Events notified include: new
// nodes coming online, nodes updating their authenticated attributes, new
// channels being advertised, updates in the routing policy for a directional
// channel edge, and when channels are closed on-chain.
//
// NOTE: This method produces a stream of responses, and the receive stream can
// be called zero or more times. After EOF error is returned, no more responses
// will be produced.
func SubscribeChannelGraph(msg []byte, rStream RecvStream) {
	s := &readStreamHandler{
		newProto: func() proto.Message {
			return &lnrpc.GraphTopologySubscription{}
		},
		recvStream: func(ctx context.Context,
			req proto.Message) (*receiver, func(), error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, nil, err
			}

			r := req.(*lnrpc.GraphTopologySubscription)
			stream, err := client.SubscribeChannelGraph(ctx, r)
			if err != nil {
				close()
				return nil, nil, err
			}
			return &receiver{
				recv: func() (proto.Message, error) {
					return stream.Recv()
				},
			}, close, nil
		},
	}
	s.start(msg, rStream)
}

// DebugLevel allows a caller to programmatically set the logging verbosity of
// lnd. The logging can be targeted according to a coarse daemon-wide logging
// level, or in a granular fashion to specify the logging for a target
// sub-system.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func DebugLevel(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.DebugLevelRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.DebugLevelRequest)
			return client.DebugLevel(ctx, r)
		},
	}
	s.start(msg, callback)
}

// FeeReport allows the caller to obtain a report detailing the current fee
// schedule enforced by the node globally for each channel.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func FeeReport(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.FeeReportRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.FeeReportRequest)
			return client.FeeReport(ctx, r)
		},
	}
	s.start(msg, callback)
}

// UpdateChannelPolicy allows the caller to update the fee schedule and
// channel policies for all channels globally, or a particular channel.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func UpdateChannelPolicy(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.PolicyUpdateRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.PolicyUpdateRequest)
			return client.UpdateChannelPolicy(ctx, r)
		},
	}
	s.start(msg, callback)
}

// ForwardingHistory allows the caller to query the htlcswitch for a record of
// all HTLC's forwarded within the target time range, and integer offset
// within that time range. If no time-range is specified, then the first chunk
// of the past 24 hrs of forwarding history are returned.
// 
// A list of forwarding events are returned. The size of each forwarding event
// is 40 bytes, and the max message size able to be returned in gRPC is 4 MiB.
// As a result each message can only contain 50k entries.  Each response has
// the index offset of the last entry. The index offset can be provided to the
// request to allow the caller to skip a series of records.
//
// NOTE: This method produces a single result or error, and the callback will
// be called only once.
func ForwardingHistory(msg []byte, callback Callback) {
	s := &syncHandler{
		newProto: func() proto.Message {
			return &lnrpc.ForwardingHistoryRequest{}
		},
		getSync: func(ctx context.Context,
			req proto.Message) (proto.Message, error) {

			// Get the gRPC client.
			client, close, err := getLightningClient()
			if err != nil {
				return nil, err
			}
			defer close()

			r := req.(*lnrpc.ForwardingHistoryRequest)
			return client.ForwardingHistory(ctx, r)
		},
	}
	s.start(msg, callback)
}
