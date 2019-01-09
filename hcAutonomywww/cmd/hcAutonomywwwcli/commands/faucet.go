package commands

import (
	"fmt"

	"github.com/HcashOrg/hcAutonomy/util"
)

type FaucetCmd struct {
	Args struct {
		Address       string `positional-arg-name:"address" required:"true" description:"Address to send HC to"`
		Amount        uint64 `positional-arg-name:"amount" required:"true" description:"Amount to send (in atoms)"`
		OverrideToken string `positional-arg-name:"overridetoken" description:"Override token for testnet faucet"`
	} `positional-args:"true"`
}

func (cmd *FaucetCmd) Execute(args []string) error {
	address := cmd.Args.Address
	atoms := cmd.Args.Amount
	hc := float64(atoms) / 1e8

	if address == "" && atoms == 0 {
		return fmt.Errorf("Invalid arguments. Unable to pay %v HC to %v",
			hc, address)
	}

	txID, err := util.PayWithTestnetFaucet(cfg.FaucetHost, address, atoms,
		cmd.Args.OverrideToken)
	if err != nil {
		return err
	}

	if cfg.RawJSON {
		fmt.Printf(`{"txid":"%v"}`, txID)
		fmt.Printf("\n")
	} else {
		fmt.Printf("Paid %v HC to %v with txID %v\n",
			hc, address, txID)
	}

	return nil
}
