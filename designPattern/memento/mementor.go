package memento

import "fmt"

//备忘录模式用于保存程序内部状态到外部，又不希望暴露内部状态的情形。
//
//程序内部状态使用窄接口船体给外部进行存储，从而不暴露程序实现细节。
//
//备忘录模式同时可以离线保存内部状态，如保存到数据库，文件等。

type Mementor interface {

}

type Game struct {
	hp, mp int
}

type GameMemento struct {
	hp, mp int
}

func (g *Game) Play(mpDelta, hpDelta int){
	g.mp += mpDelta
	g.hp += hpDelta
}

func (g *Game) Save()Mementor{
	return &GameMemento{g.hp,g.mp}
}

func (g *Game) Load(m Mementor){
	gm := m.(*GameMemento)
	g.mp = gm.mp
	g.hp = gm.hp
}

func (g *Game) Status() {
	fmt.Printf("Current HP:%d, MP:%d\n", g.hp, g.mp)
}