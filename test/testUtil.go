package testUtil

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
)

// CreateContainer dockertestを用いてテスト用のDBを立ち上げる
func CreateContainer() (*dockertest.Resource, *dockertest.Pool) {

	// Dockerとの接続
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
	// テーブル定義などはここで流し込むのが良さそう
	runOptions := &dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_DB=npb-analysis",
			"POSTGRES_USER=npb-analysis",
			"POSTGRES_PASSWORD=postgres",
		},
		Mounts: []string{
			"/yourdir/npb-analysis/docker/initdb:/docker-entrypoint-initdb.d", // コンテナ起動時に実行したいSQL
		},
	}

	// コンテナを起動
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

// ConnectDB テスト用のDBとの接続を行う
func ConnectDB(resource *dockertest.Resource, pool *dockertest.Pool) *sql.DB {
	// DB(コンテナ)との接続
	var db *sql.DB
	if err := pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("host=localhost port=%s password=postgres user=npb-analysis dbname=npb-analysis sslmode=disable", resource.GetPort("5432/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	return db
}

// CloseContainer テスト用のDBを切断する
func CloseContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	// コンテナの終了
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}
