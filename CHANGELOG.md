## CHANGELOG

#### v0.9.0

	* corrected element and attribute to match EAD 2002 schema https://www.loc.gov/ead/ead.xsd
	  * change element   `altformavailable` to `altformavail`
  	  * change attribute `xlink:type` to `type`
	  
    * add `analysis/` subdirectory that contains data and methods used
	  to check that all Golang `xml` element and attribute tags used in
	  the DLTS Golang EAD package have corresponding entries in the EAD
	  schema
	 
	 
