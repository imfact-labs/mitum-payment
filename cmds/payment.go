package cmds

type PaymentCommand struct {
	Deposit              DepositCommand           `cmd:"" name:"deposit" help:"deposit"`
	Withdraw             WithdrawCommand          `cmd:"" name:"withdraw" help:"withdraw"`
	Transfer             TransferCommand          `cmd:"" name:"transfer" help:"transfer"`
	UpdateAccountSetting UpdateAccountInfoCommand `cmd:"" name:"update-account-setting" help:"update account setting"`
	RegisterModel        RegisterModelCommand     `cmd:"" name:"register-model" help:"register payment model"`
}
