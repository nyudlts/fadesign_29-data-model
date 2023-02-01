## Analysis Directory

This directory contains files used to determine if the Golang EAD
Package is using the correct tag and attribute strings to parse EAD
files.

#### method

Collect the set of valid EAD 2002 tag and attributes from the EAD 2002 schema:
```
wget https://www.loc.gov/ead/ead.xsd
egrep "xs:attribute|xs:element|xs:enumeration" ead.xsd | grep name | grep -v attributeGroup | sort | uniq | egrep -o "name=\"[a-z]+\"" | cut -d\" -f2 | sort | uniq  > ead-2002-tag-and-attribute-list.txt 
```

https://www.loc.gov/ead/tglib/element_index.html
https://www.loc.gov/ead/tglib/att_gen.html

Extract all of the Golang XML tags used for parsing:
```
$ pwd
/Users/jgp222/go/src/github.com/nyulibraries/dlts-finding-aids-ead-go-packages

$ grep 'xml:' * | cut -d\` -f2- | grep -v ",innerxml" | cut -d' ' -f1 | grep ',attr' | sort | uniq > xml-parsing-attrs-used-in-ead-package.txt

# clean up list (e.g., remove ",omitempty" and ",attr" tags)
$ emacs xml-parsing-attrs-used-in-ead-package.txt

# sort and uniq the xml-parsing-attrs-used-in-ead-package.txt file

# check if any tags or attributes used in the Golang package are
# missing from the EAD 2002 list of valid tags and attributes
$ for f in `cat xml-parsing-tags-used-in-ead-package.txt `; do printf "$f,";grep $f ead-2002-tag-and-attribute-list.txt ||  echo "MISSING"; done | grep MISSING
altformavailable,MISSING
```
