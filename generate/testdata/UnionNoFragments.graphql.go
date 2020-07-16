package test

// Code generated by github.com/Khan/genql, DO NOT EDIT.

import (
	"context"

	"github.com/Khan/genql/graphql"
)

type Article struct {
	Typename *string `json:"__typename"`
}

func (v Article) implementsGraphQLInterfaceLeafContent() {}

type LeafContent interface {
	implementsGraphQLInterfaceLeafContent()
}

type UnionNoFragmentsQueryResponse struct {
	RandomLeaf LeafContent `json:"-"`
}

func (v *UnionNoFragmentsQueryResponse) UnmarshalJSON(b []byte) error {
	var firstPass struct {
		*UnionNoFragmentsQueryResponse
		RandomLeaf json.RawMessage `json:"randomLeaf"`
	}
	firstPass.UnionNoFragmentsQueryResponse = v

	err := json.Unmarshal(b, &typenames)
	if err != nil {
		return err
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err = json.Unmarshal(firstPass.RandomLeaf, &tn)
	if err != nil {
		return err
	}
	switch tn.TypeName {

	case "Article":
		v.RandomLeaf = Article{}
		err = json.Unmarshal(
			firstPass.RandomLeaf, &v.RandomLeaf)

	case "Video":
		v.RandomLeaf = Video{}
		err = json.Unmarshal(
			firstPass.RandomLeaf, &v.RandomLeaf)

	}
	if err != nil {
		return err
	}

}

type Video struct {
	Typename *string `json:"__typename"`
}

func (v Video) implementsGraphQLInterfaceLeafContent() {}

func UnionNoFragmentsQuery(client *graphql.Client) (*UnionNoFragmentsQueryResponse, error) {
	var retval UnionNoFragmentsQueryResponse
	err := client.MakeRequest(context.Background(), `
query UnionNoFragmentsQuery {
	randomLeaf {
		__typename
	}
}
`, &retval, nil)
	return &retval, err
}
