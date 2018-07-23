package main


import "fmt"

type doc struct {
	id int
	content string
	score int
}

type docList []*doc
type HotDocList docList


func (dl docList) Less(i, j int) bool {
	return dl[i].score < dl[j].score
}

func (dl docList) Len() int {
	return len(dl)
}

func test1() {

    // hd := HotDocList{}
    // fmt.Println(hd.Len())   // no method of Len

}

type Base struct{}

func (Base) Magic() {
    fmt.Println("base magic")
}

func (self Base) MoreMagic() {
    self.Magic()
    self.Magic()
}

type Voodoo struct {
    Base
}

func (Voodoo) Magic() {
    fmt.Println("voodoo magic")
}


func test2() {

	v := new(Voodoo)
	v.Magic()  // voodoo magic
	v.MoreMagic()

}

func main() {
	test2()
}
