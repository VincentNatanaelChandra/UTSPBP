package Model

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Games struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}

type Rooms struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
	id_game   Games  `json:"id_game"`
}

type Room struct {
	ID        int    `json:"id"`
	Room_name string `json:"room_name"`
}

type Participants struct {
	ID         int     `json:"id"`
	id_room    Rooms   `json:"id_room"`
	id_account Account `json:"id_account"`
}

type RoomsResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Room `json:"data"`
}

type DetailResponse struct {
	Status int                 `json:"status"`
	Data   RoomsDetailResponse `json:"Data"`
}

type RoomDetail struct {
	ID          int           `json:"id"`
	Room_name   string        `json:"room_name"`
	Participant []Participant `json:"participants"`
}

type RoomDetailResponse struct {
	RoomDetail RoomDetail `json:"room"`
}

type RoomsDetailResponse struct {
	Status  int                `json:"status"`
	Message string             `json:"message"`
	Data    RoomDetailResponse `json:"data"`
}

type Participant struct {
	ID         int    `json:"id"`
	ID_account int    `json:"id_account"`
	Username   string `json:"username"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type SuccessResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
