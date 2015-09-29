package goddos

import (
	"net/http"
	"strings"

	"github.com/pilu/traffic"
)

var (
	uriFragments map[string]bool
)

func init() {
	uriFragments = make(map[string]bool)
	uriFragments["manager/html"] = true
	uriFragments["POST_ip_port.php"] = true
	uriFragments["tmUnblock.cgi"] = true
	uriFragments["HNAP1"] = true
	uriFragments["phpMyAdmin/scripts/setup.php"] = true
	uriFragments["pma/scripts/setup.php"] = true
	uriFragments["myadmin/scripts/setup.php"] = true
	uriFragments["MyAdmin/scripts/setup.php"] = true
	uriFragments["vyvy/vyv/vy.php"] = true
	uriFragments["cgi-sys/php5"] = true
	uriFragments["cgi-bin/test-cgi"] = true
	uriFragments["cgi-bin/printenv"] = true
	uriFragments["cgi-bin/test.cgi"] = true
	uriFragments["cgi-bin/test.pl"] = true
	uriFragments["cgi-bin/test.sh"] = true
	uriFragments["cgi-bin/teste.pl"] = true
	uriFragments["cgi-bin/teste.cgi"] = true
	uriFragments["cgi-bin/teste.sh"] = true
	uriFragments["cgi-bin/print-env"] = true
	uriFragments["cgi-bin/print.pl"] = true
	uriFragments["cgi-bin/print.cgi"] = true
	uriFragments["cgi-bin/printenv.sh"] = true
	uriFragments["dpdp/dpd/dp.php"] = true
	uriFragments["upup/upu/up.php"] = true
	uriFragments["admin/fckeditor/editor/filemanager/browser/default/connectors/test.html"] = true
	uriFragments["web-console/ServerInfo.jsp"] = true
	uriFragments["vtigercrm"] = true
	uriFragments["operator/basic.shtml"] = true
	uriFragments["secure/ltx_conf.htm"] = true
	uriFragments["syslog.htm"] = true
}

func main() {

}

func CheckForDdosString(w traffic.ResponseWriter, r *traffic.Request) {
	for k, _ := range uriFragments {
		if strings.Contains(k, r.RequestURI) {
			w.WriteHeader(http.StatusGatewayTimeout)
			w.WriteText("connection timed out")
		}
	}
}
