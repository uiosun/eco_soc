package world

// Work 工种
type Work int32

// GoodsType 物品类
type GoodsType int32

func (gt GoodsType) String() string {
	return GoodsTypes[gt]
}

var (
	GoodsTypeWheat = GoodsType(1)
	GoodsTypeWood  = GoodsType(2)
	GoodsTypes     = map[GoodsType]string{
		GoodsTypeWheat: "小麦",
		GoodsTypeWood:  "原木",
	}

	Idle        = Work(0)
	CoinForger  = Work(1) // 铸币匠
	WheatFarmer = Work(2) // 农夫（小麦）
	Logger      = Work(3) // 伐木工
)

type OrderSell struct {
	HumanId   int
	GoodsType GoodsType
	WantSell  int // 出售数量
}

type World struct {
	Market     map[GoodsType][]*OrderSell
	Humans     map[int]*Human
	HumanMaxId int
}

func (w *World) Init() {
	w.Market = map[GoodsType][]*OrderSell{
		GoodsTypeWheat: {},
		GoodsTypeWood:  {},
	}
	w.Humans = map[int]*Human{
		w.HumanMaxId: {
			Id:   w.HumanMaxId,
			Work: WheatFarmer,
			Age:  18,
			Cash: 10,
			Bag:  map[GoodsType]int{},
		},
	}
	w.HumanMaxId++

	for _, human := range w.Humans {
		switch human.Work {
		case WheatFarmer:
			human.State |= 1
		}
	}
}
