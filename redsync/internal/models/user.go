package models

// User defines user model
type User struct {
	ID            string   `json:"_id"`
	Index         int      `json:"index"`
	GUID          string   `json:"guid"`
	IsActive      bool     `json:"isActive"`
	Balance       string   `json:"balance"`
	Picture       string   `json:"picture"`
	Age           int      `json:"age"`
	EyeColor      string   `json:"eyeColor"`
	Name          string   `json:"name"`
	Gender        string   `json:"gender"`
	Company       string   `json:"company"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Address       string   `json:"address"`
	About         string   `json:"about"`
	Registered    string   `json:"registered"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	Tags          []string `json:"tags"`
	Friends       []Friend `json:"friends"`
	Greeting      string   `json:"greeting"`
	FavoriteFruit string   `json:"favoriteFruit"`
}

// Friend defines friend model
type Friend struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NewUser will return new dummy user data
// Generated using https://www.json-generator.com
func NewUser() *User {
	return &User{
		ID:         "5d4eb6623c489e77ab5ba636",
		Index:      0,
		GUID:       "6f865801-45cb-46c8-bb2d-9a6e8abcd09c",
		IsActive:   false,
		Balance:    "$1,728.00",
		Picture:    "http://placehold.it/32x32",
		Age:        38,
		EyeColor:   "brown",
		Name:       "Wendi Bender",
		Gender:     "female",
		Company:    "BOLAX",
		Email:      "wendibender@bolax.com",
		Phone:      "+1 (917) 493-2106",
		Address:    "450 Ross Street, Vale, Washington, 6693",
		About:      "Lorem et commodo ipsum nisi pariatur proident dolore sint cupidatat pariatur. Nostrud id enim tempor incididunt ad id non proident exercitation do exercitation exercitation adipisicing. Adipisicing fugiat commodo fugiat pariatur fugiat. Lorem occaecat quis officia eiusmod labore cillum dolore esse. Elit proident nisi ad consequat nulla commodo ullamco est ullamco cupidatat. Incididunt ex in ipsum sit.",
		Registered: "2018-06-29T07:57:03 -07:00",
		Latitude:   -87.616308,
		Longitude:  121.133756,
		Tags:       []string{"ad", "ullamco", "elit", "fugiat", "veniam", "ex", "fugiat"},
		Friends: []Friend{
			Friend{
				ID:   0,
				Name: "Richards Osborn",
			},
			Friend{
				ID:   1,
				Name: "Juliana Woods",
			},
			Friend{
				ID:   2,
				Name: "Casey Frank",
			},
		},
		Greeting:      "Hello, Wendi Bender! You have 6 unread messages.",
		FavoriteFruit: "strawberry",
	}
}
