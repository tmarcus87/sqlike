package main

import (
	"context"
	"fmt"
	"github.com/tmarcus87/sqlike"
)

func init() {
	examples["tx"] = tx
}

func tx(e sqlike.Engine) error {
	// Truncate table
	if err := e.Master(context.Background()).Truncate(bookTable).Build().Execute().Error(); err != nil {
		return fmt.Errorf("failed to truncate : %w", err)
	}

	// Check empty
	{
		records, err := e.Master(context.Background()).SelectFrom(bookTable).Build().FetchMap()
		if err != nil {
			return fmt.Errorf("failed to query : %w", err)
		}
		if len(records) != 0 {
			return fmt.Errorf("unexpected number of rows : %d", len(records))
		}
	}

	// Insert w/ tx
	tx, err := e.Master(context.Background()).Begin()
	if err != nil {
		return fmt.Errorf("failed to begin tx : %w", err)
	}
	if err :=
		tx.InsertInto(bookTable).
			Columns(bookTitleColumn, bookAuthorIdColumn).
			Values("foo", 1).
			Build().
			Execute().
			Error(); err != nil {
		return fmt.Errorf("failed to insert records : %w", err)
	}

	// Check from another session w/ uncommited
	{
		records, err := e.Master(context.Background()).SelectFrom(bookTable).Build().FetchMap()
		if err != nil {
			return fmt.Errorf("failed to query : %w", err)
		}
		if len(records) != 0 {
			return fmt.Errorf("unexpected number of rows : %d", len(records))
		}
	}

	// Commit
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit tx : %w", err)
	}

	// Check from another session
	{
		records, err := e.Master(context.Background()).SelectFrom(bookTable).Build().FetchMap()
		if err != nil {
			return fmt.Errorf("failed to query : %w", err)
		}
		if len(records) != 1 {
			return fmt.Errorf("unexpected number of rows : %d", len(records))
		}
	}

	return nil

}
