package middleware

import (
	"project-devboard/pkg/apperrors"

	"github.com/gin-gonic/gin"
)

// Bu middleware'i kullanmadan önce AuthenticationMiddleware'i kullanarak kullanıcı rolünü c.Set("role", userRole) ile context'e eklemelisiniz. Bu middleware, belirtilen rollerden herhangi birine sahip olmayan kullanıcıların erişimini engeller ve uygun hata mesajları döner.
// -------------------------------------------------------------------
// requiredRoles parametresi, erişim izni verilen rollerin bir listesidir. Örneğin, AuthorizationMiddleware("admin", "editor") sadece "admin" veya "editor" rollerine sahip kullanıcıların erişimine izin verir.
func AuthorizationMiddleware(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.Error(apperrors.New(apperrors.Unauthorized, apperrors.ErrUnauthorized))
			c.Abort()
			return
		}
		userRole, ok := role.(string)
		if !ok {
			c.Error(apperrors.New(apperrors.Unauthorized, apperrors.ErrUnauthorized))
			c.Abort()
			return
		}
		for _, requiredRole := range requiredRoles {
			if userRole == requiredRole {
				c.Next()
				return
			}
		}
		c.Error(apperrors.New(apperrors.Forbidden, apperrors.ErrForbidden))
		c.Abort()
	}
}
