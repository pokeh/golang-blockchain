package blockchain

type TxOutput struct {
	Value  int
	PubKey string // for now we just put the account name. usually use a script
}

// references to previous outputs
type TxInput struct {
	ID  []byte //transaction that the output is inside of
	Out int    //index of the output in the transaction
	Sig string // pubkey
}

func (in *TxInput) CanUnlock(data string) bool {
	// if this is true, the account (data) owns the information in the output referenfed in the input
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlocked(data string) bool {
	// if this is true, the account (data) owns the information in the output
	return out.PubKey == data
}
