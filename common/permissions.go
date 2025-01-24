package common

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	path = "users.txt"
)

var admins = []int64{AdminChatId}

type PermissionsService struct {
	usersLimit int
	users      []int64
}

func NewPermissionsService(usersLimit int) *PermissionsService {
	service := PermissionsService{usersLimit: usersLimit, users: admins}
	service.users = append(service.users, readUsersFromFile()...)
	return &service
}

func (p *PermissionsService) SaveUser(ctx context.Context, bot *bot.Bot, userId int64) bool {
	if p.usersLimit > 0 {
		AppendToFile(path, strconv.FormatInt(userId, 10))
		p.users = append(p.users, userId)
		p.usersLimit--
		SendTextToAdminChannel(ctx, bot, fmt.Sprintf("С ботом начал общаться %d", userId))
		return true
	} else {
		SendTextToAdminChannel(ctx, bot, fmt.Sprintf("С ботом *пытался* общаться %d", userId))
		return false
	}
}

func (p *PermissionsService) RemoveUsers() error {
	p.users = admins
	p.usersLimit = 0
	return CreateEmptyFile(path)
}

func (p *PermissionsService) getLimit() int {
	return p.usersLimit
}

func (p *PermissionsService) increaseLimit() {
	p.usersLimit++
}

func (p *PermissionsService) CheckUserIsAcceptable(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if !slices.Contains(p.users, update.Message.Chat.ID) && !p.SaveUser(ctx, b, update.Message.Chat.ID) {
			return
		}

		next(ctx, b, update)
	}
}
func (p *PermissionsService) RemoveAllUsers(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message.Chat.ID != AdminChatId {
			return
		}
		p.RemoveUsers()
		SendTextToAdminChannel(ctx, b, fmt.Sprintf("Забрали доступ у всех посетителей. Лимит = 0"))
		next(ctx, b, update)
	}
}

func (p *PermissionsService) AddUser(next bot.HandlerFunc) bot.HandlerFunc {
	return func(ctx context.Context, b *bot.Bot, update *models.Update) {
		if update.Message.Chat.ID != AdminChatId {
			return
		}
		p.increaseLimit()
		SendTextToAdminChannel(ctx, b, fmt.Sprintf("Повышен лимит. Текущий лимит = %d", p.usersLimit))
		next(ctx, b, update)
	}
}

func EmptyHandler(ctx context.Context, b *bot.Bot, update *models.Update) {

}

func readUsersFromFile() []int64 {
	users := ReadStringFromFile(path)
	split := strings.Split(users, "\n")
	result := make([]int64, len(split))
	for _, userId := range split {
		atoi, _ := strconv.Atoi(userId)
		result = append(result, int64(atoi))
	}
	return result
}
