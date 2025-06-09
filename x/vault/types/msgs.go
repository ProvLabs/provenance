package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AllRequestMsgs defines all the Msg*Request messages.
var AllRequestMsgs = []sdk.Msg{
	(*MsgCreateVaultRequest)(nil),
	(*MsgDepositRequest)(nil),
	(*MsgWithdrawRequest)(nil),
	(*MsgRedeemRequest)(nil),
}

// ValidateBasic returns a not implemented error for MsgCreateVaultRequest.
func (m MsgCreateVaultRequest) ValidateBasic() error {
	return errors.New("ValidateBasic not implemented for MsgCreateVaultRequest")
}

// ValidateBasic returns a not implemented error for MsgDepositRequest.
func (m MsgDepositRequest) ValidateBasic() error {
	return errors.New("ValidateBasic not implemented for MsgDepositRequest")
}

// ValidateBasic returns a not implemented error for MsgWithdrawRequest.
func (m MsgWithdrawRequest) ValidateBasic() error {
	return errors.New("ValidateBasic not implemented for MsgWithdrawRequest")
}

// ValidateBasic returns a not implemented error for MsgRedeemRequest.
func (m MsgRedeemRequest) ValidateBasic() error {
	return errors.New("ValidateBasic not implemented for MsgRedeemRequest")
}
