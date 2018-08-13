package patient

import (
	"fmt"
	"io"
	"parkinsons/database"
	"encoding/json"
)

// Note the fields in the struct must be capitalized otherwise they won't be exported
// and json.Marshal won't be able to access those fields.

type Patient struct {
	Id        int			`json:"id"`
	FirstName string	`json:"first_name"`
	LastName  string	`json:"last_name"`
	Age       int			`json:"age"`
	User_Id   int			`json:"user_id"`
}

var firstname, lastname string
var id, user_id, age int

func All() []Patient {
	var patients []Patient
	rows, err := database.DB.Query("SELECT * FROM patients")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&id, &firstname, &lastname, &age, &user_id); err != nil {
			fmt.Println(err)
		}
		p := Patient{ 
			Id: id, 
			FirstName: firstname, 
			LastName: lastname, 
			Age: age, 
			User_Id: user_id,
		}
		patients = append(patients, p)
	}
	return patients
}

func Find(id int) Patient {
	var p Patient
	err := database.DB.QueryRow("SELECT * FROM patients WHERE id=$1", id).
										Scan(&id, &firstname, &lastname, &age, &user_id)
	if err != nil {
		fmt.Println(err)
	}
	p = Patient{ 
		Id: id, 
		FirstName: firstname, 
		LastName: lastname, 
		Age: age, 
		User_Id: user_id,
	}
	return p
}

func Create(body io.Reader) string {
	var p Patient
	err := json.NewDecoder(body).Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	_, err = database.DB.Exec("INSERT INTO patients (firstname, lastname, age, user_id) VALUES ($1, $2, $3, $4)",
														p.FirstName, p.LastName, p.Age, p.User_Id)
	return "Patient created successfully"
}

func Update(id int, body io.Reader) string {
	var p Patient
	err := json.NewDecoder(body).Decode(&p)
	if err != nil {
		fmt.Println(err)
	}
	res, err := database.DB.Exec("UPDATE patients SET firstname=$1, lastname=$2, age=$3, user_id=$4 WHERE id = $5",
														p.FirstName, p.LastName, p.Age, p.User_Id, id)
	row, _ := res.RowsAffected()
	if row == 0 {
		return fmt.Sprintf("Patient %d not found", id)
	} else {
		return fmt.Sprintf("Patient %d updated successfully", id)
	}
}

func Delete(id int) string {
	_, err := database.DB.Exec("DELETE FROM patients WHERE id=$1", id)
	if err != nil {
		return fmt.Sprintf("Error when trying to delete patient with id %d: %v", id, err)
	} else {
		return fmt.Sprintf("Patient with id %d was deleted", id)
	}
}