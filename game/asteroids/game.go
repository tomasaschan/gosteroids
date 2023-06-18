package asteroids

type asteroids struct{}

func (a *asteroids) Name() string { return "Asteroids" }

func (a *asteroids) Slug() any {
	return &slug{}
}

func NewGame() *asteroids {
	return &asteroids{}
}
