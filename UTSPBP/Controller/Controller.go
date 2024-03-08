package Controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	m "UTSPBP/Model"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	id_game := r.URL.Query().Get("id_game")
	query := `SELECT id, room_name
			  FROM rooms t WHERE id_game = ?`
	roomsRow, err := db.Query(query, id_game)

	if err != nil {
		log.Println(err)
		return
	}

	var room m.Room
	var rooms []m.Room
	for roomsRow.Next() {
		if err := roomsRow.Scan(&room.ID, &room.Room_name); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		return
	}

	id_room, _ := strconv.Atoi(r.Form.Get("id_room"))
	id_account, _ := strconv.Atoi(r.Form.Get("id_account"))

	query := "SELECT COUNT(id_account) FROM participants WHERE r.id = ?"
	total_player, err := db.Query(query, id_room)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w, "Id Not Found")
		return
	}

	countplay, err := getRowCount(total_player)

	query2 := "SELECT g.max_player FROM games g JOIN rooms r ON g.id = r.id WHERE r.id = ?"
	max_player, err := db.Query(query2, id_room)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w, "Id Not Found")
		return
	}

	countmax, err := getRowCount(max_player)

	if countmax >= countplay {
		_, errQuery := db.Exec("INSERT INTO participants(id_room, id_account) values (?,?)",
			id_room,
			id_account,
		)
		if errQuery == nil {
			sendSuccessResponse(w, "Insert Success")
		} else {
			sendErrorResponse(w, "Insert Failed")
		}
	} else {
		sendErrorResponse(w, "Max Player has been reached")
	}
}

func GetDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	id_room := r.URL.Query().Get("id_room")

	query := `SELECT p.id_room, r.room_name, p.id, p.id_account, a.username
			  FROM participants p JOIN accounts a ON p.id_account = a.id
			  JOIN rooms r ON p.id_room = r.id WHERE p.rooms_id = ?`
	roomRow, err := db.Query(query, id_room)
	if err != nil {
		print(err.Error())
		sendErrorResponse(w, "No Data")
		return
	}
	var room m.RoomDetailResponse
	for roomRow.Next() {
		var participants m.Participant
		if err := roomRow.Scan(
			&room.RoomDetail.ID, &room.RoomDetail.Room_name, &participants.ID, &participants.ID_account, &participants.Username); err != nil {
			print(err.Error())
			sendErrorResponse(w, "No Data")
			return
		} else {
			room.RoomDetail.Participant = append(room.RoomDetail.Participant, participants)
		}
	}

	var response m.RoomsDetailResponse
	response.Status = 200
	response.Data = room
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func getRowCount(rows *sql.Rows) (int, error) {
	var count int
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			return 0, err
		}
		break
	}
	return count, nil
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	var response m.SuccessResponse
	response.Status = 200
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
