package data

import (
	"database/sql"
	"net"
	"strings"
	"time"

	"github.com/google/uuid"

	//_ ...
	"github.com/asalih/guardian_ns/models"
	_ "github.com/lib/pq"
)

/*DNSDBHelper The database query helper*/
type DNSDBHelper struct {
}

/*GetTargetsList Reads the Target from database*/
func (h *DNSDBHelper) GetTargetsList() map[string]net.IP {
	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	rows, qerr := conn.Query("SELECT \"Domain\" FROM public.\"Targets\" where \"State\"=1")

	if qerr != nil {
		panic(qerr)
	}

	result := make(map[string]net.IP, 0)
	for rows.Next() {
		var target string

		ferr := rows.Scan(&target)

		if ferr != nil {
			panic(ferr)
		}

		result[target+"."] = net.ParseIP(models.Configuration.GuardianIPAddress)

		if strings.HasPrefix(target, "www.") {
			result[strings.ReplaceAll(target, "www.", "")+"."] = net.ParseIP(models.Configuration.GuardianIPAddress)
		} else {
			result["www."+target+"."] = net.ParseIP(models.Configuration.GuardianIPAddress)
		}
	}

	return result
}

//LogThrottle ...
func (h *DNSDBHelper) LogThrottleRequest(ipAddress string) {
	conn, err := sql.Open("postgres", models.Configuration.ConnectionString)
	defer conn.Close()

	if err != nil {
		panic(err)
	}

	sqlStatement := `
INSERT INTO "ThrottleLogs" ("Id", "CreatedAt", "IPAddress", "ThrottleType")
VALUES ($1, $2, $3, $4)`

	_, err = conn.Exec(sqlStatement,
		uuid.New(),
		time.Now(),
		ipAddress,
		0)

	if err != nil {
		panic(err)
	}
}
