package model

import "github.com/g8rswimmer/go-twitter/v2"

var (
	AllExpansions = []twitter.Expansion{
		twitter.ExpansionAttachmentsPollIDs,
		twitter.ExpansionAttachmentsMediaKeys,
		twitter.ExpansionAuthorID,
		twitter.ExpansionEntitiesMentionsUserName,
		twitter.ExpansionGeoPlaceID,
		twitter.ExpansionInReplyToUserID,
		twitter.ExpansionReferencedTweetsID,
		twitter.ExpansionReferencedTweetsIDAuthorID,
		twitter.ExpansionPinnedTweetID,
		twitter.ExpansionOwnerID,
		twitter.ExpansionCreatorID,
		twitter.ExpansionSpeakerIDS,
		twitter.ExpansionInvitedUserIDs,
		twitter.ExpansionHostIDs,
	}
	ALLTweetFields = []twitter.TweetField{
		twitter.TweetFieldID,
		twitter.TweetFieldText,
		twitter.TweetFieldAttachments,
		twitter.TweetFieldAuthorID,
		twitter.TweetFieldContextAnnotations,
		twitter.TweetFieldConversationID,
		twitter.TweetFieldCreatedAt,
		twitter.TweetFieldEntities,
		twitter.TweetFieldGeo,
		twitter.TweetFieldInReplyToUserID,
		twitter.TweetFieldLanguage,
		twitter.TweetFieldNonPublicMetrics,
		twitter.TweetFieldPublicMetrics,
		twitter.TweetFieldOrganicMetrics,
		twitter.TweetFieldPromotedMetrics,
		twitter.TweetFieldPossiblySensitve,
		twitter.TweetFieldReferencedTweets,
		twitter.TweetFieldSource,
		twitter.TweetFieldWithHeld,
	}
	AllUserFields = []twitter.UserField{
		twitter.UserFieldCreatedAt,
		twitter.UserFieldDescription,
		twitter.UserFieldLocation,
		twitter.UserFieldName,
		twitter.UserFieldPinnedTweetID,
		twitter.UserFieldProfileImageURL,
		twitter.UserFieldProtected,
		twitter.UserFieldPublicMetrics,
		twitter.UserFieldURL,
		twitter.UserFieldUserName,
		twitter.UserFieldVerified,
		twitter.UserFieldWithHeld,
	}
	AllMediaFields = []twitter.MediaField{
		twitter.MediaFieldDurationMS,
		twitter.MediaFieldHeight,
		twitter.MediaFieldMediaKey,
		twitter.MediaFieldPreviewImageURL,
		twitter.MediaFieldType,
		twitter.MediaFieldURL,
		twitter.MediaFieldWidth,
		twitter.MediaFieldPublicMetrics,
		twitter.MediaFieldNonPublicMetrics,
		twitter.MediaFieldOrganicMetrics,
		twitter.MediaFieldPromotedMetrics,
		twitter.MediaFieldAltText,
		twitter.MediaFieldVariants,
	}
	AllPlaceFields = []twitter.PlaceField{
		twitter.PlaceFieldContainedWithin,
		twitter.PlaceFieldCountry,
		twitter.PlaceFieldCountryCode,
		twitter.PlaceFieldFullName,
		twitter.PlaceFieldGeo,
		twitter.PlaceFieldID,
		twitter.PlaceFieldName,
		twitter.PlaceFieldPlaceType,
	}
	AllPollFields = []twitter.PollField{
		twitter.PollFieldDurationMinutes,
		twitter.PollFieldEndDateTime,
		twitter.PollFieldID,
		twitter.PollFieldOptions,
		twitter.PollFieldVotingStatus,
	}
)
