package structs

//Post data RISK_PROFILE
type RiskProfile struct{
	UserID 			int 		`json:"userid"`
	MM				float32 	`json:"mm"`
	Bond		 	float32 	`json:"bond"`
	Stock		 	float32 	`json:"stock"`
}

//Post data User
type User struct{
	UserID 			int 			`json:"userid" gorm:"primary_key"`
	NAME 			string 			`json:"name"`
	AGE				int				`json:"age"`
	RiskProfiles	[]RiskProfile 	`json:"details,omitempty" gorm:"foreignkey:UserID"`
	
}

// Result is an array of post
type Result struct {
	Code 	int 		`json:"code"`
	Data 	interface{} `json:"data"`
	Message string 		`json:"message"`
}