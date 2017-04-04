// Package podhost implements the domain employed by the app

package podhost

import (
	"github.com/solher/arangolite"
)

/*************************
********          ********
********   USER   ********
********          ********
*************************/

type User struct {
	arangolite.Document
	FirstName string
	LastName  string
	Nickname  string
	Username  string
	Email     string
	Password  string
	Active    bool
}

type UserService interface {
	User(key string) (*User, error)
	CreateUser(u *User) error
	DeleteUser(key string) error
}

// ROLES

type UserRole struct {
	arangolite.Document
	Role        string
	Description string
	RoleType    string
}

type UserRoleService interface {
	UserRole(key string) (*UserRole, error)
	UserRoles(user_key string) ([]*UserRole, error)
	CreateUserRole(u *UserRole) error
	DeleteUserRole(key string) error
}

type UserHasRole struct {
	arangolite.Document
	Weight int
}

type UserHasRoleService interface {
	UserHasRole(key string) (*UserHasRole, error)
	UserhasRoles(user_key string) ([]*UserHasRole, error)
	GiveUserRole(u *UserHasRole) error
	RemoveUserRole(key string) error
}

type UserHasPodcast struct {
	arangolite.Document
}

type UserHasPodcastService interface {
	UserHasPodcast(key string) (*UserHasPodcast, error)
	UserHasPodcasts(user_key string) ([]*UserHasPodcast, error)
	GiveUserPodcast(u *UserHasPodcast) error
	RemoveUserPodcast(key string) error
}

/*************************
*******           ********
*******  PODCAST  ********
*******           ********
*************************/

type Podcast struct {
	arangolite.Document
	Title         string
	Author        string
	Description   string
	Link          string
	Language      string
	Copyright     string
	ContentRating string
	PubDate       string
	Active        bool
}

type PodcastService interface {
	Podcast(key string) (*Podcast, error)
	CreatePodcast(*Podcast) error
	DeletePodcast(key string) error
}

// EPISODES

type Episode struct {
	arangolite.Document
	Title   string
	PubDate string
	URL     string
	Length  int
	Link    string
	Number  string
}

type EpisodeService interface {
	Episode(key string) (*Episode, error)
	Episodes(podcast_key string) ([]*Episode, error)
	CreateEpisode(e *Episode) error
	DeleteEpisode(key string) error
}

type PodcastHasEpisode struct {
	arangolite.Document
	Number int
}

type PodcastHasEpisodeService interface {
	PodcastHasEpisode(key string) (*PodcastHasEpisode, error)
	PodcastHasEpisodes(podcast_key string) ([]*PodcastHasEpisode, error)
	LinkPodcastEpisode(p *PodcastHasEpisode) error
	UnlinkPodcastEpisode(key string) error
}

// RSS FEEDS

type RssFeed struct {
	arangolite.Document
	Name        string
	URL         string
	Description string
	Active      bool
}

type RssFeedService interface {
	RssFeed(key string) (*RssFeed, error)
	CreateRssFeed(r *RssFeed) error
	DeleteRssFeed(key string) error
}

type PodcastHasRssFeed struct {
	arangolite.Document
	Main   bool
	Weight int
}

type PodcastHasRssFeedService interface {
	PodcastHasRssFeed(key string) (*PodcastHasRssFeed, error)
	PodcastHasRssFeeds(podcast_key string) ([]*PodcastHasRssFeed, error)
	LinkPodcastRssFeed(p *PodcastHasRssFeed) error
	UnlinkPodcastRssFeed(key string) error
}

// CATEGORY

type PodcastCategory struct {
	arangolite.Document
	Category    string
	Description string
}

type PodcastCategoryService interface {
	PodcastCategory(key string) (*PodcastCategory, error)
	PodcastHasCategories(podcast_key string) ([]*PodcastCategory, error)
	CreatePodcastCategory(p *PodcastCategory) error
	DeletePodcastCategory(key string) error
}

type PodcastHasCategory struct {
	arangolite.Document
	Weight int
}

type PodcastHasCategoryService interface {
	PodcastHasCategory(key string) (*PodcastHasCategory, error)
}
