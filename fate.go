package fate_go_sdk

import (
	"net/http"
	"net/url"
	"strconv"
	pb "github.com/zm-dev/fate-go-sdk/pb"
	"context"
	"google.golang.org/grpc"
)

func New(opt ...Option) (*Fate, error) {
	opts := defaultOptions
	for _, o := range opt {
		o(&opts)
	}
	conn, err := grpc.Dial(opts.FateRpcHostname, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	atc := NewAccessTokenClient(
		opts.AppID,
		opts.AppSecret,
		pb.NewAccessTokenServiceClient(conn),
		opts.RpcTimeout,
		opts.Cache,
	)

	lc := NewLoginChecker(
		atc,
		pb.NewLoginCheckerClient(conn),
		opts.RpcTimeout,
		opts.AccessTokenKey,
	)
	return &Fate{lc: lc, opts: opts}, nil
}

type Fate struct {
	opts options
	lc   *LoginChecker
}

const loginMethodRedirect = "redirect"

func (f *Fate) Login(w http.ResponseWriter, r *http.Request) {
	switch f.opts.LoginMethod {
	case loginMethodRedirect:
		f.RedirectToLogin(w, r)
	default:
		f.RedirectToLogin(w, r)
	}
}

func (f *Fate) RedirectToLogin(w http.ResponseWriter, r *http.Request) {
	q := make(url.Values)
	q.Set("app_id", strconv.Itoa(int(f.opts.AppID)))
	var scheme string
	if r.TLS == nil {
		scheme = "http://"
	} else {
		scheme = "https://"
	}
	q.Set("callback", scheme+r.Host+r.RequestURI)
	http.Redirect(w, r, f.opts.FateUrl+"/?"+q.Encode(), http.StatusFound)
}

func (f *Fate) RedirectToLoginWithCallback(callback string, w http.ResponseWriter, r *http.Request) {
	q := make(url.Values)
	q.Set("app_id", strconv.Itoa(int(f.opts.AppID)))
	q.Set("callback", callback)
	http.Redirect(w, r, f.opts.FateUrl+"/?"+q.Encode(), http.StatusFound)
}

func (f *Fate) RedirectToLogout(callback string, w http.ResponseWriter, r *http.Request) {
	q := make(url.Values)
	q.Set("app_id", strconv.Itoa(int(f.opts.AppID)))
	q.Set("callback", callback)
	http.Redirect(w, r, f.opts.FateUrl+"/logout?"+q.Encode(), http.StatusFound)
}

func (f *Fate) getLoginCheckResContext(r *http.Request) (*http.Request, error) {
	if _, ok := fromLoginCheckResContext(r.Context()); !ok {
		if res, err := f.Check(r); err != nil {
			return r, err
		} else {
			return r.WithContext(newLoginCheckResContext(r.Context(), res)), nil
		}
	}
	return r, nil
}

type loginCheckResContextKey struct{}

func newLoginCheckResContext(ctx context.Context, res *pb.LoginCheckRes) context.Context {
	return context.WithValue(ctx, loginCheckResContextKey{}, res)
}

func fromLoginCheckResContext(ctx context.Context) (*pb.LoginCheckRes, bool) {
	res, ok := ctx.Value(loginCheckResContextKey{}).(*pb.LoginCheckRes)
	return res, ok
}

func (f *Fate) Check(r *http.Request) (*pb.LoginCheckRes, error) {
	ticketID, err := r.Cookie(f.opts.TicketIdCookieKey)
	if err != nil {
		return &pb.LoginCheckRes{IsLogin: false}, nil
	}
	if res, err := f.lc.Check(ticketID.Value); err != nil {
		return nil, err
	} else {
		return res, err
	}
}

func (f *Fate) Logout(ticketID string) (error) {
	return f.lc.Logout(ticketID)
}

func GetIsLogin(ctx context.Context) bool {
	if res, ok := fromLoginCheckResContext(ctx); ok {
		return res.GetIsLogin();
	} else {
		return false;
	}
}

func GetUserID(ctx context.Context) int64 {
	if res, ok := fromLoginCheckResContext(ctx); ok {
		return res.GetUserId();
	} else {
		return 0;
	}
}
