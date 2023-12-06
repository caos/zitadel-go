package main

import (
	"context"
	"embed"
	_ "embed"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"

	"github.com/zitadel/zitadel-go/v3/pkg/authentication"
	openid "github.com/zitadel/zitadel-go/v3/pkg/authentication/oidc"
	"github.com/zitadel/zitadel-go/v3/pkg/zitadel"
)

var (
	// flags to be provided for running the example server
	domain      = flag.String("domain", "", "your ZITADEL instance domain (in the form: https://<instance>.zitadel.cloud or https://<yourdomain>)")
	key         = flag.String("key", "", "path to your key.json")
	clientID    = flag.String("clientID", "", "clientID provided by ZITADEL")
	redirectURI = flag.String("redirectURI", "", "redirectURI registered at ZITADEL")
	port        = flag.String("port", "8089", "port to run the server on (default is 8089)")

	//go:embed "templates/*.html"
	templates embed.FS

	// tasks are used to store an in-memory list used in the protected endpoint
	tasks []string
)

func main() {
	flag.Parse()

	ctx := context.Background()

	t, err := template.New("").ParseFS(templates, "templates/*.html")
	if err != nil {
		slog.Error("unable to parse template", "error", err)
		os.Exit(1)
	}

	// Initiate the zitadel sdk by providing its domain
	// and as this example will focus on authentication (using OIDC / OAuth2 PKCE flow),
	// you will also need to initialize that with the generated client_id.
	//
	// it's a short form of:
	// cookieHandler := http2.NewCookieHandler([]byte(*key), []byte(*key))
	//z, err := zitadel.New(*domain,
	//	zitadel.WithAuthentication(ctx, *key,
	//		openid.WithCodeFlow[*openid.UserInfoContext[*oidc.IDTokenClaims, *oidc.UserInfo], *oidc.IDTokenClaims, *oidc.UserInfo](
	//			openid.PKCEAuthentication(*clientID, *redirectURI, []string{"openid","profile","email"}, cookieHandler,
	//		),
	//	),
	//)
	z, err := zitadel.New(*domain,
		zitadel.WithAuthentication(ctx, *key,
			openid.DefaultAuthentication(*clientID, *redirectURI, *key),
		),
	)
	if err != nil {
		slog.Error("zitadel sdk could not initialize", "error", err)
		os.Exit(1)
	}

	mw := authentication.Middleware(z.Authentication)

	router := http.NewServeMux()

	router.Handle("/auth/", z.Authentication)
	router.Handle("/profile", mw.RequireAuthentication()(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authCtx := mw.Context(req.Context())
		data, err := json.MarshalIndent(authCtx.UserInfo, "", "  ")
		_ = err
		s := string(data)
		_ = s
		err = t.ExecuteTemplate(w, "profile.html", authCtx.UserInfo)
		_ = err
	})))
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if authentication.IsAuthenticated(req.Context()) {
			http.Redirect(w, req, "/profile", http.StatusFound)
			return
		}
		err = t.ExecuteTemplate(w, "home.html", nil)
	}))

	// start the server on the specified port (default http://localhost:8089)
	lis := fmt.Sprintf(":%s", *port)
	slog.Info("server listening, press ctrl+c to stop", "addr", "http://localhost"+lis)
	err = http.ListenAndServe(lis, router)
	if !errors.Is(err, http.ErrServerClosed) {
		slog.Error("server terminated", "error", err)
		os.Exit(1)
	}
}

func parseTemplates() {

}
