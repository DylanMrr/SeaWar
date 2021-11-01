package input

var arr []string = []string{"a1", "a3", "a5", "a7", "a9", "0", "c1", "0", "c3", "0", "c5", "0", "c7", "0", "f1", "0"}
var i int

func Get() string {
	temp := arr[i]
	i++
	return temp
}
