package main

import "net/http"

func main() {
	http.HandleFunc("/login/getrsakey", loginGetrsakey)
	_ = http.ListenAndServe(":9090", nil)
}

func loginGetrsakey(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")

	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// Steam always returns success
	w.WriteHeader(http.StatusOK)

	if username == "" {
		_, _ = w.Write(getLoginGetrsakeyFailure())
	} else {
		// mangosteam username
		_, _ = w.Write(getLoginGetrsakeySuccess())
	}

}

func getLoginGetrsakeySuccess() []byte {
	return []byte("{\"success\":true,\"publickey_mod\":\"F0BE4CA9BE9265A96E7079B7DAA042BF7CD6B8CB9EE7A2AA6BBF028FEC9EB7510D29E0F80382700D9EDF9784BC3FA7D53578C47C64E126C73FB5781732F05F15D25A7270EC152F2039E7F89144A966F0CB9F0C340210F81F5C2A6D9A00AFC6E4B563AB81C6EE135815DD1843AAA98F4C91CF9C517944D04393ADA743FA5DE276390069D3745A6F0755C09276107644CA0EE2AAD9F5D32A962CC4D74D51979D33585BD6D95A87C2D8B0BF5726E476896F45B91478167F9419A3520F6A61CD6FD2EAEFA7F186BAF31359BABCF55E2FF4063AFE3C0E35FB11B48C3A052AA8B7893932E03C7B4ECC75A0F9E44EC68339F7A73153FE2FB4A3F0F5B7E18EC3686B1147\",\"publickey_exp\":\"010001\",\"timestamp\":\"580281350000\",\"steamid\":\"76561198063143983\",\"token_gid\":\"8213f42ccbc1395\"}")
}

func getLoginGetrsakeyFailure() []byte {
	return []byte("{\"success\":false}")
}
