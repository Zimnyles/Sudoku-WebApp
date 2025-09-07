package middleware

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ActivityMiddleware(store *session.Store, dbpool *pgxpool.Pool) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}
		if userEmail := sess.Get("email"); userEmail != nil {
			_, err := dbpool.Exec(context.Background(),
				`UPDATE users 
                 SET is_active = true, last_seen = now() 
                 WHERE email = $1`, userEmail)
			if err != nil {
				fmt.Println("failed to update activity:", err)
			}
		}
		return c.Next()
	}
}
