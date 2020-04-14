package scope
import ( 
	"fmt"
 )

func init() {
	thing = append(thing, "apple")
	thing = append(thing, "orange")
}

func GetFirst() string {
	fmt.Println("scope first ele is ", thing[0])
	return thing[0]
}
