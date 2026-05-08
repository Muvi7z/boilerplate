package converter

import (
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	entity2 "github.com/Muvi7z/boilerplate/iam/internal/repository/entity"
)

func SessionToRedisView(session entity.Session) entity2.SessionRedisView {
	return entity2.SessionRedisView{
		Uuid:   session.Uuid,
		UserId: session.UserId,
	}
}

func SessionFromRedisView(view entity2.SessionRedisView) entity.Session {
	return entity.Session{
		Uuid:   view.Uuid,
		UserId: view.UserId,
	}
}
