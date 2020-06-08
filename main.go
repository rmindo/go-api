package main

import (
	"fmt"
	"net/http"

	"github.com/rmindo/go-api/lib"
)

/**
 * Main function
 */
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}

/**
 * Handler
 */
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, syncuser()+"\n\n")
	fmt.Fprintf(w, userinfo()+"\n\n")
	fmt.Fprintf(w, recharge()+"\n\n")
	fmt.Fprintf(w, getbetlimit()+"\n\n")
	fmt.Fprintf(w, rechargestatus()+"\n\n")

	fmt.Fprintf(w, bethistory()+"\n\n")
	fmt.Fprintf(w, updatebetlimit()+"\n\n")
	fmt.Fprintf(w, transferhistory()+"\n\n")
	fmt.Fprintf(w, h5link()+"\n\n")
	fmt.Fprintf(w, betdetail()+"\n\n")
	fmt.Fprintf(w, report()+"\n\n")
}

/**
 * Create User
 */
func syncuser() string {
	var obj = make(map[string]string)

	obj["lang"] = "zh_cn"
	obj["username"] = "xx"
	obj["channel_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("syncuser", obj, true)
}

/**
* User Info
 */
func userinfo() string {
	var obj = make(map[string]string)

	obj["username"] = "xx"
	obj["channel_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("userinfo", obj, false)
}

/**
* Recharge
 */
func recharge() string {
	var obj = make(map[string]string)

	obj["money"] = "1234"
	obj["username"] = "xx"
	obj["timestamp"] = "xx"
	obj["channel_id"] = "xx"
	obj["external_transaction_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("recharge", obj, false)
}

/**
* Get Bet Limit
 */
func getbetlimit() string {
	var obj = make(map[string]string)

	obj["room_type"] = "xx"
	obj["channel_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("getbetlimit", obj, false)
}

/**
* Recharge Status
 */
func rechargestatus() string {
	var obj = make(map[string]string)

	obj["channel_id"] = "xx"
	obj["external_transaction_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("rechargestatus", obj, false)
}

/**
* Bet history
 */
func bethistory() string {
	var obj = make(map[string]string)

	obj["end"] = "xx"
	obj["start"] = "xx"
	obj["username"] = "xx"
	obj["timestamp"] = "xx"
	obj["room_type"] = "xx"
	obj["page_size"] = "xx"
	obj["page_index"] = "xx"
	obj["bet_limits"] = "xx"
	obj["channel_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("bethistory", obj, false)
}

/**
* Update Bet Limit
 */
func updatebetlimit() string {
	var obj = make(map[string]string)

	obj["username"] = "xx"
	obj["room_type"] = "xx"
	obj["bet_limits"] = "xx"
	obj["channel_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("updatebetlimit", obj, false)
}

/**
* Transfer History
 */
func transferhistory() string {
	var obj = make(map[string]string)

	obj["end"] = "xx"
	obj["start"] = "xx"
	obj["username"] = "xx"
	obj["timestamp"] = "xx"
	obj["page_size"] = "xx"
	obj["page_index"] = "xx"
	obj["channel_id"] = "xx"
	obj["external_transaction_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("transferhistory", obj, false)
}

/**
* Get H5link
 */
func h5link() string {
	var obj = make(map[string]string)

	obj["ip"] = "xx"
	obj["username"] = "xx"
	obj["timestamp"] = "xx"
	obj["channel_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("h5link", obj, false)
}

/**
* Get Slots Bet Detail
 */
func betdetail() string {
	var obj = make(map[string]string)

	obj["round_id"] = "xx"
	obj["username"] = "xx"
	obj["timestamp"] = "xx"
	obj["channel_id"] = "xx"
	obj["sequence_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("betdetail", obj, false)
}

/**
* Get Slots Bet Detail
 */
func report() string {
	var obj = make(map[string]string)

	obj["timezone"] = "xx"
	obj["end_date"] = "xx"
	obj["timestamp"] = "xx"
	obj["start_date"] = "xx"
	obj["channel_id"] = "xx"
	obj["signature"] = lib.Sign(lib.Query(obj))

	return lib.Post("report", obj, false)
}
