package response


/*
Defines a Single Message Response structure with only 
a message sent by the server to the user 
*/
type SingleMessageResponse struct {
	Message string `json:"message"`
}


type UpdateRequest struct {
	Table string `json:"table"`

	Where map[string]interface{} `json:"where"`

	Set map[string]interface{} `json:"set"`
}
