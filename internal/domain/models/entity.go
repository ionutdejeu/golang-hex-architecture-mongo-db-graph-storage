
package models 


type EntityPropertyTriplet struct { 
	Subject interface
	Property string 
	Object string
}
type GenericEntity struct { 
	UUID string 
	Properties EntityPropertyTriplet[]
}