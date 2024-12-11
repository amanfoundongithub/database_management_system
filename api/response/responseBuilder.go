package response



/*
Builds a single message response for sending quick messages to
the client
*/
func CreateSingleMessageResponse(message string) SingleMessageResponse {
	return SingleMessageResponse{
		Message: message,
	}
}



