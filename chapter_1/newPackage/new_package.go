// A package is a collection of files in a folder, hence why this new package is in a new folder called "newPackage"
package newPackage

import "fmt"

var VariableFromOtherPackage = "fskdvk"

const ConstFromOtherPackage = 123

func PrintPackageValue() {
	fmt.Println("Value being printed from other package")
}
