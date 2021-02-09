type Ead struct {
	XMLName        xml.Name `xml:"ead"`
	Text           string   `xml:",chardata"`
	Xmlns          string   `xml:"xmlns,attr"`
	Ns2            string   `xml:"ns2,attr"`
	Xsi            string   `xml:"xsi,attr"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Xlink          string   `xml:"xlink,attr"`
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
				Titleproper struct {
					Text string   `xml:",chardata"`
					Lb   []string `xml:"lb"`
					Num  string   `xml:"num"`
					Emph []struct {
						Text   string `xml:",chardata"`
						Render string `xml:"render,attr"`
					} `xml:"emph"`
				} `xml:"titleproper"`
				Author  string `xml:"author"`
				Sponsor string `xml:"sponsor"`
			} `xml:"titlestmt"`
			Publicationstmt struct {
				Text      string `xml:",chardata"`
				Publisher string `xml:"publisher"`
				Address   struct {
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
				P struct {
					Text string `xml:",chardata"`
					Date string `xml:"date"`
				} `xml:"p"`
			} `xml:"publicationstmt"`
		} `xml:"filedesc"`
		Profiledesc struct {
			Text     string `xml:",chardata"`
			Creation struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
			} `xml:"creation"`
			Langusage string `xml:"langusage"`
			Descrules string `xml:"descrules"`
		} `xml:"profiledesc"`
		Revisiondesc struct {
			Text   string `xml:",chardata"`
			Change []struct {
				Text string `xml:",chardata"`
				Date string `xml:"date"`
				Item string `xml:"item"`
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
				Text  string `xml:",chardata"`
				Title []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"title"`
			} `xml:"unittitle"`
			Unitid   string `xml:"unitid"`
			Physdesc []struct {
				Text      string `xml:",chardata"`
				Altrender string `xml:"altrender,attr"`
				Extent    []struct {
					Text      string `xml:",chardata"`
					Altrender string `xml:"altrender,attr"`
				} `xml:"extent"`
			} `xml:"physdesc"`
			Unitdate []struct {
				Text   string `xml:",chardata"`
				Normal string `xml:"normal,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"unitdate"`
			Container []struct {
				Text      string `xml:",chardata"`
				ID        string `xml:"id,attr"`
				Label     string `xml:"label,attr"`
				Type      string `xml:"type,attr"`
				Altrender string `xml:"altrender,attr"`
			} `xml:"container"`
			Origination []struct {
				Text     string `xml:",chardata"`
				Label    string `xml:"label,attr"`
				Corpname struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
					Rules  string `xml:"rules,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"corpname"`
				Persname struct {
					Text   string `xml:",chardata"`
					Source string `xml:"source,attr"`
					Rules  string `xml:"rules,attr"`
					Role   string `xml:"role,attr"`
				} `xml:"persname"`
			} `xml:"origination"`
			Abstract struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"abstract"`
			Physloc struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"physloc"`
			Langmaterial struct {
				Text     string `xml:",chardata"`
				ID       string `xml:"id,attr"`
				Language struct {
					Text       string `xml:",chardata"`
					Langcode   string `xml:"langcode,attr"`
					Scriptcode string `xml:"scriptcode,attr"`
				} `xml:"language"`
			} `xml:"langmaterial"`
		} `xml:"did"`
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
				Lb []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"bioghist"`
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
				Lb []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"arrangement"`
		Accessrestrict struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"accessrestrict"`
		Custodhist struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"custodhist"`
		Scopecontent struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"scopecontent"`
		Controlaccess struct {
			Text      string `xml:",chardata"`
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
			Persname []struct {
				Text   string `xml:",chardata"`
				Role   string `xml:"role,attr"`
				Source string `xml:"source,attr"`
				Rules  string `xml:"rules,attr"`
			} `xml:"persname"`
			Occupation struct {
				Text   string `xml:",chardata"`
				Source string `xml:"source,attr"`
			} `xml:"occupation"`
			Corpname []struct {
				Text   string `xml:",chardata"`
				Rules  string `xml:"rules,attr"`
				Source string `xml:"source,attr"`
				Role   string `xml:"role,attr"`
			} `xml:"corpname"`
		} `xml:"controlaccess"`
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
					} `xml:"unittitle"`
					Unitid struct {
						Text       string `xml:",chardata"`
						Audience   string `xml:"audience,attr"`
						Identifier string `xml:"identifier,attr"`
						Type       string `xml:"type,attr"`
					} `xml:"unitid"`
					Origination []struct {
						Text     string `xml:",chardata"`
						Audience string `xml:"audience,attr"`
						Label    string `xml:"label,attr"`
						Corpname struct {
							Text           string `xml:",chardata"`
							Role           string `xml:"role,attr"`
							Rules          string `xml:"rules,attr"`
							Source         string `xml:"source,attr"`
							Authfilenumber string `xml:"authfilenumber,attr"`
						} `xml:"corpname"`
						Persname struct {
							Text   string `xml:",chardata"`
							Role   string `xml:"role,attr"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"persname"`
					} `xml:"origination"`
					Unitdate []struct {
						Text   string `xml:",chardata"`
						Normal string `xml:"normal,attr"`
						Type   string `xml:"type,attr"`
					} `xml:"unitdate"`
					Container []struct {
						Text      string `xml:",chardata"`
						ID        string `xml:"id,attr"`
						Label     string `xml:"label,attr"`
						Type      string `xml:"type,attr"`
						Parent    string `xml:"parent,attr"`
						Altrender string `xml:"altrender,attr"`
					} `xml:"container"`
					Physdesc []struct {
						Text      string `xml:",chardata"`
						ID        string `xml:"id,attr"`
						Altrender string `xml:"altrender,attr"`
						Extent    []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"extent"`
						Dimensions struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"dimensions"`
					} `xml:"physdesc"`
					Langmaterial struct {
						Text     string `xml:",chardata"`
						Language struct {
							Text       string `xml:",chardata"`
							Langcode   string `xml:"langcode,attr"`
							Scriptcode string `xml:"scriptcode,attr"`
						} `xml:"language"`
					} `xml:"langmaterial"`
					Dao []struct {
						Text     string `xml:",chardata"`
						Actuate  string `xml:"actuate,attr"`
						Audience string `xml:"audience,attr"`
						Href     string `xml:"href,attr"`
						Role     string `xml:"role,attr"`
						Show     string `xml:"show,attr"`
						Title    string `xml:"title,attr"`
						Type     string `xml:"type,attr"`
						Daodesc  struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"daodesc"`
					} `xml:"dao"`
					Daogrp struct {
						Text    string `xml:",chardata"`
						Title   string `xml:"title,attr"`
						Type    string `xml:"type,attr"`
						Daodesc struct {
							Text string `xml:",chardata"`
							P    string `xml:"p"`
						} `xml:"daodesc"`
						Daoloc []struct {
							Text     string `xml:",chardata"`
							Audience string `xml:"audience,attr"`
							Href     string `xml:"href,attr"`
							Role     string `xml:"role,attr"`
							Title    string `xml:"title,attr"`
							Type     string `xml:"type,attr"`
						} `xml:"daoloc"`
					} `xml:"daogrp"`
				} `xml:"did"`
				C []struct {
					Text  string `xml:",chardata"`
					ID    string `xml:"id,attr"`
					Level string `xml:"level,attr"`
					Did   struct {
						Text      string `xml:",chardata"`
						Unittitle struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
						} `xml:"unittitle"`
						Unitid struct {
							Text       string `xml:",chardata"`
							Audience   string `xml:"audience,attr"`
							Identifier string `xml:"identifier,attr"`
							Type       string `xml:"type,attr"`
						} `xml:"unitid"`
						Origination struct {
							Text     string `xml:",chardata"`
							Audience string `xml:"audience,attr"`
							Label    string `xml:"label,attr"`
							Corpname struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"corpname"`
							Persname struct {
								Text   string `xml:",chardata"`
								Role   string `xml:"role,attr"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"persname"`
						} `xml:"origination"`
						Unitdate []struct {
							Text   string `xml:",chardata"`
							Normal string `xml:"normal,attr"`
							Type   string `xml:"type,attr"`
						} `xml:"unitdate"`
						Container []struct {
							Text      string `xml:",chardata"`
							ID        string `xml:"id,attr"`
							Label     string `xml:"label,attr"`
							Type      string `xml:"type,attr"`
							Parent    string `xml:"parent,attr"`
							Altrender string `xml:"altrender,attr"`
						} `xml:"container"`
						Dao struct {
							Text     string `xml:",chardata"`
							Actuate  string `xml:"actuate,attr"`
							Audience string `xml:"audience,attr"`
							Href     string `xml:"href,attr"`
							Show     string `xml:"show,attr"`
							Title    string `xml:"title,attr"`
							Type     string `xml:"type,attr"`
							Role     string `xml:"role,attr"`
							Daodesc  struct {
								Text string `xml:",chardata"`
								P    struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"p"`
							} `xml:"daodesc"`
						} `xml:"dao"`
						Physdesc []struct {
							Text      string `xml:",chardata"`
							Altrender string `xml:"altrender,attr"`
							ID        string `xml:"id,attr"`
							Extent    []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
							} `xml:"extent"`
						} `xml:"physdesc"`
						Langmaterial struct {
							Text     string `xml:",chardata"`
							ID       string `xml:"id,attr"`
							Language struct {
								Text       string `xml:",chardata"`
								Langcode   string `xml:"langcode,attr"`
								Scriptcode string `xml:"scriptcode,attr"`
							} `xml:"language"`
						} `xml:"langmaterial"`
					} `xml:"did"`
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
					Controlaccess struct {
						Text      string `xml:",chardata"`
						Genreform []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"genreform"`
						Subject []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"subject"`
						Geogname []struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"geogname"`
						Occupation struct {
							Text   string `xml:",chardata"`
							Source string `xml:"source,attr"`
						} `xml:"occupation"`
						Corpname struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"corpname"`
						Persname []struct {
							Text   string `xml:",chardata"`
							Rules  string `xml:"rules,attr"`
							Source string `xml:"source,attr"`
						} `xml:"persname"`
					} `xml:"controlaccess"`
					Arrangement struct {
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
					} `xml:"arrangement"`
					Scopecontent struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text string `xml:",chardata"`
							Emph []struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"emph"`
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
							} `xml:"unittitle"`
							Unitdate []struct {
								Text   string `xml:",chardata"`
								Normal string `xml:"normal,attr"`
								Type   string `xml:"type,attr"`
							} `xml:"unitdate"`
							Container []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
								ID        string `xml:"id,attr"`
								Label     string `xml:"label,attr"`
								Type      string `xml:"type,attr"`
								Parent    string `xml:"parent,attr"`
							} `xml:"container"`
							Dao []struct {
								Text     string `xml:",chardata"`
								Actuate  string `xml:"actuate,attr"`
								Audience string `xml:"audience,attr"`
								Href     string `xml:"href,attr"`
								Show     string `xml:"show,attr"`
								Title    string `xml:"title,attr"`
								Type     string `xml:"type,attr"`
								Role     string `xml:"role,attr"`
								Daodesc  struct {
									Text string `xml:",chardata"`
									P    []struct {
										Text string `xml:",chardata"`
										Emph []struct {
											Text   string `xml:",chardata"`
											Render string `xml:"render,attr"`
										} `xml:"emph"`
									} `xml:"p"`
								} `xml:"daodesc"`
							} `xml:"dao"`
							Physdesc []struct {
								Text      string `xml:",chardata"`
								Altrender string `xml:"altrender,attr"`
								ID        string `xml:"id,attr"`
								Extent    []struct {
									Text      string `xml:",chardata"`
									Altrender string `xml:"altrender,attr"`
								} `xml:"extent"`
								Emph struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
								Physfacet struct {
									Text  string `xml:",chardata"`
									ID    string `xml:"id,attr"`
									Label string `xml:"label,attr"`
								} `xml:"physfacet"`
							} `xml:"physdesc"`
							Unitid       string `xml:"unitid"`
							Langmaterial struct {
								Text     string `xml:",chardata"`
								ID       string `xml:"id,attr"`
								Language struct {
									Text       string `xml:",chardata"`
									Langcode   string `xml:"langcode,attr"`
									Scriptcode string `xml:"scriptcode,attr"`
								} `xml:"language"`
							} `xml:"langmaterial"`
							Origination struct {
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
									Rules  string `xml:"rules,attr"`
									Source string `xml:"source,attr"`
								} `xml:"persname"`
							} `xml:"origination"`
						} `xml:"did"`
						Scopecontent struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    []struct {
								Text string `xml:",chardata"`
								Emph []struct {
									Text   string `xml:",chardata"`
									Render string `xml:"render,attr"`
								} `xml:"emph"`
							} `xml:"p"`
						} `xml:"scopecontent"`
						Odd struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"odd"`
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
									Title struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"title"`
								} `xml:"unittitle"`
								Unitid   string `xml:"unitid"`
								Unitdate struct {
									Text   string `xml:",chardata"`
									Normal string `xml:"normal,attr"`
									Type   string `xml:"type,attr"`
								} `xml:"unitdate"`
								Container []struct {
									Text      string `xml:",chardata"`
									Altrender string `xml:"altrender,attr"`
									ID        string `xml:"id,attr"`
									Label     string `xml:"label,attr"`
									Type      string `xml:"type,attr"`
									Parent    string `xml:"parent,attr"`
								} `xml:"container"`
								Dao []struct {
									Text     string `xml:",chardata"`
									Actuate  string `xml:"actuate,attr"`
									Href     string `xml:"href,attr"`
									Show     string `xml:"show,attr"`
									Title    string `xml:"title,attr"`
									Type     string `xml:"type,attr"`
									Audience string `xml:"audience,attr"`
									Role     string `xml:"role,attr"`
									Daodesc  struct {
										Text string `xml:",chardata"`
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
									} `xml:"daodesc"`
								} `xml:"dao"`
								Physdesc []struct {
									Text      string `xml:",chardata"`
									ID        string `xml:"id,attr"`
									Label     string `xml:"label,attr"`
									Altrender string `xml:"altrender,attr"`
									Extent    struct {
										Text      string `xml:",chardata"`
										Altrender string `xml:"altrender,attr"`
									} `xml:"extent"`
								} `xml:"physdesc"`
								Origination struct {
									Text     string `xml:",chardata"`
									Label    string `xml:"label,attr"`
									Corpname struct {
										Text   string `xml:",chardata"`
										Rules  string `xml:"rules,attr"`
										Source string `xml:"source,attr"`
									} `xml:"corpname"`
								} `xml:"origination"`
							} `xml:"did"`
							Scopecontent struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    []struct {
									Text string `xml:",chardata"`
									Emph []struct {
										Text   string `xml:",chardata"`
										Render string `xml:"render,attr"`
									} `xml:"emph"`
								} `xml:"p"`
							} `xml:"scopecontent"`
							Odd struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"odd"`
							Separatedmaterial struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"separatedmaterial"`
							Controlaccess struct {
								Text     string `xml:",chardata"`
								Geogname []struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"geogname"`
								Subject struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"subject"`
								Genreform []struct {
									Text   string `xml:",chardata"`
									Source string `xml:"source,attr"`
								} `xml:"genreform"`
							} `xml:"controlaccess"`
							C []struct {
								Text  string `xml:",chardata"`
								ID    string `xml:"id,attr"`
								Level string `xml:"level,attr"`
								Did   struct {
									Text      string `xml:",chardata"`
									Unittitle string `xml:"unittitle"`
									Unitid    string `xml:"unitid"`
									Unitdate  struct {
										Text   string `xml:",chardata"`
										Normal string `xml:"normal,attr"`
										Type   string `xml:"type,attr"`
									} `xml:"unitdate"`
									Container []struct {
										Text      string `xml:",chardata"`
										Altrender string `xml:"altrender,attr"`
										ID        string `xml:"id,attr"`
										Label     string `xml:"label,attr"`
										Type      string `xml:"type,attr"`
										Parent    string `xml:"parent,attr"`
									} `xml:"container"`
									Dao struct {
										Text     string `xml:",chardata"`
										Actuate  string `xml:"actuate,attr"`
										Audience string `xml:"audience,attr"`
										Href     string `xml:"href,attr"`
										Role     string `xml:"role,attr"`
										Show     string `xml:"show,attr"`
										Title    string `xml:"title,attr"`
										Type     string `xml:"type,attr"`
										Daodesc  struct {
											Text string   `xml:",chardata"`
											P    []string `xml:"p"`
										} `xml:"daodesc"`
									} `xml:"dao"`
									Physdesc struct {
										Text string `xml:",chardata"`
										ID   string `xml:"id,attr"`
									} `xml:"physdesc"`
								} `xml:"did"`
								Scopecontent struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Head string `xml:"head"`
									P    []struct {
										Text string `xml:",chardata"`
										Emph []struct {
											Text   string `xml:",chardata"`
											Render string `xml:"render,attr"`
										} `xml:"emph"`
									} `xml:"p"`
								} `xml:"scopecontent"`
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
								Controlaccess struct {
									Text     string `xml:",chardata"`
									Geogname struct {
										Text   string `xml:",chardata"`
										Source string `xml:"source,attr"`
									} `xml:"geogname"`
								} `xml:"controlaccess"`
								Odd struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Head string `xml:"head"`
									P    string `xml:"p"`
								} `xml:"odd"`
							} `xml:"c"`
							Bioghist struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"bioghist"`
							Phystech struct {
								Text string `xml:",chardata"`
								ID   string `xml:"id,attr"`
								Head string `xml:"head"`
								P    string `xml:"p"`
							} `xml:"phystech"`
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
							Userestrict struct {
								Text string   `xml:",chardata"`
								ID   string   `xml:"id,attr"`
								Head string   `xml:"head"`
								P    []string `xml:"p"`
							} `xml:"userestrict"`
						} `xml:"c"`
						Controlaccess struct {
							Text      string `xml:",chardata"`
							Genreform []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"genreform"`
							Subject []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"subject"`
							Geogname []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
							} `xml:"geogname"`
							Corpname struct {
								Text   string `xml:",chardata"`
								Rules  string `xml:"rules,attr"`
								Source string `xml:"source,attr"`
							} `xml:"corpname"`
							Persname []struct {
								Text   string `xml:",chardata"`
								Source string `xml:"source,attr"`
								Rules  string `xml:"rules,attr"`
							} `xml:"persname"`
						} `xml:"controlaccess"`
						Accessrestrict struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"accessrestrict"`
						Bioghist struct {
							Text string   `xml:",chardata"`
							ID   string   `xml:"id,attr"`
							Head string   `xml:"head"`
							P    []string `xml:"p"`
						} `xml:"bioghist"`
						Phystech struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"phystech"`
						Relatedmaterial struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
							Head string `xml:"head"`
							P    string `xml:"p"`
						} `xml:"relatedmaterial"`
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
							P    string `xml:"p"`
						} `xml:"separatedmaterial"`
					} `xml:"c"`
					Accessrestrict struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"accessrestrict"`
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
						} `xml:"p"`
					} `xml:"bioghist"`
					Separatedmaterial struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    []struct {
							Text  string `xml:",chardata"`
							Title struct {
								Text   string `xml:",chardata"`
								Render string `xml:"render,attr"`
							} `xml:"title"`
						} `xml:"p"`
					} `xml:"separatedmaterial"`
					Odd struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"odd"`
					Processinfo struct {
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
					} `xml:"processinfo"`
					Accruals struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"accruals"`
					Originalsloc struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"originalsloc"`
					Relatedmaterial struct {
						Text string   `xml:",chardata"`
						ID   string   `xml:"id,attr"`
						Head string   `xml:"head"`
						P    []string `xml:"p"`
					} `xml:"relatedmaterial"`
					Phystech struct {
						Text string `xml:",chardata"`
						ID   string `xml:"id,attr"`
						Head string `xml:"head"`
						P    string `xml:"p"`
					} `xml:"phystech"`
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
				} `xml:"c"`
				Dao []struct {
					Text    string `xml:",chardata"`
					Actuate string `xml:"actuate,attr"`
					Href    string `xml:"href,attr"`
					Role    string `xml:"role,attr"`
					Show    string `xml:"show,attr"`
					Title   string `xml:"title,attr"`
					Type    string `xml:"type,attr"`
					Daodesc struct {
						Text string   `xml:",chardata"`
						P    []string `xml:"p"`
					} `xml:"daodesc"`
				} `xml:"dao"`
				Controlaccess struct {
					Text      string `xml:",chardata"`
					Genreform []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"genreform"`
					Subject []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"subject"`
					Geogname []struct {
						Text   string `xml:",chardata"`
						Source string `xml:"source,attr"`
					} `xml:"geogname"`
					Occupation []struct {
						Text           string `xml:",chardata"`
						Source         string `xml:"source,attr"`
						Authfilenumber string `xml:"authfilenumber,attr"`
					} `xml:"occupation"`
					Corpname []struct {
						Text   string `xml:",chardata"`
						Rules  string `xml:"rules,attr"`
						Source string `xml:"source,attr"`
						Role   string `xml:"role,attr"`
					} `xml:"corpname"`
					Persname struct {
						Text   string `xml:",chardata"`
						Rules  string `xml:"rules,attr"`
						Source string `xml:"source,attr"`
					} `xml:"persname"`
				} `xml:"controlaccess"`
				Arrangement struct {
					Text string   `xml:",chardata"`
					ID   string   `xml:"id,attr"`
					Head string   `xml:"head"`
					P    []string `xml:"p"`
				} `xml:"arrangement"`
				Scopecontent struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    []struct {
						Text string `xml:",chardata"`
						Emph []struct {
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
					} `xml:"p"`
				} `xml:"bioghist"`
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
				Relatedmaterial struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    struct {
						Text   string `xml:",chardata"`
						Extref struct {
							Text string `xml:",chardata"`
							Href string `xml:"href,attr"`
						} `xml:"extref"`
					} `xml:"p"`
				} `xml:"relatedmaterial"`
				Custodhist struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"custodhist"`
				Userestrict struct {
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
				} `xml:"userestrict"`
				Processinfo struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"processinfo"`
				Phystech struct {
					Text string `xml:",chardata"`
					ID   string `xml:"id,attr"`
					Head string `xml:"head"`
					P    string `xml:"p"`
				} `xml:"phystech"`
			} `xml:"c"`
		} `xml:"dsc"`
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
			} `xml:"p"`
		} `xml:"phystech"`
		Acqinfo struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"acqinfo"`
		Userestrict struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"userestrict"`
		Prefercite struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string   `xml:",chardata"`
				Lb   []string `xml:"lb"`
			} `xml:"p"`
		} `xml:"prefercite"`
		Processinfo struct {
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
			} `xml:"p"`
		} `xml:"processinfo"`
		Appraisal []struct {
			Text string   `xml:",chardata"`
			ID   string   `xml:"id,attr"`
			Head string   `xml:"head"`
			P    []string `xml:"p"`
		} `xml:"appraisal"`
		Separatedmaterial struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    []struct {
				Text string `xml:",chardata"`
				Emph []struct {
					Text   string `xml:",chardata"`
					Render string `xml:"render,attr"`
				} `xml:"emph"`
			} `xml:"p"`
		} `xml:"separatedmaterial"`
		Relatedmaterial []struct {
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
				Extref struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"extref"`
			} `xml:"p"`
		} `xml:"relatedmaterial"`
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
			P    string `xml:"p"`
		} `xml:"altformavail"`
		Odd struct {
			Text string `xml:",chardata"`
			ID   string `xml:"id,attr"`
			Head string `xml:"head"`
			P    string `xml:"p"`
		} `xml:"odd"`
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
			} `xml:"bibref"`
		} `xml:"bibliography"`
	} `xml:"archdesc"`
} 

