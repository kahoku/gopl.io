package scope

func init() {
	thing = append(thing, "面条")
	thing = append(thing, "米饭")

}

func GetSec() string {
	return thing[1]
}
