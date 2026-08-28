package api

func (a *server) findUser(email string) (user, bool) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	account, exists := a.users[email]
	return account, exists
}

func (a *server) addUser(email string, salt, hash []byte) (user, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if _, exists := a.users[email]; exists {
		return user{}, false
	}
	account := user{
		ID:           a.nextID,
		Email:        email,
		PasswordSalt: salt,
		PasswordHash: hash,
	}
	a.users[email] = account
	a.nextID++
	return account, true
}

func (a *server) saveSession(token string, value session) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.sessions[token] = value
}
