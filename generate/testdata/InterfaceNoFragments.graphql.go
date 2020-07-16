package test

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"context"

	"github.com/Khan/genql/graphql"
)

type Article struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (v Article) implementsGraphQLInterfaceContent() {}

type Content interface {
	implementsGraphQLInterfaceContent()
}

type InterfaceNoFragmentsQueryResponse struct {
	Root Topic `json:"root"`
}

type Topic struct {
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Children []Content `json:"-"`
}

func (v *Topic) UnmarshalJSON(b []byte) error {
	var firstPass struct {
		*Topic
		Children json.RawMessage `json:"children"`
	}
	firstPass.Topic = v

	err := json.Unmarshal(b, &typenames)
	if err != nil {
		return err
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err = json.Unmarshal(firstPass.Children, &tn)
	if err != nil {
		return err
	}
	switch tn.TypeName {

	case "Article":
		v.Children = Article{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Video":
		v.Children = Video{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	case "Topic":
		v.Children = Topic{}
		err = json.Unmarshal(
			firstPass.Children, &v.Children)

	}
	if err != nil {
		return err
	}

}

type Video struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (v Video) implementsGraphQLInterfaceContent() {}

func InterfaceNoFragmentsQuery(client *graphql.Client) (*InterfaceNoFragmentsQueryResponse, error) {
	var retval InterfaceNoFragmentsQueryResponse
	err := client.MakeRequest(context.Background(), `
query InterfaceNoFragmentsQuery {
	root {
		id
		name
		children {
			id
			name
		}
	}
}
`, &retval, nil)
	return &retval, err
}
