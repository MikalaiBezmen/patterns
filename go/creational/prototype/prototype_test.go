//TestPrototype - testing the Prototype

package prototype

import (
	"fmt"
	"testing"
	"../../lib/asserts"
)

func TestPrototype (t *testing.T ) { 

	prototypeFactory := PrototypeFactory{NewSoupSpoon(), NewSaladFork()}
	spoon := prototypeFactory.MakeSpoon()
	fork := prototypeFactory.MakeFork()
	asserts.Equals( t, "Getting the Spoon and Fork name:", 
		"Spoon: Soup Spoon, Fork: Salad Fork", 
		fmt.Sprintf("Spoon: %s, Fork: %s" , spoon.GetName() , fork.GetName()) )


	prototypeFactory = PrototypeFactory{NewSaladSpoon(), NewSaladFork()}
	spoon = prototypeFactory.MakeSpoon()
	fork = prototypeFactory.MakeFork()
	asserts.Equals( t, "Creating a Prototype Factory with a SaladSpoon and a SaladFork", 
		"Spoon: Salad Spoon, Fork: Salad Fork",
		fmt.Sprintf("Spoon: %s, Fork: %s" , spoon.GetName() , fork.GetName() ) )

}      
