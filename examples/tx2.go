package main

import (
	"context"
	"fmt"
	"github.com/tmarcus87/sqlike"
)

func init() {
	examples["tx2"] = tx2
}

func tx2(e sqlike.Engine) error {
	// Truncate table
	if err := e.NewSession(sqlike.MarkAsExecuted(context.Background())).Truncate(bookTable).Build().Execute().Error(); err != nil {
		return fmt.Errorf("failed to truncate : %w", err)
	}

	// Check empty
	{
		records, err := e.NewSession(sqlike.MarkAsExecuted(context.Background())).SelectFrom(bookTable).Build().FetchMap()
		if err != nil {
			return fmt.Errorf("failed to query : %w", err)
		}
		if len(records) != 0 {
			return fmt.Errorf("unexpected number of rows : %d", len(records))
		}
	}

	// Begin tx
	ctx, err := e.BeginTx(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin : %w", err)
	}

	sess, err := e.GetTxSession(ctx)
	if err != nil {
		return err
	}

	// Insert w/ tx
	if err :=
		sess.InsertInto(bookTable).
			Columns(bookTitleColumn, bookAuthorIdColumn).
			Values("foo", 1).
			Build().
			Execute().
			Error(); err != nil {
		return fmt.Errorf("failed to insert records : %w", err)
	}

	// Check from another session w/ uncommited
	{
		records, err := e.NewSession(sqlike.MarkAsExecuted(context.Background())).SelectFrom(bookTable).Build().FetchMap()
		if err != nil {
			return fmt.Errorf("failed to query : %w", err)
		}
		if len(records) != 0 {
			return fmt.Errorf("unexpected number of rows : %d", len(records))
		}
	}

	// Commit
	if _, err := e.CommitTx(ctx); err != nil {
		return fmt.Errorf("failed to commit tx : %w", err)
	}

	// Check from another session
	{
		records, err := e.NewSession(sqlike.MarkAsExecuted(context.Background())).SelectFrom(bookTable).Build().FetchMap()
		if err != nil {
			return fmt.Errorf("failed to query : %w", err)
		}
		if len(records) != 1 {
			return fmt.Errorf("unexpected number of rows : %d", len(records))
		}
	}

	return nil

}
