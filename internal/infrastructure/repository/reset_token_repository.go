package repository

import (
	"sync"
	"time"

	"github.com/Elizabethyonas/A2SV-Portal-Project/internal/domain/entities"
)

type ResetTokenRepository struct {
	mu     sync.Mutex
	tokens map[string]entities.ResetToken
}

func NewResetTokenRepository() *ResetTokenRepository {
	return &ResetTokenRepository{
		tokens: make(map[string]entities.ResetToken),
	}
}

func (r *ResetTokenRepository) Save(token string, userID string, expiresAt time.Time) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tokens[token] = entities.ResetToken{
		Token:     token,
		UserID:    userID,
		ExpiresAt: expiresAt,
	}
}

func (r *ResetTokenRepository) Get(token string) (entities.ResetToken, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rt, exists := r.tokens[token]
	return rt, exists
}

func (r *ResetTokenRepository) Delete(token string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.tokens, token)
}
