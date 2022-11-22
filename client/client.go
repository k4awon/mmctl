// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package client

import (
	"io"
	"net/http"

	"github.com/mattermost/mattermost-server/v6/model"
)

type Client interface {
	CreateChannel(channel *model.Channel) (*model.Channel, *model.Response, error)
	RemoveUserFromChannel(channelID, userID string) (*model.Response, error)
	GetChannelMembers(channelID string, page, perPage int, etag string) (model.ChannelMembers, *model.Response, error)
	AddChannelMember(channelID, userID string) (*model.ChannelMember, *model.Response, error)
	DeleteChannel(channelID string) (*model.Response, error)
	PermanentDeleteChannel(channelID string) (*model.Response, error)
	MoveChannel(channelID, teamID string, force bool) (*model.Channel, *model.Response, error)
	GetPublicChannelsForTeam(teamID string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetDeletedChannelsForTeam(teamID string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetPrivateChannelsForTeam(teamID string, page int, perPage int, etag string) ([]*model.Channel, *model.Response, error)
	GetChannelsForTeamForUser(teamID, userID string, includeDeleted bool, etag string) ([]*model.Channel, *model.Response, error)
	RestoreChannel(channelID string) (*model.Channel, *model.Response, error)
	PatchChannel(channelID string, patch *model.ChannelPatch) (*model.Channel, *model.Response, error)
	GetChannelByName(channelName, teamID string, etag string) (*model.Channel, *model.Response, error)
	GetChannelByNameIncludeDeleted(channelName, teamID string, etag string) (*model.Channel, *model.Response, error)
	GetChannel(channelID, etag string) (*model.Channel, *model.Response, error)
	GetTeam(teamID, etag string) (*model.Team, *model.Response, error)
	GetTeamByName(name, etag string) (*model.Team, *model.Response, error)
	GetAllTeams(etag string, page int, perPage int) ([]*model.Team, *model.Response, error)
	CreateTeam(team *model.Team) (*model.Team, *model.Response, error)
	PatchTeam(teamID string, patch *model.TeamPatch) (*model.Team, *model.Response, error)
	AddTeamMember(teamID, userID string) (*model.TeamMember, *model.Response, error)
	RemoveTeamMember(teamID, userID string) (*model.Response, error)
	SoftDeleteTeam(teamID string) (*model.Response, error)
	PermanentDeleteTeam(teamID string) (*model.Response, error)
	RestoreTeam(teamID string) (*model.Team, *model.Response, error)
	UpdateTeamPrivacy(teamID string, privacy string) (*model.Team, *model.Response, error)
	SearchTeams(search *model.TeamSearch) ([]*model.Team, *model.Response, error)
	GetPost(postID string, etag string) (*model.Post, *model.Response, error)
	CreatePost(post *model.Post) (*model.Post, *model.Response, error)
	GetPostsForChannel(channelID string, page, perPage int, etag string, collapsedThreads bool) (*model.PostList, *model.Response, error)
	GetPostsSince(channelID string, since int64, collapsedThreads bool) (*model.PostList, *model.Response, error)
	DoAPIPost(url string, data string) (*http.Response, error)
	GetLdapGroups() ([]*model.Group, *model.Response, error)
	GetGroupsByChannel(channelID string, groupOpts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response, error)
	GetGroupsByTeam(teamID string, groupOpts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response, error)
	UploadLicenseFile(data []byte) (*model.Response, error)
	RemoveLicenseFile() (*model.Response, error)
	GetLogs(page, perPage int) ([]string, *model.Response, error)
	GetRoleByName(name string) (*model.Role, *model.Response, error)
	PatchRole(roleID string, patch *model.RolePatch) (*model.Role, *model.Response, error)
	UploadPlugin(file io.Reader) (*model.Manifest, *model.Response, error)
	UploadPluginForced(file io.Reader) (*model.Manifest, *model.Response, error)
	RemovePlugin(id string) (*model.Response, error)
	EnablePlugin(id string) (*model.Response, error)
	DisablePlugin(id string) (*model.Response, error)
	GetPlugins() (*model.PluginsResponse, *model.Response, error)
	GetUser(userID, etag string) (*model.User, *model.Response, error)
	GetUserByUsername(userName, etag string) (*model.User, *model.Response, error)
	GetUserByEmail(email, etag string) (*model.User, *model.Response, error)
	PermanentDeleteUser(userID string) (*model.Response, error)
	PermanentDeleteAllUsers() (*model.Response, error)
	CreateUser(user *model.User) (*model.User, *model.Response, error)
	VerifyUserEmailWithoutToken(userID string) (*model.User, *model.Response, error)
	UpdateUserRoles(userID, roles string) (*model.Response, error)
	InviteUsersToTeam(teamID string, userEmails []string) (*model.Response, error)
	SendPasswordResetEmail(email string) (*model.Response, error)
	UpdateUser(user *model.User) (*model.User, *model.Response, error)
	UpdateUserMfa(userID, code string, activate bool) (*model.Response, error)
	UpdateUserPassword(userID, currentPassword, newPassword string) (*model.Response, error)
	UpdateUserHashedPassword(userID, newHashedPassword string) (*model.Response, error)
	CreateUserAccessToken(userID, description string) (*model.UserAccessToken, *model.Response, error)
	RevokeUserAccessToken(tokenID string) (*model.Response, error)
	GetUserAccessTokensForUser(userID string, page, perPage int) ([]*model.UserAccessToken, *model.Response, error)
	ConvertUserToBot(userID string) (*model.Bot, *model.Response, error)
	ConvertBotToUser(userID string, userPatch *model.UserPatch, setSystemAdmin bool) (*model.User, *model.Response, error)
	PromoteGuestToUser(userID string) (*model.Response, error)
	DemoteUserToGuest(guestID string) (*model.Response, error)
	CreateCommand(cmd *model.Command) (*model.Command, *model.Response, error)
	ListCommands(teamID string, customOnly bool) ([]*model.Command, *model.Response, error)
	GetCommandById(cmdID string) (*model.Command, *model.Response, error)
	UpdateCommand(cmd *model.Command) (*model.Command, *model.Response, error)
	MoveCommand(teamID string, commandID string) (*model.Response, error)
	DeleteCommand(commandID string) (*model.Response, error)
	GetConfig() (*model.Config, *model.Response, error)
	UpdateConfig(*model.Config) (*model.Config, *model.Response, error)
	PatchConfig(*model.Config) (*model.Config, *model.Response, error)
	ReloadConfig() (*model.Response, error)
	MigrateConfig(from, to string) (*model.Response, error)
	SyncLdap(includeRemovedMembers bool) (*model.Response, error)
	MigrateIdLdap(toAttribute string) (*model.Response, error)
	GetUsers(page, perPage int, etag string) ([]*model.User, *model.Response, error)
	GetUsersByIds(userIDs []string) ([]*model.User, *model.Response, error)
	GetUsersInTeam(teamID string, page, perPage int, etag string) ([]*model.User, *model.Response, error)
	UpdateUserActive(userID string, activate bool) (*model.Response, error)
	UpdateTeam(team *model.Team) (*model.Team, *model.Response, error)
	UpdateChannelPrivacy(channelID string, privacy model.ChannelType) (*model.Channel, *model.Response, error)
	CreateBot(bot *model.Bot) (*model.Bot, *model.Response, error)
	PatchBot(userID string, patch *model.BotPatch) (*model.Bot, *model.Response, error)
	GetBots(page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	GetBotsIncludeDeleted(page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	GetBotsOrphaned(page, perPage int, etag string) ([]*model.Bot, *model.Response, error)
	DisableBot(botUserID string) (*model.Bot, *model.Response, error)
	EnableBot(botUserID string) (*model.Bot, *model.Response, error)
	AssignBot(botUserID, newOwnerID string) (*model.Bot, *model.Response, error)
	SetServerBusy(secs int) (*model.Response, error)
	ClearServerBusy() (*model.Response, error)
	GetServerBusy() (*model.ServerBusyState, *model.Response, error)
	CheckIntegrity() ([]model.IntegrityCheckResult, *model.Response, error)
	InstallPluginFromURL(string, bool) (*model.Manifest, *model.Response, error)
	InstallMarketplacePlugin(*model.InstallMarketplacePluginRequest) (*model.Manifest, *model.Response, error)
	GetMarketplacePlugins(*model.MarketplacePluginFilter) ([]*model.MarketplacePlugin, *model.Response, error)
	MigrateAuthToLdap(fromAuthService string, matchField string, force bool) (*model.Response, error)
	MigrateAuthToSaml(fromAuthService string, usersMap map[string]string, auto bool) (*model.Response, error)
	GetPing() (string, *model.Response, error)
	GetPingWithFullServerStatus() (map[string]string, *model.Response, error)
	CreateUpload(us *model.UploadSession) (*model.UploadSession, *model.Response, error)
	GetUpload(uploadID string) (*model.UploadSession, *model.Response, error)
	GetUploadsForUser(userID string) ([]*model.UploadSession, *model.Response, error)
	UploadData(uploadID string, data io.Reader) (*model.FileInfo, *model.Response, error)
	ListImports() ([]string, *model.Response, error)
	GetJob(id string) (*model.Job, *model.Response, error)
	GetJobs(page int, perPage int) ([]*model.Job, *model.Response, error)
	GetJobsByType(jobType string, page int, perPage int) ([]*model.Job, *model.Response, error)
	CreateJob(job *model.Job) (*model.Job, *model.Response, error)
	CancelJob(jobID string) (*model.Response, error)
	CreateIncomingWebhook(hook *model.IncomingWebhook) (*model.IncomingWebhook, *model.Response, error)
	UpdateIncomingWebhook(hook *model.IncomingWebhook) (*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhooks(page int, perPage int, etag string) ([]*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhooksForTeam(teamID string, page int, perPage int, etag string) ([]*model.IncomingWebhook, *model.Response, error)
	GetIncomingWebhook(hookID string, etag string) (*model.IncomingWebhook, *model.Response, error)
	DeleteIncomingWebhook(hookID string) (*model.Response, error)
	CreateOutgoingWebhook(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.Response, error)
	UpdateOutgoingWebhook(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooks(page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhook(hookID string) (*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooksForChannel(channelID string, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	GetOutgoingWebhooksForTeam(teamID string, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response, error)
	RegenOutgoingHookToken(hookID string) (*model.OutgoingWebhook, *model.Response, error)
	DeleteOutgoingWebhook(hookID string) (*model.Response, error)
	ListExports() ([]string, *model.Response, error)
	DeleteExport(name string) (*model.Response, error)
	DownloadExport(name string, wr io.Writer, offset int64) (int64, *model.Response, error)
	ResetSamlAuthDataToEmail(includeDeleted bool, dryRun bool, userIDs []string) (int64, *model.Response, error)
}
