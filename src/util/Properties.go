package util

import "fmt"

func main1() {

	data := []string{}
	data = testArray(data)

	fmt.Println(data)
}

func testArray(datas []string) []string {
	datas = append(datas, "ok")
	return datas
}

func propertiesTest222() {
	//	var propertiesExample = []byte(`
	//isc.log.hosts=root:dell@123:10.30.30.33:22\
	//root:dell@123:10.30.30.35:22
	//isc.name=123
	//`)
	// 读取内容
	//r := bytes.NewReader(propertiesExample)
	//pro := properties.NewProperties()
	//err := pro.Load(r)
	//if nil != err {
	//	fmt.Println("异常")
	//	return
	//}
	//
	//// 输出内容
	//str := ""
	//bf := bytes.NewBufferString(str)
	//err = pro.Store(bf)
	//if nil != err {
	//	fmt.Println("异常")
	//	return
	//}
	//fmt.Println(bf.String())
	//

	//// [p_ppu p_batters.batter.type p_id[0] p_id[1] p_type p_name]
	//fmt.Println(pro.PropertyNames())
	//// [Cake] true
	//fmt.Println(pro.PropertySlice("p_name"))
	//// [donut] true
	//fmt.Println(pro.PropertySlice("p_type"))
	//// [] false
	//fmt.Println(pro.PropertySlice("p_id"))
}
