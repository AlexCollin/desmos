package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v2/x/staging/posts/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the posts MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper}
}

func computePostID(ctx sdk.Context, msg *types.MsgCreatePost) string {
	post := types.Post{
		ParentID:             msg.ParentID,
		Message:              msg.Message,
		Created:              ctx.BlockTime(),
		CommentsState:        msg.CommentsState,
		Subspace:             msg.Subspace,
		AdditionalAttributes: msg.AdditionalAttributes,
		Creator:              msg.Creator,
		Attachments:          msg.Attachments,
		Poll:                 msg.Poll,
	}

	bytes, err := post.Marshal()
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post := types.NewPost(
		computePostID(ctx, msg),
		msg.ParentID,
		msg.Message,
		msg.CommentsState,
		msg.Subspace,
		msg.AdditionalAttributes,
		msg.Attachments,
		msg.Poll,
		time.Time{},
		ctx.BlockTime(),
		msg.Creator,
	)

	// Check if the subspace exists and if the user is allowed to perform the operation on it
	if err := k.CheckUserPermissionOnSubspace(ctx, post.Subspace, post.Creator); err != nil {
		return nil, err
	}

	// Validate the post
	if err := k.ValidatePost(ctx, post); err != nil {
		return nil, err
	}

	// Check for double posting
	if k.DoesPostExist(ctx, post.PostID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the provided post conflicts with the one having id %s", post.PostID)
	}

	// Check if any of the tags have blocked the post creator
	if err := k.IsCreatorBlockedBySomeTags(ctx, post.Attachments, post.Creator, post.Subspace); err != nil {
		return nil, err
	}

	// If valid, check the parent post
	if types.IsValidPostID(post.ParentID) {
		parentPost, found := k.GetPost(ctx, post.ParentID)
		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"parent post with id %s not found", post.ParentID)
		}

		if parentPost.CommentsState == types.CommentsStateBlocked {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"post with id %s does not allow comments", parentPost.PostID)
		}

		if parentPost.Subspace != post.Subspace {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the subspace of the comment is not same as the parent post with id %s", parentPost.PostID)
		}
	}

	// Save the post
	k.SavePost(ctx, post)

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostCreated,
		sdk.NewAttribute(types.AttributeKeyPostID, post.PostID),
		sdk.NewAttribute(types.AttributeKeyPostParentID, post.ParentID),
		sdk.NewAttribute(types.AttributeKeyPostCreationTime, post.Created.Format(time.RFC3339)),
		sdk.NewAttribute(types.AttributeKeyPostOwner, post.Creator),
	))

	return &types.MsgCreatePostResponse{}, nil
}

func (k msgServer) EditPost(goCtx context.Context, msg *types.MsgEditPost) (*types.MsgEditPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the existing post
	existing, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", msg.PostID)
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Editor != existing.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check the validity of the current block height respect to the creation date of the post
	if existing.Created.After(ctx.BlockTime()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "edit date cannot be before creation date")
	}

	// Check if the subspace exists and if the user is allowed to perform the operation on it
	if err := k.CheckUserPermissionOnSubspace(ctx, existing.Subspace, existing.Creator); err != nil {
		return nil, err
	}

	// Edit the post
	existing.Message = msg.Message

	if msg.Attachments != nil {
		// Check if any of the tags have blocked the post creator
		if err := k.IsCreatorBlockedBySomeTags(ctx, msg.Attachments, existing.Creator, existing.Subspace); err != nil {
			return nil, err
		}
		existing.Attachments = msg.Attachments
	}

	if msg.Poll != nil {
		existing.Poll = msg.Poll
	}

	if msg.CommentsState != types.CommentsStateUnspecified {
		existing.CommentsState = msg.CommentsState
	}

	existing.LastEdited = ctx.BlockTime()

	if err := k.ValidatePost(ctx, existing); err != nil {
		return nil, err
	}
	k.SavePost(ctx, existing)

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostEdited,
		sdk.NewAttribute(types.AttributeKeyPostID, existing.PostID),
		sdk.NewAttribute(types.AttributeKeyPostEditTime, existing.LastEdited.Format(time.RFC3339)),
	))

	return &types.MsgEditPostResponse{}, nil
}

func (k msgServer) AddPostReaction(goCtx context.Context, msg *types.MsgAddPostReaction) (*types.MsgAddPostReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the post
	post, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", msg.PostID)
	}

	// Check if the subspace exists and if the user is allowed to perform the operation on it
	if err := k.CheckUserPermissionOnSubspace(ctx, post.Subspace, msg.User); err != nil {
		return nil, err
	}

	reactionShortcode, reactionValue, err := k.ExtractReactionValueAndShortcode(ctx, msg.Reaction, post.Subspace)
	if err != nil {
		return nil, err
	}

	postReaction := types.NewPostReaction(msg.PostID, reactionShortcode, reactionValue, msg.User)
	if err := k.SavePostReaction(ctx, postReaction); err != nil {
		return nil, err
	}

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostReactionAdded,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID),
		sdk.NewAttribute(types.AttributeKeyPostReactionOwner, msg.User),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reactionValue),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reactionShortcode),
	))

	return &types.MsgAddPostReactionResponse{}, nil
}

func (k msgServer) RemovePostReaction(goCtx context.Context, msg *types.MsgRemovePostReaction) (*types.MsgRemovePostReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the post
	post, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", msg.PostID)
	}

	// Check if the subspace exists and if the user is allowed to perform the operation on it
	if err := k.CheckUserPermissionOnSubspace(ctx, post.Subspace, msg.User); err != nil {
		return nil, err
	}

	reactionShortcode, reactionValue, err := k.ExtractReactionValueAndShortcode(ctx, msg.Reaction, post.Subspace)
	if err != nil {
		return nil, err
	}

	// Remove the registeredReactions
	reaction := types.NewPostReaction(msg.PostID, reactionShortcode, reactionValue, msg.User)
	if err := k.DeletePostReaction(ctx, reaction); err != nil {
		return nil, err
	}

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostReactionRemoved,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID),
		sdk.NewAttribute(types.AttributeKeyPostReactionOwner, msg.User),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reactionValue),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reactionShortcode),
	))

	return &types.MsgRemovePostReactionResponse{}, nil
}

func (k msgServer) RegisterReaction(goCtx context.Context, msg *types.MsgRegisterReaction) (*types.MsgRegisterReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists and if the user is allowed to perform the operation on it
	if err := k.CheckUserPermissionOnSubspace(ctx, msg.Subspace, msg.Creator); err != nil {
		return nil, err
	}

	// Check if the shortcode is associated with an emoji
	if _, found := types.GetEmojiByShortCodeOrValue(msg.ShortCode); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"shortcode %s represents an emoji and thus can't be used to register a new registeredReactions", msg.ShortCode)
	}

	// Make sure the given reaction isn't already registered
	if _, isAlreadyRegistered := k.GetRegisteredReaction(ctx, msg.ShortCode, msg.Subspace); isAlreadyRegistered {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"registeredReactions with shortcode %s and subspace %s has already been registered", msg.ShortCode, msg.Subspace)
	}

	reaction := types.NewRegisteredReaction(msg.Creator, msg.ShortCode, msg.Value, msg.Subspace)
	k.SaveRegisteredReaction(ctx, reaction)

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRegisterReaction,
		sdk.NewAttribute(types.AttributeKeyReactionCreator, reaction.Creator),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reaction.ShortCode),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reaction.Value),
		sdk.NewAttribute(types.AttributeKeyReactionSubSpace, reaction.Subspace),
	))

	return &types.MsgRegisterReactionResponse{}, nil
}

func (k msgServer) AnswerPoll(goCtx context.Context, msg *types.MsgAnswerPoll) (*types.MsgAnswerPollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks if the post exists
	post, found := k.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"post with id %s doesn't exist", msg.PostID)
	}

	// Check if the subspace exists and if the user is allowed to perform the operation on it
	if err := k.CheckUserPermissionOnSubspace(ctx, post.Subspace, msg.Answerer); err != nil {
		return nil, err
	}

	// Make sure the post has a poll
	if post.Poll == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"no poll associated with ID: %s", msg.PostID)
	}

	// Make sure the poll is not closed
	if post.Poll.EndDate.Before(ctx.BlockTime()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the poll associated with ID %s was closed at %s", post.PostID, post.Poll.EndDate)
	}

	// Check if the poll allows multiple answers
	if len(msg.Answers) > 1 && !post.Poll.AllowsMultipleAnswers {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the poll associated with ID %s doesn't allow multiple answers", post.PostID)
	}

	// Check if the user answers are more than the answers provided by the poll
	if len(msg.Answers) > len(post.Poll.ProvidedAnswers) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			"user's answers are more than the available ones inside the poll")
	}

	// Make sure that each answer provided by the user matches with one of the provided ones by the poll creator
	for _, answer := range msg.Answers {
		var found = false
		for _, providedAnswer := range post.Poll.ProvidedAnswers {
			if answer == providedAnswer.ID {
				found = true
				break
			}
		}

		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"answer with ID %s isn't one of the poll's provided answers", answer)
		}
	}

	_, found = k.GetUserAnswer(ctx, post.PostID, msg.Answerer)

	// Check if the poll allows to edit previous answers
	if found && !post.Poll.AllowsAnswerEdits {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"post with ID %s doesn't allow answers' edits", post.PostID)
	}

	userAnswer := types.NewUserAnswer(post.PostID, msg.Answerer, msg.Answers)
	k.SaveUserAnswer(ctx, userAnswer)

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAnsweredPoll,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID),
		sdk.NewAttribute(types.AttributeKeyPollAnswerer, msg.Answerer),
	))

	return &types.MsgAnswerPollResponse{}, nil
}

func (k msgServer) ReportPost(goCtx context.Context, msg *types.MsgReportPost) (*types.MsgReportPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the post to report exists
	postID := msg.PostID
	if !types.IsValidPostID(postID) {
		return nil, sdkerrors.Wrap(types.ErrInvalidPostID, postID)
	}

	var post types.Post
	var exist bool
	if post, exist = k.GetPost(ctx, postID); !exist {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with ID: %s doesn't exist", postID)
	}

	// Check if the subspace exists and if the user is allowed to perform the operation on it
	if err := k.CheckUserPermissionOnSubspace(ctx, post.Subspace, msg.User); err != nil {
		return nil, err
	}

	// Create and store the report
	report := types.NewReport(postID, msg.ReportType, msg.Message, msg.User)
	err := k.SaveReport(ctx, report)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostReported,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID),
		sdk.NewAttribute(types.AttributeKeyReportOwner, msg.User),
	))

	return &types.MsgReportPostResponse{}, nil
}
