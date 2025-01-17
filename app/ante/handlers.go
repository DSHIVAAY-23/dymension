package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcante "github.com/cosmos/ibc-go/v7/modules/core/ante"
	"github.com/dymensionxyz/dymension/v3/x/common/types"
	ethante "github.com/evmos/ethermint/app/ante"
	txfeesante "github.com/osmosis-labs/osmosis/v15/x/txfees/ante"

	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	lightclientante "github.com/dymensionxyz/dymension/v3/x/lightclient/ante"
)

func newEthAnteHandler(options HandlerOptions) sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ethante.NewEthSetUpContextDecorator(options.EvmKeeper),

		// TODO: need to allow universal fees for Eth as well
		ethante.NewEthMempoolFeeDecorator(options.EvmKeeper),                           // Check eth effective gas price against minimal-gas-prices
		ethante.NewEthMinGasPriceDecorator(options.FeeMarketKeeper, options.EvmKeeper), // Check eth effective gas price against the global MinGasPrice

		ethante.NewEthValidateBasicDecorator(options.EvmKeeper),
		ethante.NewEthSigVerificationDecorator(options.EvmKeeper),
		ethante.NewEthAccountVerificationDecorator(options.AccountKeeper, options.EvmKeeper),
		ethante.NewCanTransferDecorator(options.EvmKeeper),
		ethante.NewEthGasConsumeDecorator(options.EvmKeeper, options.MaxTxGasWanted),
		ethante.NewEthIncrementSenderSequenceDecorator(options.AccountKeeper), // innermost AnteDecorator.
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
		ethante.NewEthEmitEventDecorator(options.EvmKeeper), // emit eth tx hash and index at the very last ante handler.
	)
}

// newLegacyCosmosAnteHandlerEip712 creates an AnteHandler to process legacy EIP-712
// transactions, as defined by the presence of an ExtensionOptionsWeb3Tx extension.
func newLegacyCosmosAnteHandlerEip712(options HandlerOptions) sdk.AnteHandler {
	mempoolFeeDecorator := txfeesante.NewMempoolFeeDecorator(*options.TxFeesKeeper, options.FeeMarketKeeper)
	deductFeeDecorator := txfeesante.NewDeductFeeDecorator(*options.TxFeesKeeper, options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper)

	return sdk.ChainAnteDecorators(
		/*
			See https://jumpcrypto.com/writing/bypassing-ethermint-ante-handlers/
			for an explanation of these message blocking decorators
		*/
		// reject MsgEthereumTxs and disable the Msg types that cannot be included on an authz.MsgExec msgs field
		NewRejectMessagesDecorator().WithPredicate(
			BlockTypeUrls(
				1,
				// Only blanket rejects depth greater than zero because we have our own custom logic for depth 0
				// Note that there is never a genuine reason to pass both ibc update client and misbehaviour submission through gov or auth,
				// it's always done by relayers directly.
				sdk.MsgTypeURL(&ibcclienttypes.MsgUpdateClient{}),
				sdk.MsgTypeURL(&ibcclienttypes.MsgSubmitMisbehaviour{}))).
			WithPredicate(BlockTypeUrls(
				0,
				sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}))),

		ante.NewSetUpContextDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),

		// Use Mempool Fee TransferEnabledDecorator from our txfees module instead of default one from auth
		mempoolFeeDecorator,
		deductFeeDecorator,

		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		// SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewSetPubKeyDecorator(options.AccountKeeper),
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, ethante.DefaultSigVerificationGasConsumer),
		// Note: signature verification uses EIP instead of the cosmos signature validator
		NewLegacyEip712SigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		lightclientante.NewIBCMessagesDecorator(*options.LightClientKeeper, options.IBCKeeper.ClientKeeper, options.IBCKeeper.ChannelKeeper, options.RollappKeeper),
		types.NewIBCProofHeightDecorator(),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}

func newCosmosAnteHandler(options HandlerOptions) sdk.AnteHandler {
	mempoolFeeDecorator := txfeesante.NewMempoolFeeDecorator(*options.TxFeesKeeper, options.FeeMarketKeeper)
	deductFeeDecorator := txfeesante.NewDeductFeeDecorator(*options.TxFeesKeeper, options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper)

	return sdk.ChainAnteDecorators(
		// reject MsgEthereumTxs and disable the Msg types that cannot be included on an authz.MsgExec msgs field
		NewRejectMessagesDecorator().WithPredicate(
			BlockTypeUrls(
				1,
				// Only blanket rejects depth greater than zero because we have our own custom logic for depth 0
				// Note that there is never a genuine reason to pass both ibc update client and misbehaviour submission through gov or auth,
				// it's always done by relayers directly.
				sdk.MsgTypeURL(&ibcclienttypes.MsgUpdateClient{}),
				sdk.MsgTypeURL(&ibcclienttypes.MsgSubmitMisbehaviour{}))).
			WithPredicate(BlockTypeUrls(
				0,
				sdk.MsgTypeURL(&evmtypes.MsgEthereumTx{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreateVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePeriodicVestingAccount{}),
				sdk.MsgTypeURL(&vestingtypes.MsgCreatePermanentLockedAccount{}))),
		ante.NewSetUpContextDecorator(),
		ante.NewExtensionOptionsDecorator(options.ExtensionOptionChecker),
		// Use Mempool Fee TransferEnabledDecorator from our txfees module instead of default one from auth
		mempoolFeeDecorator,
		deductFeeDecorator,
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, ethante.DefaultSigVerificationGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
		types.NewIBCProofHeightDecorator(),
		lightclientante.NewIBCMessagesDecorator(*options.LightClientKeeper, options.IBCKeeper.ClientKeeper, options.IBCKeeper.ChannelKeeper, options.RollappKeeper),
		ibcante.NewRedundantRelayDecorator(options.IBCKeeper),
		ethante.NewGasWantedDecorator(options.EvmKeeper, options.FeeMarketKeeper),
	)
}
