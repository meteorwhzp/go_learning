/*
go语言实现设计模型
 */
package pattern

import "fmt"

/*
工厂方法模式
用于弥补简单工厂的不足之处：
当新增一种产品时，需要修改工厂逻辑
简单工厂没有继承的结构
一旦工厂瘫痪，整个系统都瘫痪
*/

type Factory interface {
	Create() Pen
}

type Pen interface {
	Write() string
}

type PencilFactory struct {
}

type Pencil struct {
}

type BrushPenFactory struct {
}

type BrushPen struct {
}

func (factory *PencilFactory) Create() Pen {
	return &Pencil{}
}

func (factory *BrushPenFactory) Create() Pen {
	return &BrushPen{}
}

func (pen *Pencil) Write() string {
	return "pencil"
}

func (pen *BrushPen) Write() string {
	return "brushPen"
}

func GetPenFactory(name string) Factory {
	switch name {
	case "pencil":
		return &PencilFactory{}
	case "brushPen":
		return &BrushPenFactory{}
	}
	fmt.Println("unknown name")
	return nil
}