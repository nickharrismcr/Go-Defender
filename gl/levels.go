package gl

type Level struct {
	BaiterCount  int
	LanderCount  int
	HumanCount   int
	BomberCount  int
	PodCount     int
	SwarmerCount int
	BulletTime   float64
}

func CurrentLevel() Level {
	return levels[currentLevel]
}

var currentLevel int = 0
var levels []Level

func init() {
	levels = []Level{}
	levels = append(levels, Level{
		BaiterCount:  0,
		LanderCount:  18,
		HumanCount:   18,
		BomberCount:  0,
		PodCount:     0,
		SwarmerCount: 0,
		BulletTime:   2,
	})
	levels = append(levels, Level{
		BaiterCount:  1,
		LanderCount:  24,
		HumanCount:   24,
		BomberCount:  1,
		PodCount:     0,
		SwarmerCount: 0,
		BulletTime:   2,
	})
	levels = append(levels, Level{
		BaiterCount:  1,
		LanderCount:  24,
		HumanCount:   24,
		BomberCount:  2,
		PodCount:     1,
		SwarmerCount: 10,
		BulletTime:   2,
	})
	levels = append(levels, Level{
		BaiterCount:  1,
		LanderCount:  30,
		HumanCount:   30,
		BomberCount:  3,
		PodCount:     1,
		SwarmerCount: 10,
		BulletTime:   2,
	})
	levels = append(levels, Level{
		BaiterCount:  2,
		LanderCount:  33,
		HumanCount:   33,
		BomberCount:  3,
		PodCount:     1,
		SwarmerCount: 10,
		BulletTime:   1.5,
	})
	levels = append(levels, Level{
		BaiterCount:  2,
		LanderCount:  33,
		HumanCount:   33,
		BomberCount:  4,
		PodCount:     2,
		SwarmerCount: 10,
		BulletTime:   1.4,
	})
	levels = append(levels, Level{
		BaiterCount:  3,
		LanderCount:  36,
		HumanCount:   36,
		BomberCount:  5,
		PodCount:     3,
		SwarmerCount: 10,
		BulletTime:   1.3,
	})
	levels = append(levels, Level{
		BaiterCount:  4,
		LanderCount:  39,
		HumanCount:   39,
		BomberCount:  6,
		PodCount:     4,
		SwarmerCount: 10,
		BulletTime:   1.2,
	})
}
