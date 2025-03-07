package migrations

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func Migrations(ctx context.Context, pg *pgx.Conn) {
	var exist bool
	err := pg.QueryRow(ctx, "SELECT EXISTS (SELECT FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'todo')").Scan(&exist)
	if exist {
		return
	}

	_, err = pg.Exec(ctx, "CREATE TABLE todo (id SERIAL  PRIMARY KEY, title TEXT NOT NULL, description TEXT ,status TEXT CHECK(status IN ('new', 'in_progress', 'done')) DEFAULT 'new' ,created_at TIMESTAMP DEFAULT now(), updated_at TIMESTAMP DEFAULT now())")
	if err != nil {
		panic(err)
	}
}
