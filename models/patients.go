package patients

// Note the fields in the struct must be capitalized otherwise they won't be exported
// and json.Marshal won't be able to access those fields.

type Patient struct {
	Id        int
	FirstName string
	LastName  string
	Age       int
	User_Id   int
}

var AllPatients = []Patient{
	{Id: 1, FirstName: "Lee", LastName: "C", Age: 33, User_Id: 1},
	{Id: 2, FirstName: "Jay", LastName: "C", Age: 33, User_Id: 2},
	{Id: 3, FirstName: "Mom", LastName: "C", Age: 33, User_Id: 3},
	{Id: 4, FirstName: "Ba", LastName: "C", Age: 33, User_Id: 4},
}

func GetAllPatients() []Patient {
	return AllPatients
}

func GetOnePatient(id int) Patient {
	var p Patient
	for _, patient := range AllPatients {
		if patient.Id == id {
			p = patient
			break
		}
	}
	return p
}
