package dtos

type LoginInput struct {
	Email    string 
	Password string 
}

type LoginOutput struct {
	Message string 
}

type SignupInput struct {
	Email       string  
	Password    string  
	Firstname   string  
	Lastname    string  
	PhoneNumber *string 
}
