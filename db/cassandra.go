


package db

import (
	"log"
	pb "github.com/akililab/transaction/proto"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"fmt"
)
// CassandraSession exported
var CassandraSession *gocql.Session

// GetTransactionDB : Fetch from database and return values
func GetTransactionDB(req *pb.TransactionRequest) (*pb.TransactionReply,error) {

	// Create a query which uses the built query and populates it with the
	// values in the new item
	stmt, names := qb.Select("akililab.transactions").Where(qb.Eq("account_id")).ToCql()
	query := gocqlx.Query(CassandraSession.Query(stmt), names).BindMap(qb.M{
        "account_id": req.AccountId,
	})
	var res  *pb.TransactionReply
	// Run that query and release it when done
	err := query.GetRelease(&res)

	return res, err
}







// Initialize : Initialize database
func init() {

	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "system"
	CassandraSession, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Here is called")

	// Close session when done initializing it
	defer CassandraSession.Close()
}