package poker

import (
	"fmt"
	"math/rand"
	"time"
)

// CustomShuffle
// 实现一个函数，对输入的扑克牌执行洗牌，保证其是均匀分布的，也就是说列表中的每一张扑克牌出现在列表的每一个位置上的概率必须相同。
func CustomShuffle(pokers []*Poker) []*Poker {
	max := len(pokers)
	rand.Seed(time.Now().UnixNano())
	for i := range pokers {
		j := rand.Intn(max - 1)
		pokers[i], pokers[j] = pokers[j], pokers[i]
	}
	return pokers
}

const PokerNum52 = 52
const PokerNum54 = 54

// 扑克牌的花色和号码
var suit = []string{"♠", "♥", "♣", "♦"}
var rank = []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}
var joker = []string{"red joker", "black joker"}

type PokersStruct struct {
	Numbers int //牌数，54包括大小王，52不包括
	Pokers  []*Poker
	Cards   map[int]string // key对应的牌面
}

type Poker struct {
	suit  string
	rank  string
	joker string
	card  string
	key   int
}

func NewPokers() *PokersStruct {
	p := initPokers(PokerNum54)
	//p.Shuffle()
	return p
}

//初始化
func initPokers(num int) *PokersStruct {
	pokers := make([]*Poker, 54)
	cards := make(map[int]string)
	k := 0

	if num == 54 {
		//生成大小王牌面
		for i := 0; i < 2; i++ {
			temp := &Poker{
				joker: joker[i],
				card:  joker[i],
				key:   k,
			}
			pokers[k] = temp
			cards[k] = temp.card
			k++
		}
	}

	//生成其他牌面
	for i := 0; i < 13; i++ {
		for j := 0; j < 4; j++ {
			temp := &Poker{
				suit: suit[j],
				rank: rank[i],
				card: suit[j] + rank[i],
				key:  k,
			}
			pokers[k] = temp
			cards[k] = temp.card
			k++
		}
	}

	return &PokersStruct{num, pokers, cards}
}

func (p *PokersStruct) GetPokers() []*Poker {
	return p.Pokers
}

func (p *PokersStruct) GetCard(key int) string {
	return p.Pokers[key].card
}

func (p *PokersStruct) EachPokersTest() {
	for _, i1 := range p.Pokers {
		fmt.Println(i1.card)
	}
}
func EachOtherPokersTest(l []*Poker) {
	for _, i1 := range l {
		fmt.Println(i1.card)
	}
}

func (p *PokersStruct) EachCardsTest() {
	for k, i2 := range p.Cards {
		fmt.Println(k, i2)
	}
}

// Shuffle 洗牌
func (p *PokersStruct) Shuffle() {
	max := len(p.Pokers)
	rand.Seed(time.Now().UnixNano())
	for i := range p.Pokers {
		j := rand.Intn(max - 1)
		p.Pokers[i], p.Pokers[j] = p.Pokers[j], p.Pokers[i]
	}
}
