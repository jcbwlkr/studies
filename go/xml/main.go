package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

const doc = `<Document>
  <SubDocument>
    <component>
         <section>
             <id value="0" valueType="num"/>
             <title>Foo-Type Section</title>
             <foo>FOO</foo>
         </section>
    </component>
    <component>
         <section>
             <id value="1" valueType="num"/>
             <title>Bar-Type Section</title>
             <bar>BAR</bar>
         </section>
    </component>
  </SubDocument>
</Document>`

type ID struct {
	Value int    `xml:"value,attr,omitempty"`
	Type  string `xml:"valueType,attr,omitempty"`
}

type Document struct {
	Components []Component `xml:"SubDocument>component"`
}

type Component struct {
	Section interface{} `xml:"section"`
}

type FooSection struct {
	ID    ID     `xml:"id"`
	Title string `xml:"title"`
	Foo   string `xml:"foo"`
}

type BarSection struct {
	ID    ID     `xml:"id"`
	Title string `xml:"title"`
	Bar   string `xml:"bar"`
}

func (c *Component) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	// tmp holds the data for this Component. We can only call d.DecodeElement
	// once so we have to put it somewhere so it can be reused.
	tmp := struct {
		Data []byte `xml:",innerxml"`
	}{}
	if err := d.DecodeElement(&tmp, &start); err != nil {
		return err
	}

	// which holds just enough information to tell us what kind of section to
	// make. We'll unmarshal tmp.Section.Data into this once then inspect it
	which := struct {
		ID ID `xml:"id"`
	}{}
	if err := xml.Unmarshal(tmp.Data, &which); err != nil {
		return err
	}

	switch which.ID.Value {
	case 0:
		var f FooSection
		if err := xml.Unmarshal(tmp.Data, &f); err != nil {
			return err
		}
		c.Section = f

	case 1:
		var b BarSection
		if err := xml.Unmarshal(tmp.Data, &b); err != nil {
			return err
		}
		c.Section = b
	}

	return nil
}

func main() {
	var d Document
	if err := xml.Unmarshal([]byte(doc), &d); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", d)

	b, err := xml.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
