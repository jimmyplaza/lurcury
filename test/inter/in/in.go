package main
import ("fmt")
type Coder interface{
    Speak() string
}

type PHP struct{}
type Java struct{}
type C struct{}
type Go struct{}
type Python struct{}

func (h PHP) Speak() string{
    return "PHP is the best"//?
}
func (j Java)Speak() string{
    return "Java is cool"
}

func (c C)Speak() string{
    return "C is old"
}
func (g Go)Speak() string{
    return "GO GO GO"
}
func (p Python)Speak() string{
    return "faster"
}
func main(){
    coder:=[]Coder{PHP{},Java{},C{},Go{},Python{}}
    for _,man := range coder{
        fmt.Println(man.Speak())
}
}

