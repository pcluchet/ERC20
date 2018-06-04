package main

type SimpleAsset struct {
}

type AllowanceCouple struct {
	Spender string
	Amount  uint64
}

type AllowanceCouples []AllowanceCouple

type UserInfos struct {
	Amount     uint64
	Allowances AllowanceCouples
}
