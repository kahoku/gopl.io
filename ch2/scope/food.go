package scope

func init() {
	//可以访问包内的变量
	thing = append(thing, "面条")
	thing = append(thing, "米饭")
}

func GetSec() string {
	return thing[1]
}
