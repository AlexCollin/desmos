package simulation_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/desmos-labs/desmos/app"

	"github.com/desmos-labs/desmos/x/staging/posts/simulation"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func TestDecodeStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	dec := simulation.NewDecodeStore(cdc)

	timeZone, _ := time.LoadLocation("UTC")
	address := ed25519.GenPrivKey().PubKey().Address().String()

	post := types.NewPost(
		"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		"h1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		"Post message",
		types.CommentsStateAllowed,
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		nil,
		types.NewAttachments(
			types.NewAttachment(
				"https://uri.com",
				"text/plain",
				[]string{address},
			),
		),
		types.NewPollData(
			"title",
			time.Date(2100, 1, 1, 10, 0, 0, 0, timeZone),
			types.NewPollAnswers(
				types.NewPollAnswer("0", "first"),
				types.NewPollAnswer("1", "second"),
			),
			true,
			true,
		),
		time.Time{},
		time.Date(2020, 1, 1, 15, 15, 00, 000, timeZone),
		address,
	)

	postReaction := types.NewPostReaction(
		"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		"blue_heart:",
		"💙",
		address,
	)

	comment := "g1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163"

	registeredReaction := types.NewRegisteredReaction(
		address,
		":smile:",
		"https://smile.jpg",
		"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
	)

	report := types.NewReport(
		"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
		[]string{"scam"},
		"it offends me",
		address,
	)

	kvPairs := kv.Pairs{Pairs: []kv.Pair{
		{
			Key:   types.PostStoreKey(post.PostID),
			Value: cdc.MustMarshalBinaryBare(&post),
		},
		{
			Key:   types.CommentsStoreKey(post.PostID, comment),
			Value: []byte(comment),
		},
		{
			Key:   types.PostReactionsStoreKey(postReaction.PostID, postReaction.Owner, postReaction.ShortCode),
			Value: cdc.MustMarshalBinaryBare(&postReaction),
		},
		{
			Key:   types.RegisteredReactionsStoreKey(registeredReaction.Subspace, registeredReaction.ShortCode),
			Value: cdc.MustMarshalBinaryBare(&registeredReaction),
		},
		{
			Key:   types.ReportStoreKey(post.PostID, report.User),
			Value: cdc.MustMarshalBinaryBare(&report),
		},
	}}

	tests := []struct {
		name        string
		expectedLog string
	}{
		{"Post", fmt.Sprintf("PostA: %s\nPostB: %s\n", post.String(), post.String())},
		{"Comment", fmt.Sprintf("CommentA: %s\nCommentB: %s\n", comment, comment)},
		{"PostReaction", fmt.Sprintf("PostReactionA: %s\nPostReactionB: %s\n", postReaction, postReaction)},
		{"RegisteredReaction", fmt.Sprintf("RegisteredReactionA: %s\nRegisteredReactionB: %s\n", registeredReaction, registeredReaction)},
		{"Report", fmt.Sprintf("ReportA: %s\nReportB: %s\n", report, report)},
		{"other", ""},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case len(tests) - 1:
				require.Panics(t, func() { dec(kvPairs.Pairs[i], kvPairs.Pairs[i]) }, tt.name)
			default:
				require.Equal(t, tt.expectedLog, dec(kvPairs.Pairs[i], kvPairs.Pairs[i]), tt.name)
			}
		})
	}
}
