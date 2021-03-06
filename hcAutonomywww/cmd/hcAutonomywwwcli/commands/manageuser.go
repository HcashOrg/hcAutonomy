package commands

import (
	"fmt"
	"strconv"

	"github.com/HcashOrg/hcAutonomy/hcAutonomywww/api/v1"
)

type ManageUserCmd struct {
	Args struct {
		UserID string `positional-arg-name:"userid" description:"User ID"`
		Action string `positional-arg-name:"action" description:"(Admin) edit user action"`
		Reason string `positional-arg-name:"reason" description:"Reason for editing the user"`
	} `positional-args:"true" required:"true"`
}

func (cmd *ManageUserCmd) Execute(args []string) error {
	ManageActions := map[string]v1.UserManageActionT{
		"expirenewuser":       1,
		"expireupdatekey":     2,
		"expireresetpassword": 3,
		"clearpaywall":        4,
		"unlock":              5,
		"deactivate":          6,
		"reactivate":          7,
	}

	// Parse edit user action.  This can be either the numeric
	// action code or the human readable equivalent.
	var action v1.UserManageActionT
	a, err := strconv.ParseUint(cmd.Args.Action, 10, 32)
	if err == nil {
		// Numeric action code found
		action = v1.UserManageActionT(a)
	} else if a, ok := ManageActions[cmd.Args.Action]; ok {
		// Human readable action code found
		action = a
	} else {
		return fmt.Errorf("Invalid useredit action.  Valid actions are:\n  " +
			"expirenewuser         expires new user verification\n  " +
			"expireupdatekey       expires update user key verification\n  " +
			"expireresetpassword   expires reset password verification\n  " +
			"clearpaywall          clears user registration paywall\n  " +
			"unlock                unlocks user account from failed logins\n  " +
			"deactivate            deactivates user account\n  " +
			"reactivate            reactivates user account")
	}

	// Setup request
	mu := &v1.ManageUser{
		UserID: cmd.Args.UserID,
		Action: action,
		Reason: cmd.Args.Reason,
	}

	// Print request details
	err = Print(mu, cfg.Verbose, cfg.RawJSON)
	if err != nil {
		return err
	}

	// Send request
	mur, err := c.ManageUser(mu)
	if err != nil {
		return err
	}

	// Print response details
	return Print(mur, cfg.Verbose, cfg.RawJSON)
}
