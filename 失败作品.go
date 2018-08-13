package main
import (
	"fmt"
)
type Vertex struct{
	a int
	b int
	c int
	d int
	e int
}
func main(){
	v := Vertex{2,5,3,6,1}
	for i:=1;v.a>v.b&&v.b>v.c&&v.c>v.d&&v.d>v.e;i++ {
		if v.a<v.b {
		}
		if v.b<v.c {
			v.b,v.c = v.c,v.b
		}
		if v.c<v.d {
			v.c,v.d=v.d,v.c
		}
		if v.d<v.e {
			v.d,v.e=v.e,v.d
		}
		fmt.Println(v)
	}
	
	
}
