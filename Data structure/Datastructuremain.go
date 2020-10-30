package main
import (
"encoding/json"
"fmt"
"os"
)
type person struct {
Name string
Age int
}
func main() {
p1 := person{"karthik", 21}
json.NewEncoder(os.Stdout).Encode(p1) 
fmt.Println(p1.Name, "--", p1.Age)
}
