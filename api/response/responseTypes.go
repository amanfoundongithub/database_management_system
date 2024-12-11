package response

/*
Defines an update request to the server 
*/
type UpdateRequest struct {
	Table string `json:"table"`
	Where map[string]interface{} `json:"where"`
	Set map[string]interface{} `json:"set"`
}

/*
Defines an delete request to the server 
*/
type TableDeleteRequest struct {
	Table string `json:"table"`

	Password string `json:"password"`
}


/*
Defines a Single Message Response with only
a message field 
*/
type SingleMessageResponse struct {
	Message string `json:"message"`
}