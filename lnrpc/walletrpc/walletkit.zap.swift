//
// DO NOT EDIT.
//
// Generated by `protoc-gen-zap`.
// Source: walletrpc/walletkit.proto
//

#if !REMOTEONLY
import Lndmobile
#endif
import Logger

// MARK: - WalletKit
protocol WalletKitConnection {
  func deriveNextKey(_ request: Walletrpc_KeyReq, completion: @escaping ApiCompletion<Signrpc_KeyDescriptor>)
  func deriveKey(_ request: Signrpc_KeyLocator, completion: @escaping ApiCompletion<Signrpc_KeyDescriptor>)
  func nextAddr(_ request: Walletrpc_AddrRequest, completion: @escaping ApiCompletion<Walletrpc_AddrResponse>)
  func publishTransaction(_ request: Walletrpc_Transaction, completion: @escaping ApiCompletion<Walletrpc_PublishResponse>)
  func sendOutputs(_ request: Walletrpc_SendOutputsRequest, completion: @escaping ApiCompletion<Walletrpc_SendOutputsResponse>)
  func estimateFee(_ request: Walletrpc_EstimateFeeRequest, completion: @escaping ApiCompletion<Walletrpc_EstimateFeeResponse>)
  func pendingSweeps(_ request: Walletrpc_PendingSweepsRequest, completion: @escaping ApiCompletion<Walletrpc_PendingSweepsResponse>)
  func bumpFee(_ request: Walletrpc_BumpFeeRequest, completion: @escaping ApiCompletion<Walletrpc_BumpFeeResponse>)
}

#if !REMOTEONLY
class StreamingWalletKitConnection: WalletKitConnection {
  func deriveNextKey(_ request: Walletrpc_KeyReq, completion: @escaping ApiCompletion<Signrpc_KeyDescriptor>) {
    LndmobileDeriveNextKey(try? request.serializedData(), LndCallback(completion))
  }
  func deriveKey(_ request: Signrpc_KeyLocator, completion: @escaping ApiCompletion<Signrpc_KeyDescriptor>) {
    LndmobileDeriveKey(try? request.serializedData(), LndCallback(completion))
  }
  func nextAddr(_ request: Walletrpc_AddrRequest, completion: @escaping ApiCompletion<Walletrpc_AddrResponse>) {
    LndmobileNextAddr(try? request.serializedData(), LndCallback(completion))
  }
  func publishTransaction(_ request: Walletrpc_Transaction, completion: @escaping ApiCompletion<Walletrpc_PublishResponse>) {
    LndmobilePublishTransaction(try? request.serializedData(), LndCallback(completion))
  }
  func sendOutputs(_ request: Walletrpc_SendOutputsRequest, completion: @escaping ApiCompletion<Walletrpc_SendOutputsResponse>) {
    LndmobileSendOutputs(try? request.serializedData(), LndCallback(completion))
  }
  func estimateFee(_ request: Walletrpc_EstimateFeeRequest, completion: @escaping ApiCompletion<Walletrpc_EstimateFeeResponse>) {
    LndmobileEstimateFee(try? request.serializedData(), LndCallback(completion))
  }
  func pendingSweeps(_ request: Walletrpc_PendingSweepsRequest, completion: @escaping ApiCompletion<Walletrpc_PendingSweepsResponse>) {
    LndmobilePendingSweeps(try? request.serializedData(), LndCallback(completion))
  }
  func bumpFee(_ request: Walletrpc_BumpFeeRequest, completion: @escaping ApiCompletion<Walletrpc_BumpFeeResponse>) {
    LndmobileBumpFee(try? request.serializedData(), LndCallback(completion))
  }
}
#endif

final class RPCWalletKitConnection: WalletKitConnection {
  let service: Walletrpc_WalletKitService
      
  public init(configuration: RPCCredentials) {
      service = Walletrpc_WalletKitServiceClient(configuration: configuration)
  }

  func deriveNextKey(_ request: Walletrpc_KeyReq, completion: @escaping ApiCompletion<Signrpc_KeyDescriptor>) {
    _ = try? service.deriveNextKey(request, completion: createHandler(completion))
  }

  func deriveKey(_ request: Signrpc_KeyLocator, completion: @escaping ApiCompletion<Signrpc_KeyDescriptor>) {
    _ = try? service.deriveKey(request, completion: createHandler(completion))
  }

  func nextAddr(_ request: Walletrpc_AddrRequest, completion: @escaping ApiCompletion<Walletrpc_AddrResponse>) {
    _ = try? service.nextAddr(request, completion: createHandler(completion))
  }

  func publishTransaction(_ request: Walletrpc_Transaction, completion: @escaping ApiCompletion<Walletrpc_PublishResponse>) {
    _ = try? service.publishTransaction(request, completion: createHandler(completion))
  }

  func sendOutputs(_ request: Walletrpc_SendOutputsRequest, completion: @escaping ApiCompletion<Walletrpc_SendOutputsResponse>) {
    _ = try? service.sendOutputs(request, completion: createHandler(completion))
  }

  func estimateFee(_ request: Walletrpc_EstimateFeeRequest, completion: @escaping ApiCompletion<Walletrpc_EstimateFeeResponse>) {
    _ = try? service.estimateFee(request, completion: createHandler(completion))
  }

  func pendingSweeps(_ request: Walletrpc_PendingSweepsRequest, completion: @escaping ApiCompletion<Walletrpc_PendingSweepsResponse>) {
    _ = try? service.pendingSweeps(request, completion: createHandler(completion))
  }

  func bumpFee(_ request: Walletrpc_BumpFeeRequest, completion: @escaping ApiCompletion<Walletrpc_BumpFeeResponse>) {
    _ = try? service.bumpFee(request, completion: createHandler(completion))
  }

}

class MockWalletKitConnection: WalletKitConnection {
  private let deriveNextKey: Signrpc_KeyDescriptor?
  private let deriveKey: Signrpc_KeyDescriptor?
  private let nextAddr: Walletrpc_AddrResponse?
  private let publishTransaction: Walletrpc_PublishResponse?
  private let sendOutputs: Walletrpc_SendOutputsResponse?
  private let estimateFee: Walletrpc_EstimateFeeResponse?
  private let pendingSweeps: Walletrpc_PendingSweepsResponse?
  private let bumpFee: Walletrpc_BumpFeeResponse?

  init(
    deriveNextKey: Signrpc_KeyDescriptor? = nil,
    deriveKey: Signrpc_KeyDescriptor? = nil,
    nextAddr: Walletrpc_AddrResponse? = nil,
    publishTransaction: Walletrpc_PublishResponse? = nil,
    sendOutputs: Walletrpc_SendOutputsResponse? = nil,
    estimateFee: Walletrpc_EstimateFeeResponse? = nil,
    pendingSweeps: Walletrpc_PendingSweepsResponse? = nil,
    bumpFee: Walletrpc_BumpFeeResponse? = nil
  ) {
    self.deriveNextKey = deriveNextKey
    self.deriveKey = deriveKey
    self.nextAddr = nextAddr
    self.publishTransaction = publishTransaction
    self.sendOutputs = sendOutputs
    self.estimateFee = estimateFee
    self.pendingSweeps = pendingSweeps
    self.bumpFee = bumpFee
  }
  func deriveNextKey(_ request: Walletrpc_KeyReq, completion: @escaping ApiCompletion<Signrpc_KeyDescriptor>) {
    completion(Result(value: deriveNextKey, error: LndApiError.unknownError))
  }
  func deriveKey(_ request: Signrpc_KeyLocator, completion: @escaping ApiCompletion<Signrpc_KeyDescriptor>) {
    completion(Result(value: deriveKey, error: LndApiError.unknownError))
  }
  func nextAddr(_ request: Walletrpc_AddrRequest, completion: @escaping ApiCompletion<Walletrpc_AddrResponse>) {
    completion(Result(value: nextAddr, error: LndApiError.unknownError))
  }
  func publishTransaction(_ request: Walletrpc_Transaction, completion: @escaping ApiCompletion<Walletrpc_PublishResponse>) {
    completion(Result(value: publishTransaction, error: LndApiError.unknownError))
  }
  func sendOutputs(_ request: Walletrpc_SendOutputsRequest, completion: @escaping ApiCompletion<Walletrpc_SendOutputsResponse>) {
    completion(Result(value: sendOutputs, error: LndApiError.unknownError))
  }
  func estimateFee(_ request: Walletrpc_EstimateFeeRequest, completion: @escaping ApiCompletion<Walletrpc_EstimateFeeResponse>) {
    completion(Result(value: estimateFee, error: LndApiError.unknownError))
  }
  func pendingSweeps(_ request: Walletrpc_PendingSweepsRequest, completion: @escaping ApiCompletion<Walletrpc_PendingSweepsResponse>) {
    completion(Result(value: pendingSweeps, error: LndApiError.unknownError))
  }
  func bumpFee(_ request: Walletrpc_BumpFeeRequest, completion: @escaping ApiCompletion<Walletrpc_BumpFeeResponse>) {
    completion(Result(value: bumpFee, error: LndApiError.unknownError))
  }
}
