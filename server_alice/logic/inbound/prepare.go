package inbound

import (
	"context"
	"errors"

	common "GolandProjects/apaxos-gautamsardana/api_common"
	"GolandProjects/apaxos-gautamsardana/server_alice/config"
	"GolandProjects/apaxos-gautamsardana/server_alice/utils"
)

func Prepare(ctx context.Context, req *common.Prepare, conf *config.Config) error {
	if !isValidBallot(req, conf) {
		//invalid prepare request -- do nothing (check)
		return errors.New("invalid ballot")
	}
	// valid prepare request -- process
	utils.UpdateBallot(conf, req.BallotNum.TermNumber, req.BallotNum.ServerNumber)

	// need to check for older promises here -- check whether to send old acceptNum,Val or nil with local txns

	//outbound.Promise(true, &common.Ballot{TermNumber: req.BallotNum.TermNumber, ServerNumber: req.BallotNum.ServerNumber})

	return nil
}

func isValidBallot(req *common.Prepare, conf *config.Config) bool {
	if req.BallotNum.TermNumber < conf.CurrBallot.TermNumber {
		return false
	}
	return true
}