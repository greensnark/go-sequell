package action

import (
	"fmt"
	"os"

	"github.com/greensnark/go-sequell/crawl/data"
	cdb "github.com/greensnark/go-sequell/crawl/db"
	"github.com/greensnark/go-sequell/pg"
	"github.com/greensnark/go-sequell/schema"
)

func CrawlSchema() *schema.Schema {
	schema, err := cdb.LoadSchema(data.CrawlData())
	if err != nil {
		panic(err)
	}
	return schema.Schema()
}

func PrintSchema(skipIndexes, dropIndexes, createIndexes bool) {
	s := CrawlSchema()
	sel := schema.SelTablesIndexes
	if skipIndexes {
		sel = schema.SelTables
	}
	if dropIndexes {
		sel = schema.SelDropIndexes
	}
	if createIndexes {
		sel = schema.SelIndexes
	}
	s.Sort().Write(sel, os.Stdout)
}

func DumpSchema(dbspec pg.ConnSpec) error {
	db, err := dbspec.Open()
	if err != nil {
		return err
	}
	s, err := db.IntrospectSchema()
	if err != nil {
		return err
	}
	s.Sort().Write(schema.SelTablesIndexes, os.Stdout)
	return nil
}

func CheckDBSchema(dbspec pg.ConnSpec, applyDelta bool) error {
	db, err := dbspec.Open()
	if err != nil {
		return err
	}
	actualSchema, err := db.IntrospectSchema()
	if err != nil {
		return err
	}
	wantedSchema := CrawlSchema()
	diff := wantedSchema.DiffSchema(actualSchema)
	if len(diff.Tables) == 0 {
		fmt.Fprintf(os.Stderr, "Schema is up-to-date.\n")
		return nil
	}

	fmt.Fprintf(os.Stderr, "Schema delta:\n")
	diff.PrintDelta(os.Stderr)
	if applyDelta {
		return nil
	}
	return nil
}

func CreateDB(admin, db pg.ConnSpec) error {
	return nil
}