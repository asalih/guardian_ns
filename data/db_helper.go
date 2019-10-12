package data

import (
	"database/sql"
	//_ ...
	"github.com/asalih/guardian_ns/models"
	_ "github.com/lib/pq"
)

/*DNSDBHelper The database query helper*/
type DNSDBHelper struct {
}

/*GetTargetsList Reads the Target from database*/
func (h *DNSDBHelper) GetTargetsList() map[string]string {
	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	rows, qerr := conn.Query("SELECT \"Domain\" FROM public.\"Targets\" where \"State\"=1")

	if qerr != nil {
		panic(qerr)
	}

	result := make(map[string]string, 0)
	for rows.Next() {
		var target string

		ferr := rows.Scan(&target)

		if ferr != nil {
			panic(ferr)
		}

		result[target+"."] = "165.227.244.17"
	}

	return result
}
