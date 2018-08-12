package patient

// Note the fields in the struct must be capitalized otherwise they won't be exported
// and json.Marshal won't be able to access those fields.

type Patient struct {
	Id        int			`json:"id"`
	FirstName string	`json:"first_name"`
	LastName  string	`json:"last_name"`
	Age       int			`json:"age"`
	User_Id   int			`json:"user_id"`
}

var Patients = []Patient{
	{Id: 1, FirstName: "Lee", LastName: "C", Age: 33, User_Id: 1},
	{Id: 2, FirstName: "Jay", LastName: "C", Age: 33, User_Id: 2},
	{Id: 3, FirstName: "Mom", LastName: "C", Age: 33, User_Id: 3},
	{Id: 4, FirstName: "Ba", LastName: "C", Age: 33, User_Id: 4},
}

func All() []Patient {
	return Patients
}

func Find(id int) Patient {
	var patient Patient
	for _, p := range Patients {
		if p.Id == id {
			patient = p
		}
	}
	return patient
}