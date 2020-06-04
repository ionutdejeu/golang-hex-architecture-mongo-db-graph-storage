package dynamoadapter 

type EntityPropertyTriplet struct { 
	Subject interface{} 	`json:"subject"` 
	Property string 		`json:"property"`
	ObjectValue string		`json:"objectvalue"`
}

type EntityStoreEntity struct { 
	id string 
	Properties EntityPropertyTriplet[]
}
