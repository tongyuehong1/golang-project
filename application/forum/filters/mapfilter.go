package filters

var MapFilter map[string]interface{}

func init() {
	MapFilter = make(map[string]interface{})
	MapFilter["/user/login"] = struct{}{}
	MapFilter["/user/create"] = struct{}{}

	MapFilter["/admin/create"] = struct{}{}
	MapFilter["/admin/login"] = struct{}{}
}
