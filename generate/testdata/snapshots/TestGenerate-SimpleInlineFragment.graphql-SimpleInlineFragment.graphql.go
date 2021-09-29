package test

// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

import (
	"encoding/json"
	"fmt"

	"github.com/Khan/genqlient/graphql"
	"github.com/Khan/genqlient/internal/testutil"
)

// SimpleInlineFragmentRandomItemArticle includes the requested fields of the GraphQL type Article.
type SimpleInlineFragmentRandomItemArticle struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
	Text string      `json:"text"`
}

// SimpleInlineFragmentRandomItemContent includes the requested fields of the GraphQL interface Content.
//
// SimpleInlineFragmentRandomItemContent is implemented by the following types:
// SimpleInlineFragmentRandomItemArticle
// SimpleInlineFragmentRandomItemVideo
// SimpleInlineFragmentRandomItemTopic
// The GraphQL type's documentation follows.
//
// Content is implemented by various types like Article, Video, and Topic.
type SimpleInlineFragmentRandomItemContent interface {
	implementsGraphQLInterfaceSimpleInlineFragmentRandomItemContent()
	// GetTypename returns the receiver's concrete GraphQL type-name (see interface doc for possible values).
	GetTypename() string
	// GetId returns the interface-field "id" from its implementation.
	// The GraphQL interface field's documentation follows.
	//
	// ID is the identifier of the content.
	GetId() testutil.ID
	// GetName returns the interface-field "name" from its implementation.
	GetName() string
}

func (v *SimpleInlineFragmentRandomItemArticle) implementsGraphQLInterfaceSimpleInlineFragmentRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemArticle) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemArticle) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemArticle) GetName() string { return v.Name }

func (v *SimpleInlineFragmentRandomItemVideo) implementsGraphQLInterfaceSimpleInlineFragmentRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemVideo) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemVideo) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemVideo) GetName() string { return v.Name }

func (v *SimpleInlineFragmentRandomItemTopic) implementsGraphQLInterfaceSimpleInlineFragmentRandomItemContent() {
}

// GetTypename is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemTopic) GetTypename() string { return v.Typename }

// GetId is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemTopic) GetId() testutil.ID { return v.Id }

// GetName is a part of, and documented with, the interface SimpleInlineFragmentRandomItemContent.
func (v *SimpleInlineFragmentRandomItemTopic) GetName() string { return v.Name }

func __unmarshalSimpleInlineFragmentRandomItemContent(b []byte, v *SimpleInlineFragmentRandomItemContent) error {
	if string(b) == "null" {
		return nil
	}

	var tn struct {
		TypeName string `json:"__typename"`
	}
	err := json.Unmarshal(b, &tn)
	if err != nil {
		return err
	}

	switch tn.TypeName {
	case "Article":
		*v = new(SimpleInlineFragmentRandomItemArticle)
		return json.Unmarshal(b, *v)
	case "Video":
		*v = new(SimpleInlineFragmentRandomItemVideo)
		return json.Unmarshal(b, *v)
	case "Topic":
		*v = new(SimpleInlineFragmentRandomItemTopic)
		return json.Unmarshal(b, *v)
	case "":
		return fmt.Errorf(
			"Response was missing Content.__typename")
	default:
		return fmt.Errorf(
			`Unexpected concrete type for SimpleInlineFragmentRandomItemContent: "%v"`, tn.TypeName)
	}
}

func __marshalSimpleInlineFragmentRandomItemContent(v *SimpleInlineFragmentRandomItemContent) ([]byte, error) {

	var typename string
	switch v := (*v).(type) {
	case *SimpleInlineFragmentRandomItemArticle:
		typename = "Article"

		result := struct {
			TypeName string `json:"__typename"`
			*SimpleInlineFragmentRandomItemArticle
		}{typename, v}
		return json.Marshal(result)
	case *SimpleInlineFragmentRandomItemVideo:
		typename = "Video"

		result := struct {
			TypeName string `json:"__typename"`
			*SimpleInlineFragmentRandomItemVideo
		}{typename, v}
		return json.Marshal(result)
	case *SimpleInlineFragmentRandomItemTopic:
		typename = "Topic"

		result := struct {
			TypeName string `json:"__typename"`
			*SimpleInlineFragmentRandomItemTopic
		}{typename, v}
		return json.Marshal(result)
	case nil:
		return []byte("null"), nil
	default:
		return nil, fmt.Errorf(
			`Unexpected concrete type for SimpleInlineFragmentRandomItemContent: "%T"`, v)
	}
}

// SimpleInlineFragmentRandomItemTopic includes the requested fields of the GraphQL type Topic.
type SimpleInlineFragmentRandomItemTopic struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id   testutil.ID `json:"id"`
	Name string      `json:"name"`
}

// SimpleInlineFragmentRandomItemVideo includes the requested fields of the GraphQL type Video.
type SimpleInlineFragmentRandomItemVideo struct {
	Typename string `json:"__typename"`
	// ID is the identifier of the content.
	Id       testutil.ID `json:"id"`
	Name     string      `json:"name"`
	Duration int         `json:"duration"`
}

// SimpleInlineFragmentResponse is returned by SimpleInlineFragment on success.
type SimpleInlineFragmentResponse struct {
	RandomItem SimpleInlineFragmentRandomItemContent `json:"-"`
}

func (v *SimpleInlineFragmentResponse) UnmarshalJSON(b []byte) error {

	if string(b) == "null" {
		return nil
	}

	var firstPass struct {
		*SimpleInlineFragmentResponse
		RandomItem json.RawMessage `json:"randomItem"`
		graphql.NoUnmarshalJSON
	}
	firstPass.SimpleInlineFragmentResponse = v

	err := json.Unmarshal(b, &firstPass)
	if err != nil {
		return err
	}

	{
		dst := &v.RandomItem
		src := firstPass.RandomItem
		if len(src) != 0 && string(src) != "null" {
			err = __unmarshalSimpleInlineFragmentRandomItemContent(
				src, dst)
			if err != nil {
				return fmt.Errorf(
					"Unable to unmarshal SimpleInlineFragmentResponse.RandomItem: %w", err)
			}
		}
	}
	return nil
}

type __premarshalSimpleInlineFragmentResponse struct {
	RandomItem json.RawMessage `json:"randomItem"`
}

func (v *SimpleInlineFragmentResponse) MarshalJSON() ([]byte, error) {
	premarshaled, err := v.__premarshalJSON()
	if err != nil {
		return nil, err
	}
	return json.Marshal(premarshaled)
}

func (v *SimpleInlineFragmentResponse) __premarshalJSON() (*__premarshalSimpleInlineFragmentResponse, error) {
	var retval __premarshalSimpleInlineFragmentResponse

	{

		dst := &retval.RandomItem
		src := v.RandomItem
		var err error
		*dst, err = __marshalSimpleInlineFragmentRandomItemContent(
			&src)
		if err != nil {
			return nil, fmt.Errorf(
				"Unable to marshal SimpleInlineFragmentResponse.RandomItem: %w", err)
		}
	}
	return &retval, nil
}

func SimpleInlineFragment(
	client graphql.Client,
) (*SimpleInlineFragmentResponse, error) {
	var err error

	var retval SimpleInlineFragmentResponse
	err = client.MakeRequest(
		nil,
		"SimpleInlineFragment",
		`
query SimpleInlineFragment {
	randomItem {
		__typename
		id
		name
		... on Article {
			text
		}
		... on Video {
			duration
		}
	}
}
`,
		&retval,
		nil,
	)
	return &retval, err
}

