type Ead struct {
	XMLName        xml.Name `xml:"ead"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Xlink          string   `xml:"xlink,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Eadheader      struct {
		Text               string `xml:",chardata"`
		Countryencoding    string `xml:"countryencoding,attr"`
		Dateencoding       string `xml:"dateencoding,attr"`
		Findaidstatus      string `xml:"findaidstatus,attr"`
		Langencoding       string `xml:"langencoding,attr"`
		Repositoryencoding string `xml:"repositoryencoding,attr"`
		Eadid              struct {
			Text           string `xml:",chardata"`
			Countrycode    string `xml:"countrycode,attr"`
			Mainagencycode string `xml:"mainagencycode,attr"`
			URL            string `xml:"url,attr"`
		} `xml:"eadid"`
		Filedesc struct {
			Text      string `xml:",chardata"`
			Titlestmt struct {
				Text        string `xml:",chardata"`
				Titleproper []struct {
					Text string   `xml:",chardata"`
					Type string   `xml:"type,attr"`
					Num  []string `xml:"num"`
					Lb   []string `xml:"lb"`
					Date []struct {
						Text     string `xml:",chardata"`
						Normal   string `xml:"normal,attr"`
						Type     string `xml:"type,attr"`
						Calendar string `xml:"calendar,attr"`
						Era      string `xml:"era,attr"`
					} `xml:"date"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"titleproper"`
				Author   string `xml:"author"`
				Subtitle string `xml:"subtitle"`
				Sponsor  string `xml:"sponsor"`
			} `xml:"titlestmt"`
			Publicationstmt struct {
				Text      string `xml:",chardata"`
				Publisher string `xml:"publisher"`
				P         struct {
					Text string `xml:",chardata"`
					Date string `xml:"date"`
				} `xml:"p"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text   string `xml:",chardata"`
						Extptr struct {
							Text  string `xml:",chardata"`
							Href  string `xml:"href,attr"`
							Show  string `xml:"show,attr"`
							Title string `xml:"title,attr"`
							Type  string `xml:"type,attr"`
						} `xml:"extptr"`
					} `xml:"addressline"`
				} `xml:"address"`
			} `xml:"publicationstmt"`
			Editionstmt struct {
				Text string `xml:",chardata"`
				P    struct {
					Text string `xml:",chardata"`
					Date struct {
						Text     string `xml:",chardata"`
						Calendar string `xml:"calendar,attr"`
						Era      string `xml:"era,attr"`
					} `xml:"date"`
				} `xml:"p"`
			} `xml:"editionstmt"`
			Notestmt struct {
				Text string `xml:",chardata"`
				Note struct {
					Text string `xml:",chardata"`
					P    string `xml:"p"`
				} `xml:"note"`
			} `xml:"notestmt"`
		} `xml:"filedesc"`
		Profiledesc struct {
			Text     string `xml:",chardata"`
			Creation struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"creation"`
			Langusage struct {
				Text     string `xml:",chardata"`
				Language struct {
					Text           string `xml:",chardata"`
					Langcode       string `xml:"langcode,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"language"`
			} `xml:"langusage"`
			Descrules string `xml:"descrules"`
		} `xml:"profiledesc"`
		Revisiondesc struct {
			Text   string `xml:",chardata"`
			Change []struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				Item struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"item"`
			} `xml:"change"`
		} `xml:"revisiondesc"`
	} `xml:"eadheader"`
	Archdesc struct {
		Text  string `xml:",chardata"`
		Level string `xml:"level,attr"`
		Did   struct {
			Text       string `xml:",chardata"`
			Repository struct {
				Text     string `xml:",chardata"`
				Corpname string `xml:"corpname"`
			} `xml:"repository"`
			Unittitle struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Lb string `xml:"lb"`
			} `xml:"unittitle"`
			Origination []struct {
				Text     string `xml:",chardata"`
				Label    string `xml:"label,attr"`
				Corpname struct {
					Text   string `xml:",chardata"`
					Rules  string `xml:"rules,attr"`
					Source string `xml:"source,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"corpname"`
				Persname struct {
					Text   string `xml:",chardata"`
					Rules  string `xml:"rules,attr"`
					Source string `xml:"source,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"persname"`
				Famname struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
					Rules  string `xml:"rules,attr"`
				} `xml:"famname"`
			} `xml:"origination"`
			Unitid   string `xml:"unitid"`
			Physdesc []struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
				ID        string `xml:"id,attr"`
				Label     string `xml:"label,attr"`
				Extent    []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
				} `xml:"extent"`
				Physfacet string `xml:"physfacet"`
			} `xml:"physdesc"`
			Unitdate []struct {
				Text   string `xml:",chardata"`
				Normal string `xml:"normal,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"unitdate"`
			Langmaterial []struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Language struct {
					Text       string `xml:",chardata"`
					Langcode   string `xml:"langcode,attr"`
					Scriptcode string `xml:"scriptcode,attr"`
				} `xml:"language"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"langmaterial"`
			Abstract []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Lb []string `xml:"lb"`
			} `xml:"abstract"`
			Physloc []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"physloc"`
			Container []struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
				ID        string `xml:"id,attr"`
				Label     string `xml:"label,attr"`
				Type      string `xml:"type,attr"`
				Parent    string `xml:"parent,attr"`
			} `xml:"container"`
		} `xml:"did"`
		Phystech []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Lb []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"phystech"`
		Dsc struct {
			Text string `xml:",chardata"`
			C    []struct {
				Text       string `xml:",chardata"`
				ID         string `xml:"id,attr"`
				Level      string `xml:"level,attr"`
				Otherlevel string `xml:"otherlevel,attr"`
				Did        struct {
					Text      string `xml:",chardata"`
					Unittitle struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Name struct {
							Text string `xml:",chardata"`
							Role string `xml:"role,attr"`
						} `xml:"name"`
						Lb       []string `xml:"lb"`
						Corpname string   `xml:"corpname"`
						Persname string   `xml:"persname"`
					} `xml:"unittitle"`
					Langmaterial struct {
						Text     string `xml:",chardata"`
						ID       string `xml:"id,attr"`
						Language []struct {
							Text       string `xml:",chardata"`
							Langcode   string `xml:"langcode,attr"`
							Scriptcode string `xml:"scriptcode,attr"`
						} `xml:"language"`
					} `xml:"langmaterial"`
					Unitdate []struct {
						Text      string `xml:",chardata"`
						Normal    string `xml:"normal,attr"`
						Type      string `xml:"type,attr"`
						Certainty string `xml:"certainty,attr"`
					} `xml:"unitdate"`
					Container []struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Type      string `xml:"type,attr"`
						Parent    string `xml:"parent,attr"`
					} `xml:"container"`
					Physdesc []struct {
						Text      string `xml:",chardata"`
						Altrender string `xml:"altrender,attr"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Extent    []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"extent"`
						Physfacet string `xml:"physfacet"`
						Emph      struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"physdesc"`
					Unitid string `xml:"unitid"`
					Dao    []struct {
						Text    string `xml:",chardata"`
						Actuate string `xml:"actuate,attr"`
						Href    string `xml:"href,attr"`
						Role    string `xml:"role,attr"`
						Show    string `xml:"show,attr"`
						Title   string `xml:"title,attr"`
						Type    string `xml:"type,attr"`
						Daodesc struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"daodesc"`
					} `xml:"dao"`
					Origination []struct {
						Text     string `xml:",chardata"`
						Label    string `xml:"label,attr"`
						Persname struct {
							Text   string `xml:",chardata"`
							Role   string `xml:"role,attr"`
							Source string `xml:"source,attr"`
							Rules  string `xml:"rules,attr"`
						} `xml:"persname"`
						Corpname struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"corpname"`
					} `xml:"origination"`
					Physloc struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"physloc"`
					Abstract struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
					} `xml:"abstract"`
					Daogrp struct {
						Text    string `xml:",chardata"`
						Title   string `xml:"title,attr"`
						Type    string `xml:"type,attr"`
						Daodesc struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"daodesc"`
						Daoloc []struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
							Role string `xml:"role,attr"`
							Type string `xml:"type,attr"`
						} `xml:"daoloc"`
					} `xml:"daogrp"`
				} `xml:"did"`
				C []struct {
					Text       string `xml:",chardata"`
					ID         string `xml:"id,attr"`
					Level      string `xml:"level,attr"`
					Otherlevel string `xml:"otherlevel,attr"`
					Did        struct {
						Text      string `xml:",chardata"`
						Unittitle struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
								Emph   struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Lb []string `xml:"lb"`
							} `xml:"emph"`
							Lb    []string `xml:"lb"`
							Title []struct {
								Text   string   `xml:",chardata"`
								Render string   `xml:"render,attr"`
								Lb     []string `xml:"lb"`
							} `xml:"title"`
							Name struct {
								Text string `xml:",chardata"`
								Role string `xml:"role,attr"`
							} `xml:"name"`
							Corpname string `xml:"corpname"`
						} `xml:"unittitle"`
						Unitid   string `xml:"unitid"`
						Physdesc []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Extent    []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
								Unit      string `xml:"unit,attr"`
							} `xml:"extent"`
							Dimensions struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Label string `xml:"label,attr"`
							} `xml:"dimensions"`
							Physfacet struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Label string `xml:"label,attr"`
								Emph  struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"physfacet"`
						} `xml:"physdesc"`
						Unitdate []struct {
							Text      string `xml:",chardata"`
							Normal    string `xml:"normal,attr"`
							Type      string `xml:"type,attr"`
							Certainty string `xml:"certainty,attr"`
						} `xml:"unitdate"`
						Langmaterial struct {
							Text     string `xml:",chardata"`
							ID       string `xml:"id,attr"`
							Language struct {
								Text       string `xml:",chardata"`
								Langcode   string `xml:"langcode,attr"`
								Scriptcode string `xml:"scriptcode,attr"`
							} `xml:"language"`
						} `xml:"langmaterial"`
						Container []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Type      string `xml:"type,attr"`
							Parent    string `xml:"parent,attr"`
						} `xml:"container"`
						Origination []struct {
							Text     string `xml:",chardata"`
							Label    string `xml:"label,attr"`
							Corpname struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"corpname"`
							Persname struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"persname"`
						} `xml:"origination"`
						Dao []struct {
							Text    string `xml:",chardata"`
							Actuate string `xml:"actuate,attr"`
							Href    string `xml:"href,attr"`
							Role    string `xml:"role,attr"`
							Show    string `xml:"show,attr"`
							Title   string `xml:"title,attr"`
							Type    string `xml:"type,attr"`
							Daodesc struct {
								Text string `xml:",chardata"`
								P    struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
									Name string `xml:"name"`
									Lb   string `xml:"lb"`
								} `xml:"p"`
							} `xml:"daodesc"`
						} `xml:"dao"`
						Physloc struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"physloc"`
						Abstract struct {
							Text string   `xml:",chardata"`
							ID   string   `xml:"id,attr"`
							Lb   []string `xml:"lb"`
						} `xml:"abstract"`
						Materialspec struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"materialspec"`
					} `xml:"did"`
					Altformavail []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"altformavail"`
					Scopecontent []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head struct {
							Text string `xml:",chardata"`
							Lb   string `xml:"lb"`
						} `xml:"head"`
						P []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string   `xml:",chardata"`
								Render string   `xml:"render,attr"`
								Lb     []string `xml:"lb"`
							} `xml:"emph"`
							Lb    []string `xml:"lb"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"p"`
					} `xml:"scopecontent"`
					Phystech []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head struct {
							Text string `xml:",chardata"`
							Lb   string `xml:"lb"`
						} `xml:"head"`
						P []struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Lb []string `xml:"lb"`
						} `xml:"p"`
					} `xml:"phystech"`
					Odd []struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Lb []string `xml:"lb"`
						} `xml:"p"`
					} `xml:"odd"`
					Relatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text   string `xml:",chardata"`
							Extref []struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
								Show    string `xml:"show,attr"`
							} `xml:"extref"`
						} `xml:"p"`
					} `xml:"relatedmaterial"`
					Controlaccess struct {
						Text     string `xml:",chardata"`
						Persname []struct {
							Text   string `xml:",chardata"`
							Role   string `xml:"role,attr"`
							Source string `xml:"source,attr"`
							Rules  string `xml:"rules,attr"`
						} `xml:"persname"`
						Corpname []struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
							Role   string `xml:"role,attr"`
						} `xml:"corpname"`
						Geogname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"geogname"`
						Subject []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"subject"`
						Genreform []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"genreform"`
						Title []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"title"`
						Famname struct {
							Text   string `xml:",chardata"`
							Role   string `xml:"role,attr"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"famname"`
					} `xml:"controlaccess"`
					Custodhist struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"custodhist"`
					Otherfindaid struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"otherfindaid"`
					Bioghist struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Title []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"p"`
					} `xml:"bioghist"`
					Accessrestrict []struct {
						Text        string `xml:",chardata"`
						ID          string `xml:"id,attr"`
						Legalstatus struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"legalstatus"`
						Head string   `xml:"head"`
						P    []string `xml:"p"`
					} `xml:"accessrestrict"`
					Userestrict struct {
						Text string   `xml:",chardata"`
						ID   string   `xml:"id,attr"`
						Head string   `xml:"head"`
						P    []string `xml:"p"`
					} `xml:"userestrict"`
					Arrangement struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
							Lb []string `xml:"lb"`
						} `xml:"p"`
					} `xml:"arrangement"`
					C []struct {
						Text       string `xml:",chardata"`
						ID         string `xml:"id,attr"`
						Level      string `xml:"level,attr"`
						Otherlevel string `xml:"otherlevel,attr"`
						Did        struct {
							Text      string `xml:",chardata"`
							Unittitle struct {
								Text  string `xml:",chardata"`
								Title []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
									Emph   struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"title"`
								Lb   []string `xml:"lb"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Name     []string `xml:"name"`
								Persname []string `xml:"persname"`
							} `xml:"unittitle"`
							Unitdate []struct {
								Text      string `xml:",chardata"`
								Normal    string `xml:"normal,attr"`
								Type      string `xml:"type,attr"`
								Certainty string `xml:"certainty,attr"`
							} `xml:"unitdate"`
							Container []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
								ID        string `xml:"id,attr"`
								Label     string `xml:"label,attr"`
								Type      string `xml:"type,attr"`
								Parent    string `xml:"parent,attr"`
							} `xml:"container"`
							Langmaterial struct {
								Text     string `xml:",chardata"`
								ID       string `xml:"id,attr"`
								Language struct {
									Text       string `xml:",chardata"`
									Langcode   string `xml:"langcode,attr"`
									Scriptcode string `xml:"scriptcode,attr"`
								} `xml:"language"`
							} `xml:"langmaterial"`
							Physdesc []struct {
								Text      string `xml:",chardata"`
								ID        string `xml:"id,attr"`
								Label     string `xml:"label,attr"`
								Altrender string `xml:"altrender,attr"`
								Extent    []struct {
									Text      string `xml:",chardata"`
									Altrender string `xml:"altrender,attr"`
									Lb        string `xml:"lb"`
								} `xml:"extent"`
								Dimensions struct {
									Text  string `xml:",chardata"`
									ID    string `xml:"id,attr"`
									Label string `xml:"label,attr"`
								} `xml:"dimensions"`
								Physfacet struct {
									Text  string `xml:",chardata"`
									ID    string `xml:"id,attr"`
									Label string `xml:"label,attr"`
									Emph  struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"physfacet"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"physdesc"`
							Unitid  string `xml:"unitid"`
							Physloc struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"physloc"`
							Origination []struct {
								Text     string `xml:",chardata"`
								Label    string `xml:"label,attr"`
								Persname struct {
									Text   string `xml:",chardata"`
									Rules  string `xml:"rules,attr"`
									Source string `xml:"source,attr"`
									Role   string `xml:"role,attr"`
								} `xml:"persname"`
								Corpname struct {
									Text   string `xml:",chardata"`
									Role   string `xml:"role,attr"`
									Rules  string `xml:"rules,attr"`
									Source string `xml:"source,attr"`
								} `xml:"corpname"`
							} `xml:"origination"`
							Dao []struct {
								Text    string `xml:",chardata"`
								Actuate string `xml:"actuate,attr"`
								Href    string `xml:"href,attr"`
								Role    string `xml:"role,attr"`
								Show    string `xml:"show,attr"`
								Title   string `xml:"title,attr"`
								Type    string `xml:"type,attr"`
								Daodesc struct {
									Text string `xml:",chardata"`
									P    []struct {
										Text string `xml:",chardata"`
										Emph []struct {
											Text   string `xml:",chardata"`
											Render string `xml:"render,attr"`
										} `xml:"emph"`
										Name  []string `xml:"name"`
										Lb    []string `xml:"lb"`
										Title []struct {
											Text   string `xml:",chardata"`
											Render string `xml:"render,attr"`
										} `xml:"title"`
									} `xml:"p"`
								} `xml:"daodesc"`
							} `xml:"dao"`
							Abstract struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
							} `xml:"abstract"`
							Materialspec struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"materialspec"`
						} `xml:"did"`
						Odd []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    []struct {
								Text string   `xml:",chardata"`
								Lb   []string `xml:"lb"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Title string `xml:"title"`
								Num   struct {
									Text string `xml:",chardata"`
									Type string `xml:"type,attr"`
								} `xml:"num"`
							} `xml:"p"`
						} `xml:"odd"`
						Scopecontent []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    []struct {
								Text string `xml:",chardata"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Lb    []string `xml:"lb"`
								Title []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
							} `xml:"p"`
						} `xml:"scopecontent"`
						C []struct {
							Text       string `xml:",chardata"`
							ID         string `xml:"id,attr"`
							Level      string `xml:"level,attr"`
							Otherlevel string `xml:"otherlevel,attr"`
							Did        struct {
								Text      string `xml:",chardata"`
								Unittitle struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
									Lb    string `xml:"lb"`
									Title []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
								} `xml:"unittitle"`
								Unitdate []struct {
									Text   string `xml:",chardata"`
									Normal string `xml:"normal,attr"`
									Type   string `xml:"type,attr"`
								} `xml:"unitdate"`
								Langmaterial struct {
									Text     string `xml:",chardata"`
									Language struct {
										Text       string `xml:",chardata"`
										Langcode   string `xml:"langcode,attr"`
										Scriptcode string `xml:"scriptcode,attr"`
									} `xml:"language"`
								} `xml:"langmaterial"`
								Container []struct {
									Text      string `xml:",chardata"`
									ID        string `xml:"id,attr"`
									Label     string `xml:"label,attr"`
									Type      string `xml:"type,attr"`
									Parent    string `xml:"parent,attr"`
									Altrender string `xml:"altrender,attr"`
								} `xml:"container"`
								Unitid   string `xml:"unitid"`
								Physdesc []struct {
									Text      string `xml:",chardata"`
									ID        string `xml:"id,attr"`
									Label     string `xml:"label,attr"`
									Altrender string `xml:"altrender,attr"`
									Extent    []struct {
										Text      string `xml:",chardata"`
										Altrender string `xml:"altrender,attr"`
									} `xml:"extent"`
								} `xml:"physdesc"`
								Dao []struct {
									Text    string `xml:",chardata"`
									Actuate string `xml:"actuate,attr"`
									Href    string `xml:"href,attr"`
									Role    string `xml:"role,attr"`
									Show    string `xml:"show,attr"`
									Title   string `xml:"title,attr"`
									Type    string `xml:"type,attr"`
									Daodesc struct {
										Text string `xml:",chardata"`
										P    string `xml:"p"`
									} `xml:"daodesc"`
								} `xml:"dao"`
								Origination []struct {
									Text     string `xml:",chardata"`
									Label    string `xml:"label,attr"`
									Persname struct {
										Text   string `xml:",chardata"`
										Rules  string `xml:"rules,attr"`
										Source string `xml:"source,attr"`
									} `xml:"persname"`
									Corpname struct {
										Text   string `xml:",chardata"`
										Rules  string `xml:"rules,attr"`
										Source string `xml:"source,attr"`
									} `xml:"corpname"`
								} `xml:"origination"`
								Physloc struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
								} `xml:"physloc"`
							} `xml:"did"`
							Odd []struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    []struct {
									Text string `xml:",chardata"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
									Title struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
								} `xml:"p"`
							} `xml:"odd"`
							Controlaccess struct {
								Text     string `xml:",chardata"`
								Geogname []struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"geogname"`
								Genreform struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"genreform"`
								Subject struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"subject"`
							} `xml:"controlaccess"`
							Userestrict struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"userestrict"`
							Phystech struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"phystech"`
							C []struct {
								Text       string `xml:",chardata"`
								ID         string `xml:"id,attr"`
								Level      string `xml:"level,attr"`
								Otherlevel string `xml:"otherlevel,attr"`
								Did        struct {
									Text      string `xml:",chardata"`
									Unittitle string `xml:"unittitle"`
									Unitdate  struct {
										Text   string `xml:",chardata"`
										Normal string `xml:"normal,attr"`
										Type   string `xml:"type,attr"`
									} `xml:"unitdate"`
									Langmaterial struct {
										Text     string `xml:",chardata"`
										Language struct {
											Text     string `xml:",chardata"`
											Langcode string `xml:"langcode,attr"`
										} `xml:"language"`
									} `xml:"langmaterial"`
									Container []struct {
										Text      string `xml:",chardata"`
										ID        string `xml:"id,attr"`
										Label     string `xml:"label,attr"`
										Type      string `xml:"type,attr"`
										Altrender string `xml:"altrender,attr"`
										Parent    string `xml:"parent,attr"`
									} `xml:"container"`
									Dao struct {
										Text    string `xml:",chardata"`
										Actuate string `xml:"actuate,attr"`
										Href    string `xml:"href,attr"`
										Role    string `xml:"role,attr"`
										Show    string `xml:"show,attr"`
										Title   string `xml:"title,attr"`
										Type    string `xml:"type,attr"`
										Daodesc struct {
											Text string `xml:",chardata"`
											P    string `xml:"p"`
										} `xml:"daodesc"`
									} `xml:"dao"`
									Unitid string `xml:"unitid"`
								} `xml:"did"`
								Odd struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Head string `xml:"head"`
									P    string `xml:"p"`
								} `xml:"odd"`
							} `xml:"c"`
							Scopecontent struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    struct {
									Text string `xml:",chardata"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"p"`
							} `xml:"scopecontent"`
							Accessrestrict struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"accessrestrict"`
						} `xml:"c"`
						Separatedmaterial struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    struct {
								Text string   `xml:",chardata"`
								Lb   []string `xml:"lb"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Genreform struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"genreform"`
								Archref struct {
									Text    string `xml:",chardata"`
									Physloc string `xml:"physloc"`
								} `xml:"archref"`
							} `xml:"p"`
						} `xml:"separatedmaterial"`
						Accessrestrict []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    []struct {
								Text string `xml:",chardata"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"p"`
						} `xml:"accessrestrict"`
						Processinfo struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"processinfo"`
						Altformavail struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    struct {
								Text string `xml:",chardata"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"p"`
						} `xml:"altformavail"`
						Phystech []struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    []struct {
								Text string `xml:",chardata"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Lb []string `xml:"lb"`
							} `xml:"p"`
						} `xml:"phystech"`
						Controlaccess struct {
							Text     string `xml:",chardata"`
							Corpname []struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"corpname"`
							Persname []struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"persname"`
							Genreform []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"genreform"`
							Geogname []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"geogname"`
							Subject []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"subject"`
							Title []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"title"`
							Occupation struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"occupation"`
						} `xml:"controlaccess"`
						Bioghist struct {
							Text string   `xml:",chardata"`
							ID   string   `xml:"id,attr"`
							Head string   `xml:"head"`
							P    []string `xml:"p"`
						} `xml:"bioghist"`
						Acqinfo struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"acqinfo"`
						Userestrict struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"userestrict"`
						Relatedmaterial struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"relatedmaterial"`
						Fileplan struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    struct {
								Text string `xml:",chardata"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"p"`
						} `xml:"fileplan"`
						Originalsloc struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"originalsloc"`
						Otherfindaid struct {
							Text string   `xml:",chardata"`
							ID   string   `xml:"id,attr"`
							Head string   `xml:"head"`
							P    []string `xml:"p"`
						} `xml:"otherfindaid"`
						Arrangement struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"arrangement"`
						Custodhist struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"custodhist"`
						Accruals struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"accruals"`
						Appraisal struct {
							Text string   `xml:",chardata"`
							ID   string   `xml:"id,attr"`
							Head string   `xml:"head"`
							P    []string `xml:"p"`
						} `xml:"appraisal"`
					} `xml:"c"`
					Separatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text string   `xml:",chardata"`
							Lb   []string `xml:"lb"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"separatedmaterial"`
					Processinfo struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"p"`
					} `xml:"processinfo"`
					Originalsloc struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"originalsloc"`
					Fileplan struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"fileplan"`
					Appraisal struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"appraisal"`
					Accruals struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"accruals"`
					Index struct {
						Text       string `xml:",chardata"`
						ID         string `xml:"id,attr"`
						Head       string `xml:"head"`
						Indexentry []struct {
							Text     string `xml:",chardata"`
							Name     string `xml:"name"`
							Subject  string `xml:"subject"`
							Corpname string `xml:"corpname"`
						} `xml:"indexentry"`
					} `xml:"index"`
					Prefercite struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"prefercite"`
				} `xml:"c"`
				Scopecontent []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text     string `xml:",chardata"`
						Corpname string `xml:"corpname"`
						Date     []struct {
							Text     string `xml:",chardata"`
							Calendar string `xml:"calendar,attr"`
							Era      string `xml:"era,attr"`
						} `xml:"date"`
						Subject    []string `xml:"subject"`
						Genreform  []string `xml:"genreform"`
						Occupation string   `xml:"occupation"`
						Emph       []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Lb    []string `xml:"lb"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Extref []struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
						List struct {
							Text string `xml:",chardata"`
							Item []struct {
								Text string `xml:",chardata"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"item"`
						} `xml:"list"`
					} `xml:"p"`
				} `xml:"scopecontent"`
				Arrangement struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Lb    []string `xml:"lb"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Blockquote struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"blockquote"`
					} `xml:"p"`
				} `xml:"arrangement"`
				Accessrestrict struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"p"`
				} `xml:"accessrestrict"`
				Odd struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Lb []string `xml:"lb"`
					} `xml:"p"`
				} `xml:"odd"`
				Processinfo struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"processinfo"`
				Phystech []struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"phystech"`
				Accruals struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"accruals"`
				Appraisal struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"appraisal"`
				Bioghist struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"p"`
				} `xml:"bioghist"`
				Controlaccess struct {
					Text      string `xml:",chardata"`
					Genreform []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
					Persname []struct {
						Text   string `xml:",chardata"`
						Role   string `xml:"role,attr"`
						Source string `xml:"source,attr"`
						Rules  string `xml:"rules,attr"`
					} `xml:"persname"`
					Geogname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"geogname"`
					Subject []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"subject"`
					Corpname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
						Rules  string `xml:"rules,attr"`
					} `xml:"corpname"`
					Occupation []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"occupation"`
					Function struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"function"`
					Title struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"title"`
				} `xml:"controlaccess"`
				Otherfindaid struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"otherfindaid"`
				Userestrict struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"userestrict"`
				Separatedmaterial struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"p"`
				} `xml:"separatedmaterial"`
				Prefercite struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"prefercite"`
				Custodhist struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"custodhist"`
				Originalsloc struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"originalsloc"`
				Relatedmaterial struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"relatedmaterial"`
				Altformavail struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"altformavail"`
				Acqinfo struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"acqinfo"`
			} `xml:"c"`
		} `xml:"dsc"`
		Accessrestrict []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Extref struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Archref string `xml:"archref"`
				} `xml:"extref"`
				Lb []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"accessrestrict"`
		Acqinfo struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"acqinfo"`
		Arrangement []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				List struct {
					Text       string `xml:",chardata"`
					Numeration string `xml:"numeration,attr"`
					Type       string `xml:"type,attr"`
					Item       []struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
					} `xml:"item"`
				} `xml:"list"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Lb []string `xml:"lb"`
				} `xml:"emph"`
				Lb    []string `xml:"lb"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Extref struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
			List struct {
				Text       string `xml:",chardata"`
				Numeration string `xml:"numeration,attr"`
				Type       string `xml:"type,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
					Lb []string `xml:"lb"`
				} `xml:"item"`
			} `xml:"list"`
		} `xml:"arrangement"`
		Bioghist []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"emph"`
				List []struct {
					Text string `xml:",chardata"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"title"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Bibref struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"bibref"`
					} `xml:"item"`
					Head struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"head"`
				} `xml:"list"`
				Persname []struct {
					Text   string `xml:",chardata"`
					Normal string `xml:"normal,attr"`
				} `xml:"persname"`
				Corpname []struct {
					Text           string `xml:",chardata"`
					Source         string `xml:"source,attr"`
					Encodinganalog string `xml:"encodinganalog,attr"`
				} `xml:"corpname"`
				Geogname []string `xml:"geogname"`
				Date     []struct {
					Text     string `xml:",chardata"`
					Calendar string `xml:"calendar,attr"`
					Era      string `xml:"era,attr"`
				} `xml:"date"`
				Subject    []string `xml:"subject"`
				Occupation []string `xml:"occupation"`
				Title      []struct {
					Text   string `xml:",chardata"`
					Type   string `xml:"type,attr"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"title"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
					Emph    struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"extref"`
				Lb         []string `xml:"lb"`
				Blockquote struct {
					Text string `xml:",chardata"`
					P    string `xml:"p"`
				} `xml:"blockquote"`
				Name      []string `xml:"name"`
				Chronlist struct {
					Text      string `xml:",chardata"`
					Chronitem []struct {
						Text string `xml:",chardata"`
						Date struct {
							Text     string `xml:",chardata"`
							Calendar string `xml:"calendar,attr"`
							Era      string `xml:"era,attr"`
						} `xml:"date"`
						Event struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"event"`
						Eventgrp struct {
							Text  string `xml:",chardata"`
							Event []struct {
								Text  string `xml:",chardata"`
								Title struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"title"`
							} `xml:"event"`
						} `xml:"eventgrp"`
					} `xml:"chronitem"`
					Head struct {
						Text string `xml:",chardata"`
						Emph struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"head"`
				} `xml:"chronlist"`
				Abbr  string `xml:"abbr"`
				Table struct {
					Text   string `xml:",chardata"`
					Tgroup struct {
						Text    string `xml:",chardata"`
						Cols    string `xml:"cols,attr"`
						Colspec []struct {
							Text     string `xml:",chardata"`
							Colwidth string `xml:"colwidth,attr"`
						} `xml:"colspec"`
						Tbody struct {
							Text string `xml:",chardata"`
							Row  []struct {
								Text  string `xml:",chardata"`
								Entry []struct {
									Text string `xml:",chardata"`
									Emph struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"entry"`
							} `xml:"row"`
						} `xml:"tbody"`
					} `xml:"tgroup"`
				} `xml:"table"`
				Bibref struct {
					Text  string `xml:",chardata"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"bibref"`
			} `xml:"p"`
			List []struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Numeration string `xml:"numeration,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text string `xml:",chardata"`
					Name string `xml:"name"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"item"`
				Defitem []struct {
					Text  string `xml:",chardata"`
					Label string `xml:"label"`
					Item  struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"item"`
				} `xml:"defitem"`
			} `xml:"list"`
			Chronlist struct {
				Text      string `xml:",chardata"`
				Head      string `xml:"head"`
				Chronitem []struct {
					Text     string `xml:",chardata"`
					Date     string `xml:"date"`
					Eventgrp struct {
						Text  string `xml:",chardata"`
						Event []struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
							Emph struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"event"`
					} `xml:"eventgrp"`
				} `xml:"chronitem"`
			} `xml:"chronlist"`
		} `xml:"bioghist"`
		Custodhist []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Lb     []string `xml:"lb"`
				Extref struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"custodhist"`
		Prefercite []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Lb      []string `xml:"lb"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text string `xml:",chardata"`
						Lb   string `xml:"lb"`
					} `xml:"addressline"`
				} `xml:"address"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"prefercite"`
		Processinfo []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Lb    []string `xml:"lb"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"processinfo"`
		Scopecontent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Emph   struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"title"`
				} `xml:"emph"`
				Corpname []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"corpname"`
				Date struct {
					Text     string `xml:",chardata"`
					Calendar string `xml:"calendar,attr"`
					Era      string `xml:"era,attr"`
				} `xml:"date"`
				List []struct {
					Text string `xml:",chardata"`
					Head string `xml:"head"`
					Item []struct {
						Text  string `xml:",chardata"`
						Title []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
						Emph []struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
						Lb []string `xml:"lb"`
					} `xml:"item"`
				} `xml:"list"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Genreform  []string `xml:"genreform"`
				Name       []string `xml:"name"`
				Occupation string   `xml:"occupation"`
				Lb         []string `xml:"lb"`
				Extref     struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
				Archref struct {
					Text   string `xml:",chardata"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
				Persname string `xml:"persname"`
			} `xml:"p"`
			List []struct {
				Text string   `xml:",chardata"`
				Type string   `xml:"type,attr"`
				Head string   `xml:"head"`
				Item []string `xml:"item"`
			} `xml:"list"`
		} `xml:"scopecontent"`
		Userestrict struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text    string   `xml:",chardata"`
				Lb      []string `xml:"lb"`
				Address struct {
					Text        string `xml:",chardata"`
					Addressline []struct {
						Text string   `xml:",chardata"`
						Lb   []string `xml:"lb"`
					} `xml:"addressline"`
				} `xml:"address"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Extref string `xml:"extref"`
			} `xml:"p"`
		} `xml:"userestrict"`
		Originalsloc struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"originalsloc"`
		Controlaccess struct {
			Text    string `xml:",chardata"`
			Subject []struct {
				Text           string `xml:",chardata"`
				Source         string `xml:"source,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"subject"`
			Genreform []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"genreform"`
			Geogname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"geogname"`
			Corpname []struct {
				Text           string `xml:",chardata"`
				Rules          string `xml:"rules,attr"`
				Source         string `xml:"source,attr"`
				Role           string `xml:"role,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"corpname"`
			Persname []struct {
				Text           string `xml:",chardata"`
				Rules          string `xml:"rules,attr"`
				Source         string `xml:"source,attr"`
				Role           string `xml:"role,attr"`
				Authfilenumber string `xml:"authfilenumber,attr"`
			} `xml:"persname"`
			Famname []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
				Rules  string `xml:"rules,attr"`
				Role   string `xml:"role,attr"`
			} `xml:"famname"`
			Title []struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"title"`
			Occupation struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"occupation"`
		} `xml:"controlaccess"`
		Otherfindaid []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Extref struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"otherfindaid"`
		Separatedmaterial []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				List []struct {
					Text string `xml:",chardata"`
					Head string `xml:"head"`
					Item []struct {
						Text  string   `xml:",chardata"`
						Lb    []string `xml:"lb"`
						Title struct {
							Text   string `xml:",chardata"`
							Type   string `xml:"type,attr"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"item"`
				} `xml:"list"`
				Archref struct {
					Text   string `xml:",chardata"`
					Extref struct {
						Text string `xml:",chardata"`
						Href string `xml:"href,attr"`
					} `xml:"extref"`
				} `xml:"archref"`
				Lb    []string `xml:"lb"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
					Type   string `xml:"type,attr"`
				} `xml:"title"`
				Extref struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
			List []struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				Head string `xml:"head"`
				Item []struct {
					Text string `xml:",chardata"`
					Emph struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"item"`
			} `xml:"list"`
		} `xml:"separatedmaterial"`
		Relatedmaterial []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text    string `xml:",chardata"`
				Archref struct {
					Text   string `xml:",chardata"`
					Extref struct {
						Text  string `xml:",chardata"`
						Href  string `xml:"href,attr"`
						Title struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"title"`
					} `xml:"extref"`
				} `xml:"archref"`
				Extref []struct {
					Text    string `xml:",chardata"`
					Href    string `xml:"href,attr"`
					Title   string `xml:"title,attr"`
					Actuate string `xml:"actuate,attr"`
					Show    string `xml:"show,attr"`
				} `xml:"extref"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				List struct {
					Text string   `xml:",chardata"`
					Item []string `xml:"item"`
				} `xml:"list"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
				Lb []string `xml:"lb"`
			} `xml:"p"`
			List struct {
				Text string   `xml:",chardata"`
				Type string   `xml:"type,attr"`
				Head string   `xml:"head"`
				Item []string `xml:"item"`
			} `xml:"list"`
		} `xml:"relatedmaterial"`
		Odd []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				List struct {
					Text string   `xml:",chardata"`
					Item []string `xml:"item"`
				} `xml:"list"`
			} `xml:"p"`
			List struct {
				Text       string `xml:",chardata"`
				Type       string `xml:"type,attr"`
				Numeration string `xml:"numeration,attr"`
				Head       string `xml:"head"`
				Item       []struct {
					Text     string   `xml:",chardata"`
					Persname []string `xml:"persname"`
					Corpname string   `xml:"corpname"`
				} `xml:"item"`
			} `xml:"list"`
		} `xml:"odd"`
		Appraisal []struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"appraisal"`
		Accruals struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accruals"`
		Altformavail struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Extref []struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
				Title struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"p"`
		} `xml:"altformavail"`
		Bibliography struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Head   string `xml:"head"`
			Bibref []struct {
				Text string `xml:",chardata"`
				Emph struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"bibref"`
			P []struct {
				Text   string `xml:",chardata"`
				Bibref []struct {
					Text string `xml:",chardata"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
					Title []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
						Emph   struct {
							Text   string `xml:",chardata"`
							Render string `xml:"render,attr"`
						} `xml:"emph"`
					} `xml:"title"`
				} `xml:"bibref"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"bibliography"`
	} `xml:"archdesc"`
}

