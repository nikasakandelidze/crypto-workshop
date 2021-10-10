package common

type InsertNodeRequest struct {
	Payload string
}

type InsertNodeResponse struct {
	Status int
}

type BlockchainValidationRequest struct {
	//no content yet for request
}

type BlockchainValidationResponse struct {
	Valid   bool
	Message string
}

type BlockchainVisualizationRequest struct {
}

type BlockchainVisualizationResponse struct {
}
