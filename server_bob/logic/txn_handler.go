package logic

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	common "GolandProjects/apaxos-gautamsardana/api_common"
	"GolandProjects/apaxos-gautamsardana/server_bob/config"
)

func EnqueueTxn(ctx context.Context, req *common.TxnRequest, conf *config.Config) error {
	fmt.Printf("Server %d: received enqueue txn request:%v\n", conf.ServerNumber, req)
	conf.QueueMutex.Lock()
	defer conf.QueueMutex.Unlock()

	conf.TxnQueue = append(conf.TxnQueue, req)
	return nil
}

func ProcessTxn(ctx context.Context, req *common.TxnRequest, conf *config.Config) error {
	fmt.Printf("Server %d: received process txn request:%v\n", conf.ServerNumber, req)

	txn, err := ValidateTxnInDB(conf, req)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if txn != nil {
		err = fmt.Errorf("Server %d: duplicate txn found in db %v\n", conf.ServerNumber, req)
		return err
	}

	txn = ValidateTxnInLogs(conf, req)
	if txn != nil {
		err = fmt.Errorf("Server %d: duplicate txn found in logs %v\n", conf.ServerNumber, req)
		return err
	}

	conf.StartTime = time.Now()
	balance := conf.Balance

	if balance >= req.Amount {
		err = ExecuteTxn(ctx, req, conf)
		if err != nil {
			return err
		}
	} else {
		fmt.Println("this is where the magic happens!")
		SendPrepare(context.Background(), conf)
	}
	fmt.Printf("-------- %s\n", time.Since(conf.StartTime))
	return nil
}

func ExecuteTxn(ctx context.Context, req *common.TxnRequest, conf *config.Config) error {
	fmt.Printf("Server %d: executing txn: %v\n", conf.ServerNumber, req)
	conf.LogStore.AddTransactionLog(req)
	conf.Balance -= req.Amount
	fmt.Println(conf.LogStore.Logs, conf.Balance)
	return nil
}