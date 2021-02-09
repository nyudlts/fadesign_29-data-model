type Ead struct {
	Eadheader      struct {
		Filedesc struct {} `xml:"filedesc"`
		Profiledesc struct {} `xml:"profiledesc"`
		Revisiondesc struct {} `xml:"revisiondesc"`
	} `xml:"eadheader"`
	Archdesc struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
		Did   struct {} `xml:"did"`
		Phystech []struct {} `xml:"phystech"`
		Dsc struct {
			Text string `xml:",chardata"`
			C    []struct {
		} `xml:"dsc"`
		Accessrestrict []struct {} `xml:"accessrestrict"`
		Acqinfo struct {} `xml:"acqinfo"`
		Arrangement []struct {} `xml:"arrangement"`
		Bioghist []struct {} `xml:"bioghist"`
		Custodhist []struct {} `xml:"custodhist"`
		Prefercite []struct {} `xml:"prefercite"`
		Processinfo []struct {} `xml:"processinfo"`
		Scopecontent struct {} `xml:"scopecontent"`
		Userestrict struct {} `xml:"userestrict"`
		Originalsloc struct {} `xml:"originalsloc"`
		Controlaccess struct {} `xml:"controlaccess"`
		Otherfindaid []struct {} `xml:"otherfindaid"`
		Separatedmaterial []struct {} `xml:"separatedmaterial"`
		Relatedmaterial []struct {} `xml:"relatedmaterial"`
		Odd []struct {} `xml:"odd"`
		Appraisal []struct {} `xml:"appraisal"`
		Accruals struct {} `xml:"accruals"`
		Altformavail struct {} `xml:"altformavail"`
		Bibliography struct {} `xml:"bibliography"`
	} `xml:"archdesc"`
} 

