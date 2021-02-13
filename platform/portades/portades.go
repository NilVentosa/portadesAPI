package portades

import "database/sql"

type Data struct {
	DB *sql.DB
}

func (data *Data) GetRandomPortada() (Portada, bool) {
	row := data.DB.QueryRow("SELECT * from portades ORDER BY RANDOM() LIMIT 1")
	var id int
	var intro string
	var newspaper string
	var headline string
	var result bool
	var video string
	var episode string

	err := row.Scan(&id, &intro, &newspaper, &headline, &result, &video, &episode)
	portada := Portada{
		Id:        id,
		Intro:     intro,
		Newspaper: newspaper,
		Headline:  headline,
		Result:    result,
		Video:     video,
		Episode:   episode,
	}
	if err != nil {
		return Portada{}, false
	}
	return portada, true
}

func (data *Data) GetPortada(queryId int) (Portada, bool) {
	row := data.DB.QueryRow("SELECT * from portades where id = ?", queryId)
	var id int
	var intro string
	var newspaper string
	var headline string
	var result bool
	var video string
	var episode string

	err := row.Scan(&id, &intro, &newspaper, &headline, &result, &video, &episode)
	portada := Portada{
		Id:        id,
		Intro:     intro,
		Newspaper: newspaper,
		Headline:  headline,
		Result:    result,
		Video:     video,
		Episode:   episode,
	}
	if err != nil {
		return Portada{}, false
	}
	return portada, true
}

func NewData(db *sql.DB) *Data {
	return &Data{
		DB: db,
	}
}
