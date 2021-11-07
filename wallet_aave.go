package hdwallet

func init() {
	coins[AAVE] = newAAVE
}

type aave struct {
	*eth
}

func newAAVE(key *Key) Wallet {
	token := newETH(key).(*eth)
	token.name = "Aave Token"
	token.symbol = "AAVE"
	token.contract = "0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9"

	return &aave{eth: token}
}
